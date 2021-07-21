### Description
  * This app, for now, simply exposes 3 endpoints, 1 for HTTP and 2 for publishing to pubnub.
  * Basic web app in Golang which connects to pubnub.
    - it's recommended to use less files, and if files are using from other files, create project in GOROOT for easy import via package name.
    - SDK for pubnub/go is [here](https://github.com/pubnub/go)
  * Create an account on pubnub and get the free API keys.
    - store them in a .env file, use the godotenv package (node's dotenv) to load the keys in "os"

#### Setup, Execution
  * Just download and do "go run main.go"
  * Definitely create a .env file which has 3 key-values for the pubnub keys.
  * curl -X GET localhost:8080/[world - for publishing to a channel, msg - for publishing to another channel, hello - some http message response]

#### Models, Routes
  * Just the above 3 routes for now.
