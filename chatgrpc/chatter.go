package chatgrpc

import (
  "fmt"
  "os"
  "github.com/joho/godotenv"
  "net"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) ChatSender(ctx context.Context, in *Msg) (*Status, error) {
  fmt.Println("Received message inside Channel ChatSender: ", in.Msg)
  return &Status{Status: true, Timestamp: in.Timestamp, UniqueNum: in.UniqueNum}, nil
}

func (s *Server) ChatListener(ctx context.Context, in *Msg) (*Status, error) {
  fmt.Println("Received message inside Channel ChatListener: ", in.Msg)
  return &Status{Status: true, Timestamp: in.Timestamp, UniqueNum: in.UniqueNum}, nil
}

func CallChatSender(s *Server, ctx context.Context, in *Msg) {
  s.ChatSender(ctx, in)
}

func CallChatListener(s *Server, ctx context.Context, in *Msg) {
  s.ChatListener(ctx, in)
}

func ListenerGRPCServer() {
  env_err := godotenv.Load(".env")
  if env_err != nil {
    panic(env_err.Error())
  }

  port := os.Getenv("grpc_port")

  fmt.Println("Starting GRPC server on port: ", port)

  lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
  if err != nil {
    fmt.Println("failed to listen: ", err)
  }

  serveIt := Server{}
  grpcServer := grpc.NewServer()
  RegisterChatterInterfaceServer(grpcServer, &serveIt)
  err = grpcServer.Serve(lis)
  if err != nil {
    fmt.Println("Failed to server: ", err)
  }
}
