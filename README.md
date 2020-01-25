# GoMs Singleton

This go framework is still under development and this mention will be withdrawn when I consider that the project has sufficient maturity to start being exploited by other people.

### Why this project?

This project is based on GoMs, it serves as a repository to quickly set up a project, it will be very easy for you to post endpoints or create your services. It also serves to highlight that you don't need a GoMs git clone to use it.
You will find a project base allowing you to start quickly, I will add tooling as and when I need it.


### How to get started

You can start with git clone the skeleton that will serve as your base.
```bash
go get github.com/reversTeam/go-ms-skeleton
```

Then you can go to the folder, then launch the program to be sure that everything works well
```bash
go run $GOPATH/src/github.com/reversTeam/go-ms-skeleton/main.go
2020/01/25 01:32:06 [SYSTEM]: System is ready for catch exit's signals, To exit press CTRL+C
2020/01/25 01:32:06 [EXPORTER] Start listen on 127.0.0.1:4242
2020/01/25 01:32:06 [GRPC] Server listen on 127.0.0.1:42001
2020/01/25 01:32:06 [GRPC] Register service ping
2020/01/25 01:32:06 [HTTP] Register service ping
2020/01/25 01:32:06 [HTTP] Server listen on 127.0.0.1:8080
```
Then open another terminal and try to join the `ping` service
```bash
curl -sX GET "127.0.0.1:8080/ping"
{"Message":"Ping List"}
curl -sX POST "127.0.0.1:8080/ping"
{"Message":"Ping Create"}
curl -sX PATCH "127.0.0.1:8080/ping/42"
{"Message":"Ping View"}
curl -sX PATCH "127.0.0.1:8080/ping/42"
{"Message":"Ping Update"}
curl -sX DELETE "127.0.0.1:8080/ping/42"
{"Message":"Ping Delete"}
```

You should then have these answers on each of your endpoint is that it works well for you. You can now resume the server you started earlier and quit it using [Ctrl + c]

```bash
^C                                                              // [Ctrl + c] for exit the program 
2020/01/25 01:32:07 [SYSTEM]: Signal catch: interrupt
2020/01/25 01:32:07 [GRPC] Graceful Stop
2020/01/25 01:32:07 [HTTP] Graceful Stop
2020/01/25 01:32:07 [EXPORTER] Graceful Stop
2020/01/25 01:32:07 [HTTP] Error listen:  http: Server closed
```


### Create your own service

The best practice in the context of creating a service has gone through another repository, to do this create a git rest and follow the instructions of go-ms-service-singleton.
For a question of standardization of the documentation we recommend that you follow our guidelines, but the only limits are those of your imaginations.

As an example, in the case of a micro project, monorepos, it is normal to consider using local service at your repository.

Go to your project, for our example, and export the variable that allows you to reference modules locally
```bash
cd $GOPATH/src/github.com/reversTeam/go-ms-skeleton/
export GO111MODULE=on
```

Now create our first service, create the structure of your service:
```bash
mkdir -p services/example/protobuf
touch services/example/protobuf/example.proto services/example/service.go
```

You should therefore have this structure for your service:
```
services/example
├── protobuf
│   └── example.proto
└── service.go
```

We will start by filling out the example.proto file which will allow you to describe your endpoints:
```golang
syntax = "proto3";

package go.micro.service.example;

import "google/api/annotations.proto";
import "google/protobuf/empty.proto";
import "github.com/reversTeam/go-ms/services/goms/protobuf/goms.proto";

service Example {

	rpc List(google.protobuf.Empty) returns (goms.GoMsResponse) {
		option (google.api.http) = {
			get: "/example"
		};
	}

	rpc Create(google.protobuf.Empty) returns (goms.GoMsResponse) {
		option (google.api.http) = {
			post: "/example"
			body: "*"
		};
	}

	rpc Get(goms.GoMsEntityRequest) returns (goms.GoMsResponse) {
		option (google.api.http) = {
			get: "/example/{id}"
		};
	}

	rpc Update(goms.GoMsEntityRequest) returns (goms.GoMsResponse) {
		option (google.api.http) = {
			patch: "/example/{id}"
			body: "*"
		};
	}

	rpc Delete(goms.GoMsEntityRequest) returns (goms.GoMsResponse) {
		option (google.api.http) = {
			delete: "/example/{id}"
		};
	}
}

```

We do there, the description of several endpoint grpc and http for your service.
| GRPC   | HTTP VERBE | HTTP PATH     |
|:-------|:-----------|:--------------|
| List   | GET        | /example      |
| Create | POST       | /example      |
| Get    | GET        | /example/{id} |
| Update | PATCH      | /example/{id} |
| Delete | DELETE     | /example/{id} |



Now that we have finished the description of your proto, you can now start the generation of protobuf with the following command:
```bash
make protogen
```

If you have an error at this step, you may be missing one of the following dependencies:
```bash
brew install protobuf
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

As for launching the `make lint` command, you will need to install:
```bash
brew install golangci/tap/golangci-lint
```


Your file nomenclature should normally be this one now, we just generated the http handler, the entity structure and the swagger documentation.
```
services/example
├── protobuf
│   ├── exmaple.pb.go
│   ├── exmaple.pb.gw.go
│   ├── exmaple.proto
│   └── exmaple.swagger.json
└── service.go
```

The description of each of its endpoints can be found in the `service.go` file.
```golang
package example

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/reversTeam/go-ms/core"
	pb "github.com/reversTeam/go-ms-skeleton/services/example/protobuf"
	"github.com/reversTeam/go-ms/services/goms"
	ms "github.com/reversTeam/go-ms/services/goms/protobuf"
	"golang.org/x/net/context"
)

// Define the service structure
type ExampleService struct {
	*goms.GoMsService
}

// Instanciate the service without dependency because it's role of ServiceFactory
func NewService(name string) *ExampleService {
	s := &ExampleService{
		GoMsService: goms.NewService(name),
	}

	return s
}

// This method is required for redister your service on the Http server
func (o *ExampleService) RegisterHttp(gh *core.GoMsHttpServer, endpoint string) error {
	return pb.RegisterExampleHandlerFromEndpoint(gh.Ctx, gh.Mux, endpoint, gh.Grpc.Opts)
}

// This method is required for redister your service on the Grpc server
func (o *ExampleService) RegisterGrpc(gs *core.GoMsGrpcServer) {
	pb.RegisterExampleServer(gs.Server, o)
}

// Endpoint :
//  - grpc : List
//  - http : Get /example
func (o *ExampleService) List(ctx context.Context, in *empty.Empty) (*ms.GoMsResponse, error) {
	return &ms.GoMsResponse{
		Message: "Example List",
	}, nil
}

// Endpoint :
//  - grpc : Create
//  - http : POST /example
func (o *ExampleService) Create(ctx context.Context, in *empty.Empty) (*ms.GoMsResponse, error) {
	return &ms.GoMsResponse{
		Message: "Example Create",
	}, nil
}

// Endpoint :
//  - grpc : Get
//  - http : GET /example/{id}
func (o *ExampleService) Get(ctx context.Context, in *ms.GoMsEntityRequest) (*ms.GoMsResponse, error) {
	return &ms.GoMsResponse{
		Message: "Example View",
	}, nil
}

// Endpoint :
//  - grpc : Update
//  - http : PATCH /example/{id}
func (o *ExampleService) Update(ctx context.Context, in *ms.GoMsEntityRequest) (*ms.GoMsResponse, error) {
	return &ms.GoMsResponse{
		Message: "Example Update",
	}, nil
}

// Endpoint :
//  - grpc : Delete
//  - http : PATCH /example/{id}
func (o *ExampleService) Delete(ctx context.Context, in *ms.GoMsEntityRequest) (*ms.GoMsResponse, error) {
	return &ms.GoMsResponse{
		Message: "Example Delete",
	}, nil
}

```

The `RegisterHttp` and` RegisterGrpc` methods are methods that allow you to expose your services, you will always have to implement them.

Now it only remains to instantiate your service in main.go and add it to the servers (HTTP & GRPC).

```golang
package main

import (
	"flag"
	"github.com/reversTeam/go-ms/core"
	"github.com/reversTeam/go-ms-skeleton/services/example"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	// Default flag values for GRPC server
	GRPC_DEFAULT_HOST = "127.0.0.1"
	GRPC_DEFAULT_PORT = 42001

	// Default flag values for http server
	HTTP_DEFAULT_HOST = "127.0.0.1"
	HTTP_DEFAULT_PORT = 8080
)

var (
	// flags for Grpc server
	grpcHost = flag.String("grpc-host", GRPC_DEFAULT_HOST, "Grpc listening host")
	grpcPort = flag.Int("grpc-port", GRPC_DEFAULT_PORT, "Grpc listening port")

	// flags for http server
	httpHost = flag.String("http-host", HTTP_DEFAULT_HOST, "http gateway host")
	httpPort = flag.Int("http-port", HTTP_DEFAULT_PORT, "http gateway port")
)

func main() {
	// Instantiate context in background
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Parse flags
	flag.Parse()

	// Create a gateway configuration
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	// setup servers
	grpcServer := core.NewGoMsGrpcServer(ctx, *grpcHost, *grpcPort, opts)
	httpServer := core.NewGoMsHttpServer(ctx, *httpHost, *httpPort, grpcServer)

	// setup services
	exampleService := goms.NewService("example")

	// Register service to the grpc server
	grpcServer.AddService(gomsService)

	// Register service to the http server
	httpServer.AddService(gomsService)

	// Graceful stop servers
	core.AddServerGracefulStop(grpcServer)
	core.AddServerGracefulStop(httpServer)
	// Catch ctrl + c
	done := core.CatchStopSignals()

	// Start Grpc Server
	err := grpcServer.Start()
	if err != nil {
		log.Fatal("An error occured, the grpc server can be running", err)
	}
	// Start Http Server
	err = httpServer.Start()
	if err != nil {
		log.Fatal("An error occured, the http server can be running", err)
	}

	<-done
}
```

Try yours endpoint:
```bash
curl -sX GET "127.0.0.1:8080/example"
{"Message":"Example List"}
curl -sX POST "127.0.0.1:8080/example"
{"Message":"Example Create"}
curl -sX PATCH "127.0.0.1:8080/example/42"
{"Message":"Example View"}
curl -sX PATCH "127.0.0.1:8080/example/42"
{"Message":"Example Update"}
curl -sX DELETE "127.0.0.1:8080/example/42"
{"Message":"Example Delete"}
```


### More defails about the main.go

If we break down the code we have above we can see that we have different phases:
 - Import of libraries which are necessary to operate your hand
```golang
import (
	"flag"
	"github.com/reversTeam/go-ms/core"
	"github.com/reversTeam/go-ms/services/goms"   // Only for example
	"github.com/reversTeam/go-ms/services/child"  // Only for example
	// "github.com/yoursName/go-ms-service-what-you-want"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)
 ```
 - Initialization of the constants for the default values ​​of the flags, readability questions
 ```golang
 const (
	// Default flag values for GRPC server
	GRPC_DEFAULT_HOST = "127.0.0.1"
	GRPC_DEFAULT_PORT = 42001

	// Default flag values for http server
	HTTP_DEFAULT_HOST = "127.0.0.1"
	HTTP_DEFAULT_PORT = 8080
)
 ```
 - Initialization of program flags in global variables, questionable but extremely readable
```golang
var (
	// flags for Grpc server
	grpcHost = flag.String("grpc-host", GRPC_DEFAULT_HOST, "Grpc listening host")
	grpcPort = flag.Int("grpc-port", GRPC_DEFAULT_PORT, "Grpc listening port")

	// flags for http server
	httpHost = flag.String("http-host", HTTP_DEFAULT_HOST, "http gateway host")
	httpPort = flag.Int("http-port", HTTP_DEFAULT_PORT, "http gateway port")
)
```
 - Initialization of grpc and http servers
```golang
// Instantiate context in background
ctx := context.Background()
ctx, cancel := context.WithCancel(ctx)
defer cancel()

// Parse flags
flag.Parse()

// Create a gateway configuration
opts := []grpc.DialOption{
	grpc.WithInsecure(),
}

// setup servers
grpcServer := core.NewGoMsGrpcServer(ctx, *grpcHost, *grpcPort, opts)
httpServer := core.NewGoMsHttpServer(ctx, *httpHost, *httpPort, grpcServer)
```
 - Service initialization
   If you create your own modules try to respect this name as well as possible for your repositories, I would try afterwards to make a service manager that everyone can offer their own services.
```golang
// setup services

gomsService := goms.NewService("goms")    // import "github.com/reversTeam/go-ms/services/goms"
childService := child.NewService("child") // import "github.com/reversTeam/go-ms/services/child"
whatYouWantService := whatYouWant.NewService("what-you-want") // import "github.com/yoursName/go-ms-service-what-you-want"
```
 - Ajout des services sur les différents serveurs
```golang
// Register service to the grpc server
grpcServer.AddService(gomsService)
grpcServer.AddService(childService)

// Register service to the http server
httpServer.AddService(gomsService)
httpServer.AddService(childService)
```
 - Ajout des signaux pour couper les services
```golang
// Graceful stop servers
core.AddServerGracefulStop(grpcServer)
core.AddServerGracefulStop(httpServer)
// Catch ctrl + c
done := core.CatchStopSignals()
```
 - Launch of different servers
 If you want to start only one of the two servers, delete the code that starts the one you don't want. In case you launch an http server, you will have to give it the configuration of a functional grpc server.
```golang
// Start Grpc Server
err := grpcServer.Start()
if err != nil {
	log.Fatal("An error occured, the grpc server can be running", err)
}
// Start Http Server
err = httpServer.Start()
if err != nil {
	log.Fatal("An error occured, the http server can be running", err)
}
```
 - We are waiting for the signal telling us to finish the services, in the case of a ctrl + c for example
```golang
<-done
```

### Credits
 - golang
 - protoc
 - grpc-gateway