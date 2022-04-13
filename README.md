# A gRPC Client in Go

### Protoc compiler

To be able to generate the needed client files, you must use 'protoc' compiler.

http://google.github.io/proto-lens/installing-protoc.html

### Proto Files

You may find ready to use .proto files in the 'proto' directory. You may run the server to serve those 
services by running the project:

https://github.com/guildenstern70/SmartgRpc

### Protoc Go Gen Files

Be sure to locally install the latest version of Go files generator for protoc, here:

    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

### Server gRPC

The gRpc server for this client can be found at:

https://github.com/guildenstern70/SmartgRpc

### Generate client code

Md5 Service 

    protoc --go_out=./out --go_opt=paths=source_relative --go-grpc_out=./out --go-grpc_opt=paths=source_relative ./proto/md5.proto

Hypotenuse

    protoc --go_out=./out --go_opt=paths=source_relative --go-grpc_out=./out --go-grpc_opt=paths=source_relative ./proto/hypotenuse.proto
    