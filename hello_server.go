package main

import (
  "flag"
  pb "github.com/cncgl/grpc-sample/proto"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "log"
  "net"
)

var (
  addrFlag = flag.String("addr", ":5000", "Address host:post")
)

type server struct{}

//リクセスト(Name)を受け取り、レスポンス(Message)を返す
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
  log.Printf("New Request: %v", in.String())
  return &pb.HelloReply{Message: "Hello, " + in.Name + "!"}, nil
}

func main() {
  //requestを受け付けるportを指定する
  lis, err := net.Listen("tcp", *addrFlag)

  if err != nil {
    log.Fatalf("boo")
  }

  //新しいgRPCサーバーのインスタンスを作成
  s := grpc.NewServer()
  //gRPCサーバーを保存する
  pb.RegisterGreeterServer(s, &server{})
  s.Serve(lis)
}