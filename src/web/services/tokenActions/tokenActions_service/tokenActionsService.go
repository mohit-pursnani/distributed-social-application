/*
	New gprc service started here for tokenActions

	reference: https://tutorialedge.net/golang/go-grpc-beginners-tutorial/
*/

package main

import (
	"log"
	"net"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/tokenActions/tokenActions_pb"
	tokenActions_methods "github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/tokenActions/tokenActions_repository_storage"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Authentication Server started at Port 9000 ")
	tokenActions_listen, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen to auth server: %v", err)
	}

	tokenActionsServer := grpc.NewServer()
	tokenActions_pb.RegisterTokenActionsServiceServer(tokenActionsServer, &tokenActions_methods.Server{})
	tokenActions_methods.InitializeTokenActionsService()

	if err := tokenActionsServer.Serve(tokenActions_listen); err != nil {
		log.Fatalf("Failed to serve auth server: %v", err)
	}
}
