/*
	New gprc service started here for postService

	reference: https://tutorialedge.net/golang/go-grpc-beginners-tutorial/
*/

package main

import (
	"log"
	"net"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/post/post_pb"
	post_methods "github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/post/post_repository_storage"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Post Server started at Port 9100 ")
	post_listen, err := net.Listen("tcp", ":9100")

	if err != nil {
		log.Fatalf("Failed to listen to post server: %v", err)
	}

	postServer := grpc.NewServer()
	post_pb.RegisterPostsServiceServer(postServer, &post_methods.Server{})
	post_methods.InitializePostService()

	if err := postServer.Serve(post_listen); err != nil {
		log.Fatalf("Failed to server post server: %v", err)
	}
}
