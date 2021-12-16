package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	// ServiceGoods 商品服务
	ServiceGoods = "goods"
)

// serviceEndpointKeyPrefix 服务入口在 etcd 存储的 key 前缀
var serviceEndpointKeyPrefix = "/service_registry_discovery"

// serviceEndpoints 服务入口列表
var serviceEndpoints = map[string]map[string]string{
	ServiceGoods: {},
}

// 全局服务锁
var serviceLocker = sync.Mutex{}

// etcdCfg Etcd配置
var etcdCfg = clientv3.Config{
	Endpoints: []string{
		"http://etcd-1.etcd-headless.devops.svc.cluster.local:2379",
		"http://etcd-2.etcd-headless.devops.svc.cluster.local:2379",
		"http://etcd-3.etcd-headless.devops.svc.cluster.local:2379",
	},
	DialTimeout:          time.Second * 30,
	DialKeepAliveTimeout: time.Second * 30,
	Username:             "root",
	Password:             "90CjPHPRlxw=",
}

// etcdCfg Etcd配置
// var etcdCfg = clientv3.Config{
// 	Endpoints: []string{
// 		"127.0.0.1:12379",
// 		"127.0.0.1:22379",
// 		"127.0.0.1:32379",
// 	},
// 	DialTimeout:          time.Second * 30,
// 	DialKeepAliveTimeout: time.Second * 30,
// 	Username:             "root",
// 	Password:             "111111",
// }

// servicePort 服务端口
var servicePort int = 80

func main() {
	// 监听服务入口
	go func() {
		ServiceDiscovery()
	}()
	http.HandleFunc(
		"/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprint(w, "gateway")
		},
	)
	http.HandleFunc("/goods/list", GetGoodsList)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", servicePort), nil)
}

// GetGoodsList 获取商品列表
func GetGoodsList(w http.ResponseWriter, r *http.Request) {
	var traceID = time.Now().Unix()
	var res = map[string]interface{}{
		"trace_id": traceID,
		"code":     0,
	}
	client, endpoint, err := GetSvcEndpoin(ServiceGoods)
	fmt.Printf(
		"trace_id: %d, get goods endpoint result: client=%s, endpoint=%s, error=%s\n", traceID, client, endpoint, err,
	)
	if err != nil {
		res["code"] = -1
		res["message"] = err.Error()
		b, _ := json.Marshal(res)
		_, _ = fmt.Fprint(w, string(b))
		return
	}
	url := fmt.Sprintf("http://%s/goods/list", endpoint)
	m := map[string]interface{}{
		"trace_id": time.Now().Unix(),
	}
	body, _ := json.Marshal(m)
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	b, _ := ioutil.ReadAll(resp.Body)
	var jsonResp map[string]interface{}
	_ = json.Unmarshal(b, &jsonResp)
	res["message"] = jsonResp["message"]
	b, _ = json.Marshal(res)
	fmt.Printf(
		"trace_id: %d, get goods list result: %s\n", traceID, string(b),
	)
	_, _ = fmt.Fprint(w, string(b))
}

// ServiceDiscovery 服务发现
func ServiceDiscovery() {
	cli, err := clientv3.New(etcdCfg)
	if err != nil {
		panic(err)
	}
	for k, _ := range serviceEndpoints {
		go func(svc string) {
			ctx := context.Background()
			serviceKey := fmt.Sprintf("%s/%s", serviceEndpointKeyPrefix, svc)
			// 获取当前所有服务入口
			getRes, _ := cli.Get(ctx, serviceKey, clientv3.WithPrefix())
			serviceLocker.Lock()
			for _, v := range getRes.Kvs {
				serviceEndpoints[svc][string(v.Key)] = string(v.Value)
			}
			serviceLocker.Unlock()
			fmt.Printf(
				"[service_endpoint_change] [%s] service %s get endpoints success, %v\n", svc, svc,
				serviceEndpoints[svc],
			)
			ch := cli.Watch(ctx, serviceKey, clientv3.WithPrefix(), clientv3.WithPrevKV())
			for v := range ch {
				for _, v := range v.Events {
					key := string(v.Kv.Key)
					endpoint := string(v.Kv.Value)
					preEndpoint := ""
					if v.PrevKv != nil {
						preEndpoint = string(v.PrevKv.Value)
					}
					switch v.Type {
					// PUT，新增或替换
					case 0:
						serviceLocker.Lock()
						serviceEndpoints[svc][key] = endpoint
						serviceLocker.Unlock()
						fmt.Printf(
							"[service_endpoint_change] service %s put endpoint, key: %s, endpoint: %s\n", svc,
							key, endpoint,
						)
					// DELETE
					case 1:
						serviceLocker.Lock()
						delete(serviceEndpoints[svc], key)
						serviceLocker.Unlock()
						fmt.Printf(
							"[service_endpoint_change] service %s delete endpoint, key: %s, endpoint: %s\n",
							svc, key, preEndpoint,
						)
					}
				}
			}
		}(k)
	}
}

// GetSvcEndpoin 获取服务入口
func GetSvcEndpoin(svc string) (key, endpoint string, err error) {
	endpoints := serviceEndpoints[svc]
	if len(endpoints) == 0 {
		return "", "", errors.New(fmt.Sprintf("%s服务不可用，请稍后再试", svc))
	}
	num := len(endpoints)
	keys := make([]string, num)
	for v := range endpoints {
		keys = append(keys, v)
	}
	randomKey := keys[rand.Intn(len(keys))]
	return randomKey, endpoints[randomKey], nil
}
