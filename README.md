# go-user-service

This module manages the storage of the **refresh tokens**.

![alt text](diagram.png)

> Requires a valid refresh token, **cannot generate** one itself.

### Before all
* Have Go
* Have gRPC
* Have Docker

installed.

### gRPC settings
Everything starts with designing the [protobuf](https://github.com/google/protobuf) file
(`./proto/token.proto` for this project) and once the messages and services are designed
(or changed later) the following command should be run:
```sh
cd $GOPATH/src/github.com/vahdet/go-auth-service
protoc -I proto proto/auth.proto --go_out=plugins=grpc:proto
```

### Calling the other services from the code
We need to make gRPC calls from the **code** to containers stemming from `go-user-store` and `go-refresh-token-store` images. To do that we can use Kubernetes client API

This command will generate (or recreate if it exists) a `.pb.go` file just beside itself.
It is `./proto/auth.pb.go` here and it allows implementing Go services, repositories, in short any Go code
behind the services defined in the `.proto` file.

For this case, the server implementation is performed in `./grpcserver/server.go` file.

### Git Tips for humble beings
The [StackOverflow answer](https://stackoverflow.com/a/23328996/4636715) resembles three lines of general purpose git commands
that can be used anytime a change is made and should be committed to `master` branch:

```bash
git add .
git commit -a -m "My classical message to be replaced"
git push
```
# Create Swagger Documentation from source code

    swagger generate spec -o ./swagger.json
