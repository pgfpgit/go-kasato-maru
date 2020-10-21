# Go-kasato-maru

<div style="text-align:center">
<img src="https://github.com/ashleymcnamara/gophers/blob/master/GOPHER_SAILOR_STRIPE.png?raw=true" width="200" height="200">
</div>

`go-kasato-maru` is a template that helps you to create your new microservice architecture project. It uses go-kit/[truss](https://github.com/metaverse/truss) library.

## Prerequisites

### Truss

If you are going to create a new microservice then you need to install the truss library/cli. Please follow the truss install in https://github.com/metaverse/truss

### Docker

You need to have docker installed on your computer - Check the official install: https://docs.docker.com/get-docker/

## Container Architecture

See docker-compose.yml file

services:

- krakend
  - API gateway
- web
  - create-react-app
  - Served on Nginx alphine
- datastore
  - MongoDB
- todo
  - Example of a microservice created using truss;

## Run

Clone this repository

```
git clone git@github.com:rokoga/go-kasato-maru.git
```

then run docker-compose:

```
docker-compose up
```

## Usage

### Create a new microservice

1. Create a new directory on microservice

```
mkdir -p microservices/echo-service
```

2. Create a .proto file

```
touch microservices/echo-service/echo.proto
```

3. Edit the .proto (example):

```protobuf
// In general, while you can use proto2 (the current default protocol buffers
// version), we recommend that you use proto3 with gRPC as it lets you use the
// full range of gRPC-supported languages, as well as avoiding compatibility
// issues with proto2 clients talking to proto3 servers and vice versa.
syntax = "proto3";

// The package name determines the name of the directories that truss creates
// for `package echo;` truss will create the directory "echo-service".
package echo;

import "github.com/metaverse/truss/deftree/googlethirdparty/annotations.proto";

service Echo {
  // Echo "echos" the incoming string
  rpc Echo (EchoRequest) returns (EchoResponse) {
    option (google.api.http) = {
      // All fields (In) are query parameters of the http request unless otherwise specified
      get: "/echo"

      additional_bindings {
        // Trailing slashes are different routes
        get: "/echo/"
      }
    };
  }

  // Louder "echos" the incoming string with `Loudness` additional exclamation marks
  rpc Louder (LouderRequest) returns (EchoResponse) {
    option (google.api.http) = {
      custom {
        kind: "HEAD"
        // Loudness is accepted in the http path
        path: "/louder/{Loudness}"
      }
      additional_bindings {
        post: "/louder/{Loudness}"
        // All other fields (In) are located in the body of the http/json request
        body: "*"
      }
    };
  }

  // LouderGet is the same as Louder, but pulls fields other than Loudness (i.e. In) from query params instead of POST
  rpc LouderGet (LouderRequest) returns (EchoResponse) {
    option (google.api.http) = {
      // Loudness is accepted in the http path
      get: "/louder/{Loudness}"
    };
  }
}

message EchoRequest {
  string In = 1;
}

message LouderRequest {
  // In is the string to echo back
  string In = 1;
  // Loudness is the number of exclamations marks to add to the echoed string
  int32 Loudness = 2;
}

message EchoResponse {
  string Out = 1;
}
```

The name of the generated service will be based on the package name.

see more at: https://developers.google.com/protocol-buffers/docs/proto3 and https://github.com/metaverse/truss/blob/master/TUTORIAL.md

4. Run truss

```
truss echo.proto
```

It will be generated the service folder and this content:

```
.
├── echo-service
|   ├── cmd
|   │   └── echo
|   │       └── main.go
|   ├── echo.pb.go
|   ├── handlers
|   │   ├── handlers.go
|   │   ├── hooks.go
|   │   └── middlewares.go
|   └── svc
|       └── ...
└── echo.proto
```

5. Implement the business logic

You only need to edit the handlers.go functions with the business logic:

```go
func (s echoService) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {

	var resp pb.EchoResponse
	resp = pb.EchoResponse{
	// Out:
	}

    return &resp, nil
}
```

6. Run your service

```
go run echo-service/cmd/echo/main.go
```

You will see something like this:

```
ts=2020-10-20T23:25:14Z caller=main.go:106 transport=HTTP addr=:5050
ts=2020-10-20T23:25:14Z caller=main.go:98 transport=debug addr=:5060
ts=2020-10-20T23:25:14Z caller=main.go:124 transport=gRPC addr=:5040
```

Now you can acess your microservice with some HTTP requests on 5050

7. Create Dockerfile

   1. Copy and edit the `todo` dockerfile example to run the microservice on a container;
   2. Configure docker-compose.yaml to add the new service;

## Contributing

Pull requests are welcome.

## Special thanks

- [Truss](https://github.com/metaverse/truss/blob/master/TUTORIAL.md)

## License

[MIT](https://choosealicense.com/licenses/mit/)
