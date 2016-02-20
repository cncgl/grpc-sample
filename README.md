# grpc-sample
sample code for gRPC

## Install

### GO の場合
proto 以下には 生成されたスタブファイルがありますが、自分で生成するには protocol buffer が必要になります。
```
$ git clone https://github.com/google/protobuf.git
$ cd protobuf
$ ./autogen.sh
$ ./configure
$ make
$ make check
$ sudo make install
```

protoc の golang 用ジェネレータをプラグインとしてインストールします。
```
$ go get github.com/golang/protobuf/protoc-gen-go
```

オプション指定してコンパイルします。
```
$ protoc --go_out=plugins=grpc:. proto/helloworld.proto
```

### Ruby の場合
proto 以下には 生成されたスタブファイルがありますが、Ruby のスタブファイルの生成には protocol buffer を使用します。
```
$ protoc -I proto --ruby_out=lib --grpc_out=lib --plugin=protoc-gen-grpc=`which grpc_ruby_plugin` proto/helloworld.proto
```


## Execute

サーバー側
```
$ go run hello_server.go
```

クライアント側
```
$ go run hello_client.go
```

## License
[MIT](LICENSE)
