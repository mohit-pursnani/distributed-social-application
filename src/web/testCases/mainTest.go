package main

import (
	"log"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/post/post_pb"
	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/tokenActions/tokenActions_pb"
	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/user/user_pb"
	"google.golang.org/grpc"
)

type Clients struct {
	TokenClient tokenActions_pb.TokenActionsServiceClient
	PostClient  post_pb.PostsServiceClient
	UserClient  user_pb.UserServiceClient
}

type ClientPorts struct {
	AuthPort string
	PostPort string
	UserPort string
}

var clientObj Clients

var clientPortObj ClientPorts

func InitializeClientPorts() {
	clientPortObj.AuthPort = ":9000"
	clientPortObj.PostPort = ":9100"
	clientPortObj.UserPort = ":9200"
}

func DialToServers() {

	var authenticationConnection *grpc.ClientConn
	var postConnection *grpc.ClientConn
	var userConnection *grpc.ClientConn

	authenticationConnection, _ = grpc.Dial(clientPortObj.AuthPort, grpc.WithInsecure())
	postConnection, _ = grpc.Dial(clientPortObj.PostPort, grpc.WithInsecure())
	userConnection, _ = grpc.Dial(clientPortObj.UserPort, grpc.WithInsecure())

	clientObj.TokenClient = tokenActions_pb.NewTokenActionsServiceClient(authenticationConnection)
	clientObj.PostClient = post_pb.NewPostsServiceClient(postConnection)
	clientObj.UserClient = user_pb.NewUserServiceClient(userConnection)
}

func Initialize() {
	// initialize client port
	InitializeClientPorts()
	// initialize client connections
	DialToServers()
}

func main() {
	log.Println("Running Tests...")
	Initialize()

	//Running the regsiter tests
	RegisterTest()

	//Running the login tests
	LoginTest()

	//Running the follow tests
	FollowTest()

	//Running the unfollow tests
	UnFollowTest()

	//Running Concurrent Batch post tests
	BatchPostTest()

	//Running concurrent Batch register tests
	BatchRegister()

	//Running concurrent Batch Login tests
	BatchLogin()
}
