# go-redis-proxy
A go implement redis proxy.

Go Agent For PHP

## 定位

本机Client，为php提供扩展服务

## 协议

redis get

### Request

{"id":"","method":"","action":"","args":""}

### Response

{"em":"", "ec":"", "data":OBJECT}

## 本地服务

redis (grpc/protobuf/thrift)

## 功能列表

|function|method|action|args|response_data|
|:---|---|---|---|---|
|缓存|cache|set| {"key":"", "value":""}||
||cache|get|{"key":""}||
|redis 连接池|redis|get|{"host":","port":3,"key":""}||
||redis|set |{"host":","port":3,"key":"", "value":""}||
||redis|hset |{"host":"", "port":2,"key":"","field":"", "value":""}||
||redis|hget |{"host":"", "port":2,"key":"","field":""}||

## 其他功能

1. Mongo Proxy
2. MySQL Proxy
3. Kafka Proxy
4. ** Proxy

## 系统组件

1. 日志输出
2. 缓存查看及删除
3. 连接池状态查看
4. 系统状态
5. 自动重连