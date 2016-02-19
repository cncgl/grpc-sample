package main


import (
  "flag"
  pb "github.com/cncgl/grpc-sample/proto"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "log"
)

var (
  addrFlag = flag.String("addr", "localhost:5000", "server address host:post")
  option = grpc.WithInsecure()
)

func main() {
  //IPアドレス(ここではlocalhost)とポート番号(ここでは5000)を指定して、サーバーと接続する
  // conn, err := grpc.Dial(*addrFlag)
  conn, err := grpc.Dial(*addrFlag, option)

  if err != nil {
    log.Fatalf("Connection error: %v", err)
  }

  //接続は最後に必ず閉じる
  defer conn.Close()

  c := pb.NewGreeterClient(conn)

  //サーバーに対してリクエストを送信する
  resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "cncgl"})
  if err != nil {
    log.Fatalf("RPC error: %v", err)
  }
  log.Printf("Greeting: %s", resp.Message)
}
