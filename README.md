### Description and Stuff
  * Basic web app in Golang which connects to pubnub.
    - it's recommended to use less files, and if files are using from other files, create project in GOROOT for easy import via package name.
    - SDK for pubnub/go is [here](https://github.com/pubnub/go)
  * This app, for now, simply exposes 5 endpoints, 1 for HTTP, 2 for publishing to pubnub and 2 for dialing to gRPC.
  * Create an account on pubnub and get the free API keys.
    - store them in a .env file, use the godotenv package (node's dotenv) to load the keys in "os"
  * Adding multiple communication channels to this application
    - Pubnub - it might be using any protocol internally, ranging from websockets, MQTT, etc - depending on the product
    - Websocket - yet to add, from [here](https://github.com/gorilla/websocket/tree/master/examples/autobahn)
    - gRPC - basic client message sending - can be easily extended to send custom messages to the server
    - HTTP - this is basically working as a client now - gRPC/pubnub clients are created for sending messages in HTTP endpoints
    - HTTP/2 - yet to add
  *  More notes on protocols - HTTP, gRPC, Pubnub (whatever it uses internally), Web Socket, HTTP/2
    - Typically, the websocket is used for pushing from server to client - from client to server, we've HTTP - but can use the same channel
    - HTTP/2 and gRPC are also dual channel so can be used just like websockets. gRPC isn't implemented by any browsers though so only option is cmd clients.
    - Try to benchmark performances of various protocols?
  * Using gRPC's protocol buffer
    - In every type of communication (hardware/software), data packets and their organization is always present obviously. Sender would prepare some data to be sent,
      and the receiver will then read from that. Although how exactly this data is passed remains mysterious if we're working only via HTTP protocols (eg, REST APIs).
    - As already mentioned, the closer the communicating components are, the more prevalent this mechanism becomes when studying that communication.
    - Unix domain sockets can be like key-value pairs, where one process can store the data it wants with some kind of id for the processes that can consume it. Then,
      the other process can consume that data later - this is facilitated through OS??
    - Unfortunately, with growing distances, one must face the following issues - 
      1. Data sizes need to be preferably small - other layers for the communication protocol will add their own metadata - efficient compression.
      2. It should be easy to specify the structure of data, if it exists, and then easy to fill that structure before sending.
      3. Similarly, it should be easy to access this filled structure for the reader.
      4. Compression and decompression should be fast, apart from the readability being simpler.
      5. The structure should be such that any language can read it, efficiently - JSON, XML are present everywhere.
    - It should be noted that (2) and (3) can be achieved via existing features of languages rather than external libraries when communication is being done
      between applications of the same language? - since most languages have classes, instead of using JSON/XML etc, we can create classes which only have instance
      variables, and instance methods are concerned with getting/setting those instance variables. Now all applications just need to have this class defined in a file
      and use the methods to read/write to this data to be shared.
    - Now since classes are present everywhere, we can then create a specification that every language can process and create those classes automatically, rather
      than we having to write those classes in every application.
    - protocol buffer is basically the implementation of this theoretical specification mentioned above - one can specify the minimal amount of configurations for
      setting up a structure, and then run that specification in any language to generate classes - and specification is uniform so that any lanugage can create
      classes out of it while writing replicating specification across language is very simple.
    - Unfortunately, one will still have to import an external library which will help the language process that specification to create classes.
    - But if one is using gRPC, built on top of HTTP/2 as a full duplex communication protocol which is also fast, they can use this kind of class generator
      instead of JSON/XML - it is providing performance boost + it's also almost uniform + it's got some more specifics in the specification which can then help
      improve (1) and (4), ie, maybe they're storing only one instance of a repeated variable by storing just the other indices where the data is initially stored,
      etc to reduce size and improve compression.
    - Syntax -
      1. message - set of variable_names
         - datatype - (a) large list of basic datatypes, (b) can have enum variables, (c) can have other messages [like mongoose ORM specifications], (d) maps
         - unique index within that message structure - this isn't the value since we're talking about specification
         - field rules - (a) required, (b) optional - can have default values, (c) repeated
         - reserve - for future compatibility, make sure no one can change certain indices/variable_names
         - int32/int kind of type conversions have to be very carefully handled since golang is strongly typed
      2. messages can be nested. specifications can be imported from other specification files.
      3. Huge set of rules for updating message structure to maintain comptability.
      4. Extensions - reserve indices to be extended by other users (assuming the message is created via multiple modules of the repo)
      5. OneOf - used when multiple optional types specified and only one of them needs to be present in the final data.
      6. Maps supported with limitation. Packages can be specified at the top to prevent name clashes (specific to golang).
      7. services - in gRPC, these packed data will typically be used as req/response - this can be specified too, ie, which message which serve as the req/resp for
         which RPC handler function via <rpc Handler1 [RequestMessage] returns [ResponseMessage] {}> (golang specific).
      8. Generating the actual code - all of the above specification will generate a bunch of Interfaces (for client/server with methods specified in Services)

#### Setup, Execution
  * Just download and do "go run main.go" - let's say this is done in Console.
  * Definitely create a .env file which has 3 key-values for the pubnub keys. and some other key values for http/gRPC ports, unique nums, etc
  * curl -X GET localhost:8080 - and any of the below endpoints
    - /hello - some dummy http response
    - /pubnub/world - publishes to a channel "hello\_world\_pubnub" - a listener will respond to this on Console
    - /pubnub/msg - publishes to a channel "hello\_msg\_pubnub" - a listener will respond to this on Console
    - /grpc/world - dials to the gRPC server and calls the service "ChatSender" - gRPC server will respond accordingly on Console
    - /grpc/msg - dials to the gRPC server and calls the service "ChatListener" - gRPC server will respond accordingly on Console

#### Models, Routes
  * Just the above 5 routes for now.
