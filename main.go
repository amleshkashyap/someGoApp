package main

import (
  "fmt"
  "os"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/joho/godotenv"
  pubnub "github.com/pubnub/go"
)

func newRouter() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/hello", handler).Methods("GET")
  r.HandleFunc("/world", publishToPubnubWorld).Methods("GET")
  r.HandleFunc("/msg", publishToPubnubMsg).Methods("GET")
  return r
}

func setupGlobalConfigs() *pubnub.Config {
  config := pubnub.NewConfig()
  config.SubscribeKey = os.Getenv("pubnub_subscribe")
  config.PublishKey = os.Getenv("pubnub_publish")
  config.SecretKey = os.Getenv("pubnub_secret")
  config.UUID = "bbf80b36-d536-4297-b956-44f21ef206a3"
  return config
}

var pubnubHandleIt = func() {
  configs := setupGlobalConfigs()
  pubnubHandler := pubnub.NewPubNub(configs)
  listener := pubnub.NewListener()
  // doneConnect := make(chan bool)
  // donePublish := make(chan bool)
  fmt.Println("Listener Is Created")

  go func() {
    for {
      select {
        case status := <-listener.Status:
          switch status.Category {
          case pubnub.PNDisconnectedCategory:
          // This event happens when radio / connectivity is lost
          case pubnub.PNConnectedCategory:
          // Connect event. You can do stuff like publish, and know you'll get it.
          // Or just use the connected event to confirm you are subscribed for
          // UI / internal notifications, etc
            // doneConnect <- true
          case pubnub.PNReconnectedCategory:
          // Happens as part of our regular operation. This event happens when
          // radio / connectivity is lost, then regained.
          }
          case message := <-listener.Message:
          // Handle new message stored in message.message
          if message.Channel == "hello_world_pubnub" {
            fmt.Println("Listener has heard your WORLDLY message")
	  } else if message.Channel == "hello_msg_pubnub" {
            fmt.Println("Listener has heard your MSGY message")
          } else {
          }
          // print the received message
          if msg, ok := message.Message.(map[string]interface{}); ok {
            fmt.Println(msg["msg"])
          }
          // log these - message { Message, Subscription, Timetoken }
            // donePublish <- true
        case <-listener.Presence:
        // handle presence
      }
    }
  }()

  pubnubHandler.AddListener(listener)

  // pubnubHandler is subscribed to hello_world_pubnub and hello_msg_pubnub channels - invoked via /world and /msg routes
  pubnubHandler.Subscribe().
      Channels([]string{"hello_world_pubnub", "hello_msg_pubnub"}).
      Execute()
  // <-doneConnect
}


func main() {
  env_err := godotenv.Load(".env")
  if env_err != nil {
    panic(env_err.Error())
  }
  fmt.Println("Starting the Pubnub Listener")
  pubnubHandleIt()
  // The router is now formed by calling the `newRouter` constructor function
  // that we defined above. The rest of the code stays the same
  r := newRouter()
  fmt.Println("Starting Server at localhost:8080")
  err := http.ListenAndServe(":8080", r)
  if err != nil {
    panic(err.Error())
  }
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World! From Golang")
}

func publishToPubnubWorld(w http.ResponseWriter, r *http.Request) {
  configs := setupGlobalConfigs()
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
  configs := setupGlobalConfigs()
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
