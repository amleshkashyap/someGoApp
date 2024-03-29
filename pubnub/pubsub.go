package pubsub

import (
  "fmt"
  "os"
  "github.com/joho/godotenv"
  pubnub "github.com/pubnub/go"
)

// use upper case first letter of function name to export it - truly concise!
func SetupGlobalConfigs() *pubnub.Config {
  config := pubnub.NewConfig()
  config.SubscribeKey = os.Getenv("pubnub_subscribe")
  config.PublishKey = os.Getenv("pubnub_publish")
  config.SecretKey = os.Getenv("pubnub_secret")
  config.UUID = os.Getenv("pubnub_uuid")
  return config
}

var PubnubListener = func() {
  env_err := godotenv.Load(".env")
  if env_err != nil {
    panic(env_err.Error())
  }
  configs := SetupGlobalConfigs()
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
