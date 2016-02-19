# grpc-sample
sample code for gRPC

## Install

### GO の場合
proto 以下には 生成されたスタブファイルがありますが、自分で生成するには protocol buffer が必要になります。
```
$ git clone https://github.com/google/protobuf.git
$ cd $_
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
Gemfile でインストールされていますが、Ruby のスタブファイルの生成には Ruby 版の protocol buffer を使用します。
```
$ bundle exec protoc --ruby_out . proto/helloworld.proto
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
