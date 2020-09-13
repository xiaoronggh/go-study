## rpc

```shell script
go env -w GOPROXY=https://goproxy.io,direct
```

```shell script
# 生成服务端的go代码
protoc --go_out=plugins=grpc:./proto proto/service.proto 
```

[参考](https://www.cnblogs.com/oolo/p/11840305.html)