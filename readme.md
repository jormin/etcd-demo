Etcd Demo
============

[![Go Report Card](https://goreportcard.com/badge/github.com/jormin/etcd-demo)](https://goreportcard.com/report/github.com/jormin/etcd-demo)
[![](https://img.shields.io/badge/version-v1.0.0-success.svg)](https://github.com/jormin/etcd-demo)

用 Golang 调用 Etcd Demo。

#### 基础操作

| 模块        | 命令                   | 说明                                         | 进度 |
| ----------- | ---------------------- | -------------------------------------------- | ---- |
| version     | version                | 查看版本信息                                 | ---  |
| endpoint    | endpoint status        | 查看节点状态                                 | ✅    |
|             | endpoint hashkv        | 查看节点历史hash                             | ✅    |
|             | endpoint health        | 查看节点状态                                 | ---  |
| member      | member list            | 获取成员列表                                 | ✅    |
|             | member add             | 添加成员                                     | ✅    |
|             | member update          | 更新成员                                     | ✅    |
|             | member promote         | 提升学习者成员为常规成员                     | ✅    |
|             | member remove          | 移除成员                                     | ✅    |
| move-leader | move-leader            | 将节点设为leader                             | ✅    |
| auth        | auth status            | 查看当前鉴权开启状态                         | ✅    |
|             | auth disable           | 禁用鉴权                                     | ✅    |
|             | auth enable            | 启用鉴权                                     | ✅    |
| role        | role list              | 获取角色列表                                 | ✅    |
|             | role add               | 添加角色                                     | ✅    |
|             | role get               | 查看角色信息                                 | ✅    |
|             | role delete            | 删除角色                                     | ✅    |
|             | role grant-permission  | 角色授权                                     | ✅    |
|             | role revoke-permission | 撤销角色授权                                 | ✅    |
| user        | user add               | 添加用户                                     | ✅    |
|             | user list              | 用户列表                                     | ✅    |
|             | user get               | 获取用户信息                                 | ✅    |
|             | user delete            | 删除用户                                     | ✅    |
|             | user passwd            | 修改密码                                     | ✅    |
|             | user grant-role        | 用户绑定角色                                 | ✅    |
|             | user revoke-role       | 撤销用户角色                                 | ✅    |
| lease       | lease list             | 租约列表                                     | ✅    |
|             | lease grant            | 生成租约                                     | ✅    |
|             | lease revoke           | 撤销租约                                     | ✅    |
|             | lease timetolive       | 查看租约期限信息                             | ✅    |
|             | lease keep-alive       | 保持租约不过期                               | ✅    |
| normal      | put                    | 设置kv                                       | ✅    |
|             | get                    | 查看kv                                       | ✅    |
|             | del                    | 删除kv                                       | ✅    |
|             | watch                  | 监听key                                      | ✅    |
| txn         | txn                    | 开启事务                                     | ✅    |
| lock        | lock                   | 分布式锁                                     |      |
| snapshot    | snapshot save          | 保存快照                                     | ✅    |
|             | snapshot status        | 校验快照                                     | ---  |
|             | snapshot restore       | 从快照恢复                                   | ---  |
| alarm       | alarm list             | 获取警报列表                                 | ✅    |
|             | alarm disarm           | 解除警报                                     | ✅    |
| check       | check datascale        | 检查内存使用情况                             | ✅    |
|             | check perf             | 查看集群性能                                 | ✅    |
| other       | compaction             | 压缩 etcd 中的事件历史                       | ---  |
|             | completion             | 生成完成脚本                                 | ---  |
|             | defrag                 | 对具有给定端点的 etcd 成员的存储进行碎片整理 | ---  |
|             | elect                  | 观察并参与leader选举                         | ---  |
|             | make-mirror            | 在目标 etcd 集群上创建镜像                   | ---  |



License
-------

under the [MIT](./LICENSE) License