# grpc-sample
sample code for [gRPC](https://github.com/grpc/grpc)

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

### サーバー側

#### Go の場合
```
$ go run hello_server.go
```

#### Ruby の場合
```
$ ruby hello_server.rb
```

### クライアント側

#### Go の場合
```
$ go run hello_client.go
```

#### Ruby の場合
```
$ ruby hello_client.rb
```


## License
[MIT](LICENSE)
