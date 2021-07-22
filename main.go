package main

import (
  "fmt"
  "os"
  "strconv"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/joho/godotenv"
  "github.com/amleshkashyap/someGoApp/chatgrpc"
  "github.com/amleshkashyap/someGoApp/pubnub"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pubnub "github.com/pubnub/go"
)

func newRouter() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/golang/hello", handler).Methods("GET")
  // /pubnub/ publishes to the channels that are subscribed to by the listener started in main() - a dummy client
  r.HandleFunc("/pubnub/world", publishToPubnubWorld).Methods("GET")
  r.HandleFunc("/pubnub/msg", publishToPubnubMsg).Methods("GET")
  // /grpc/ creates a grpc client for the chatgrpc package - it dials to and sends a message to the grpc server started in main() - a dummy client
  r.HandleFunc("/grpc/world", publishToGRPCWorld).Methods("GET")
  r.HandleFunc("/grpc/msg", publishToGRPCMsg).Methods("GET")
  return r
}

func httpServer() {
  port := os.Getenv("http_port")
  // The router is now formed by calling the `newRouter` constructor function
  // that we defined above. The rest of the code stays the same
  r := newRouter()
  fmt.Println("Starting HTTP Server on Port: ", port)
  err := http.ListenAndServe(port, r)
  if err != nil {
    panic(err.Error())
  }
}


func main() {
  // load the environment variables for use by the pubnub/grpc clients below
  env_err := godotenv.Load(".env")
  if env_err != nil {
    panic(env_err.Error())
  }

  // start the grpc server hosting the chatgrpc package using protobuf
  go func() {
    chatgrpc.ListenerGRPCServer()
  }()

  // start the pubnub event listeners
  pubsub.PubnubListener()
  // start the http server
  httpServer()
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World! From Golang")
}

// write a proper client - duplex channel protocols are supposed to from server to client
func publishToPubnubWorld(w http.ResponseWriter, r *http.Request) {
  configs := pubsub.SetupGlobalConfigs()
  pubnubHandler := pubnub.NewPubNub(configs)

  message := "Hello world from Pubnub via WORLD route"

  msg := map[string]interface{}{
    "msg": message,
  }

  // fmt.Println("Now Trying To Publish In WORLD Route")

  response, status, err := pubnubHandler.Publish().
    Channel("hello_world_pubnub").Message(msg).Execute()

  if err != nil {
     // Request processing failed.
     // Handle message publish error
  }
  _, _ = response, status
  // fmt.Println(response, status, err)
}

func publishToPubnubMsg(w http.ResponseWriter, r *http.Request) {
  configs := pubsub.SetupGlobalConfigs()
  pubnubHandler := pubnub.NewPubNub(configs)

  message := "Hello world from Pubnub via MSG route"

  msg := map[string]interface{}{
    "msg": message,
  }

  // fmt.Println("Now Trying To Publish In MSG Route")

  response, status, err := pubnubHandler.Publish().
    Channel("hello_msg_pubnub").Message(msg).Execute()

  if err != nil {
     // Request processing failed.
     // Handle message publish error
  }
  _, _ = response, status
  // fmt.Println(response, status, err)
}

func publishToGRPCWorld(w http.ResponseWriter, r *http.Request) {
  var conn *grpc.ClientConn
  port := os.Getenv("grpc_port")

  conn, err := grpc.Dial(fmt.Sprintf(":%s", port), grpc.WithInsecure())
  if err != nil {
    fmt.Println("did not connect: ", err)
  }
  defer conn.Close()

  sender_unique, conv_err := strconv.Atoi(os.Getenv("grpc_send_unique"))
  if conv_err != nil {
    fmt.Println("Conversion Error from env var, using Default sender value")
    sender_unique = 80
  }

  client := chatgrpc.NewChatterInterfaceClient(conn)
  sampleSenderMessage := chatgrpc.Msg{Msg: "Hey, sending to sender", Timestamp: "none", UniqueNum: int32(sender_unique)}

  resp, err := client.ChatSender(context.Background(), &sampleSenderMessage)
  if err == nil {
    // type conversions are very strict
    if resp.UniqueNum == int32(sender_unique) {
      fmt.Println("Verified Response: ", resp)
    } else {
      fmt.Println("Unacceptable Response: ", resp)
    }
  } else {
    fmt.Println(err)
  }
}

func publishToGRPCMsg(w http.ResponseWriter, r *http.Request) {
  var conn *grpc.ClientConn
  port := os.Getenv("grpc_port")

  conn, err := grpc.Dial(fmt.Sprintf(":%s", port), grpc.WithInsecure())
  if err != nil {
    fmt.Println("did not connect: ", err)
  }
  defer conn.Close()

  listener_unique, conv_err := strconv.Atoi(os.Getenv("grpc_listen_unique"))
  if conv_err != nil {
    fmt.Println("Conversion Error from env var, using Default listener value")
    listener_unique = 90
  }

  client := chatgrpc.NewChatterInterfaceClient(conn)
  sampleSenderMessage := chatgrpc.Msg{Msg: "Hey, sending to listener", Timestamp: "none", UniqueNum: int32(listener_unique)}

  resp, err := client.ChatListener(context.Background(), &sampleSenderMessage)
  if err == nil {
    if resp.UniqueNum == int32(listener_unique) {
      fmt.Println("Verified Response: ", resp)
    } else {
      fmt.Println("Unacceptable Response: ", resp)
    }
  } else {
    fmt.Println(err)
  }
}
