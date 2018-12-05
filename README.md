# grpc_demo
gRPC演示

# protobuf 下载安装
https://github.com/protocolbuffers/protobuf/releases/tag/v3.3.0 下载平台对应的包，Linux下使用源码编译，Windows下直接使用编译好的文件
将安装目录/bin添加到环境变量里面

# proto3 接口文件语法介绍
https://blog.csdn.net/u011518120/article/details/54604615 
https://blog.csdn.net/lyjshen/article/details/52298003

# proto 工具生产代码
protoc -I protos protos/helloworld.proto --go_out=plugins=grpc:src/greeter
或者
protoc --go_out=plugins=grpc:. route_guide.proto   //基于当前目录
正式打包环境请编写一键打包sh脚本

# grpc 简介
https://smallnest.gitbooks.io/go-rpc-programming-guide/content/part1/grpc.html

# golang grpc开发包
go get google.golang.org/grpc    需要VPN下载，或者github上下载对应的包
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

# grpc demo演示
http://helight.info/2018-02-06/golang-grpc%E7%AE%80%E5%8D%95%E4%BD%BF%E7%94%A8/

# grpc-gateway：grpc转换为http协议对外提供服务
https://blog.csdn.net/dapangzi88/article/details/63686334

# golang命令行库cobra的使用
https://www.cnblogs.com/borey/p/5715641.html

# go-swagger
https://www.jianshu.com/p/dbca3911419e

# 使用 SwaggerUI 创建 Golang API 文档
https://studygolang.com/articles/12354?fr=sidebar

# Go 语言打包静态文件 go-bindata
http://www.mamicode.com/info-detail-2158143.html
go get -u github.com/jteeuwen/go-bindata/...

生成pb.go文件（在helloworld文件夹外面执行）
protoc -I helloworld -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
helloworld/helloworld.proto 


生成swagger.json文件（在helloworld文件夹外面执行）
protoc -I helloworld -I. \
-I$GOPATH/src  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--swagger_out=logtostderr=true:. \
helloworld/helloworld.proto

2 pb.gw.go文件（在helloworld文件夹外面执行）
protoc -I helloworld -I. \
-I$GOPATH/src  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--grpc-gateway_out=logtostderr=true:. \
helloworld/helloworld.proto


//仅仅参考，验证不成功
# 编译 google.api
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

#编译hello_http.proto为hello_http.pb.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=grpc-hello-world/proto/google/api:. ./hello.proto

#编译hello_http.proto为hello_http.pb.gw.proto
protoc --grpc-gateway_out=logtostderr=true:. ./hello.proto




curl -X POST -k http://localhost:8080/v1/example/echo -d '{"name": " world"}'
 
{"message":"Hello  world"}

//生成datafile.go文件（在swagger-ui目录下执行）
go-bindata --nocompress -pkg ../swagger -o ../swagger/datafile.go ./...

//安装go-bindata-assetfs,它能够使用go-bindata所生成Swagger UI的Go代码，结合net/http对外提供服务
go get github.com/elazarl/go-bindata-assetfs/...


https://github.com/EDDYCJY/grpc-hello-world.git