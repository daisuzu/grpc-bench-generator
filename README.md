grpc-bench-generator
====================

grpc-bench-generator is a tool to generate benchmark client from the protocol buffer package of Go.

Installation
------------

```sh
go get -u github.com/daisuzu/grpc-bench-generator
```

Example
-------

```sh
grpc-bench-generator -path $GOPATH/src/google.golang.org/grpc/examples/helloworld/helloworld > main.go
go run main.go -server=localhost:50051 -duration=1 -enable_ssl=false
```

License
-------

Licensed under the BSD-3-Clause License.
See [LICENSE](LICENSE).
