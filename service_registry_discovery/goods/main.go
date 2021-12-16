package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	// ServiceGoods 商品服务
	ServiceGoods = "goods"
)

// serviceEndpointKeyPrefix 服务入口在 etcd 存储的 key 前缀
var serviceEndpointKeyPrefix = "/service_registry_discovery"

// hostname 主机名
var hostname string

// endpoint 访问入口
var endpoint string

// servicePort 服务端口
var servicePort int = 80

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

// init 初始化
func init() {
	hostname, _ = os.Hostname()
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				endpoint = fmt.Sprintf("%s:%d", ipnet.IP.String(), servicePort)
				break
			}
		}
	}
}

func main() {
	// 服务注册
	go func() {
		ServiceRegistry()
	}()
	http.HandleFunc(
		"/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprint(w, "goods")
		},
	)
	http.HandleFunc("/goods/list", GetGoodsList)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", servicePort), nil)
}

// GetGoodsList 获取商品列表
func GetGoodsList(w http.ResponseWriter, r *http.Request) {
	// 生成订单信息
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	params := make(map[string]interface{})
	_ = json.Unmarshal(body, &params)
	var traceID = params["trace_id"]
	var res = map[string]interface{}{
		"trace_id": traceID,
		"message":  "get goods list success",
	}
	b, _ := json.Marshal(res)
	fmt.Printf(
		"trace_id: %v, goods list result: %s\n", traceID, string(b),
	)
	_, _ = fmt.Fprint(w, string(b))
}

// ServiceRegistry 服务注册
func ServiceRegistry() {
	hostname, _ = os.Hostname()
	cli, err := clientv3.New(etcdCfg)
	if err != nil {
		panic(err)
	}
	key := fmt.Sprintf("%s/%s/%s", serviceEndpointKeyPrefix, ServiceGoods, hostname)
	ctx := context.Background()
	// 过期时间: 3秒钟
	ttl := 3
	// 创建租约
	lease, err := cli.Grant(ctx, int64(ttl))
	if err != nil {
		panic(err)
	}
	b, _ := json.Marshal(lease)
	fmt.Printf("grant lease suucess: %s\n", string(b))
	// put kv
	res, err := cli.Put(ctx, key, endpoint, clientv3.WithLease(lease.ID))
	if err != nil {
		panic(err)
	}
	b, _ = json.Marshal(res)
	fmt.Printf("put kv with lease suucess: %s\n", string(b))
	// 保持租约不过期
	klRes, err := cli.KeepAlive(ctx, lease.ID)
	if err != nil {
		panic(err)
	}
	// 监听续约情况
	for v := range klRes {
		b, _ = json.Marshal(v)
		fmt.Printf("keep lease alive suucess: %s\n", string(b))
	}
	fmt.Println("stop keeping lease alive")
}
