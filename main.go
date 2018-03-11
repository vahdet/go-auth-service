package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/vahdet/go-auth-service/http/controllers"
	"github.com/vahdet/go-auth-service/services"
	"github.com/vahdet/go-auth-service/services/interfaces"
	reftknstrpb "github.com/vahdet/go-refresh-token-store-redis/proto"
	usrstrpb "github.com/vahdet/go-user-store-redis/proto"
	"google.golang.org/grpc"
)

// See the Kubernetes .yml for the value of the environment variables
const (
	USER_STORE_SERVICE_ENV_VAR_NAME          = "USER_STORE_SERVICE_URL"
	REFRESH_TOKEN_STORE_SERVICE_ENV_VAR_NAME = "REFRESH_TOKEN_STORE_SERVICE_URL"
)

func main() {

	// Set up a connections to the grpc servers.
	usrStrConn, err := grpc.Dial(os.Getenv(USER_STORE_SERVICE_ENV_VAR_NAME), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer usrStrConn.Close()

	rfrTknStrConn, err := grpc.Dial(os.Getenv(REFRESH_TOKEN_STORE_SERVICE_ENV_VAR_NAME), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer rfrTknStrConn.Close()

	var srv interfaces.AuthService
	srv = services.NewAuthService(usrstrpb.NewUserServiceClient(usrStrConn), reftknstrpb.NewTokenServiceClient(rfrTknStrConn))

	// grpc client
	router := httprouter.New()
	ac := controllers.NewAuthController(srv)
	router.POST("/register", ac.Register)
	router.POST("/login/:name", ac.Login)
	router.POST("/logout", ac.Logout)

	log.Fatal(http.ListenAndServe(":8080", router))

}
