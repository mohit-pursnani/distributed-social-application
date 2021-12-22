/*
	Add posts to the post list data structure and get all the needed posts to serve on the landing page after the
	user has been followed.
*/

package post_methods

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"os"
	"sync"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/raftHelper"
	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/post/post_pb"
)

var mutex sync.Mutex

type Server struct{}

type userPostListStruct struct {
	userPostList *post_pb.UserPostList
}

var userPostListObj userPostListStruct

func InitializePostService() {
	mutex.Lock()
	userPostListValue := CreateUserPosts()
	userPostListObj.userPostList = userPostListValue
	mutex.Unlock()
	_, errorVal := raftHelper.RequestForService("userPostList", userPostListObj.userPostList, "PUT")
	if errorVal != nil {
		log.Println("Error occurred while PUT Request in Initialize Metod: ", errorVal)
		os.Exit(1)
	}
	log.Println("User post list: ", userPostListObj.userPostList)

}

func CreateUserPosts() *post_pb.UserPostList {

	emptyUserPostList := &post_pb.UserPostList{
		PostList: make([]*post_pb.Post, 0),
	}
	return emptyUserPostList
}

// func (s *Server) AddPost(ctx context.Context, postInfo *post_pb.Post) (*post_pb.UserPostList, error) {
// 	mutex.Lock()
// 	post := &post_pb.Post{
// 		UserID:   postInfo.UserID,
// 		PostText: postInfo.PostText,
// 	}
// 	userPostListObj.userPostList.PostList = append(userPostListObj.userPostList.PostList, post)
// 	log.Println("User post list", userPostListObj.userPostList)
// 	postList := &post_pb.UserPostList{
// 		PostList: userPostListObj.userPostList.PostList,
// 	}
// 	_, errorValue := raftHelper.RequestForService("userPostList", userPostListObj.userPostList, "PUT")
// 	if errorValue != nil {
// 		log.Println("Error occurred while PUT Request in Post Metod: ", errorValue)
// 		os.Exit(1)
// 	}
// 	mutex.Unlock()
// 	return postList, nil
// }

func (s *Server) GetFolloweePosts(ctx context.Context, userFolloweeList *post_pb.PostPbStruct) (*post_pb.UserPostList, error) {
	mutex.Lock()
	upl := CreateUserPosts()
	followeePostList := upl
	mutex.Unlock()
	responseInBytes, errorValue := raftHelper.RequestForService("userPostList", userPostListObj.userPostList, "GET")
	if errorValue != nil {
		log.Println("Error occurred while PUT Request in Post Metod: ", errorValue)
		os.Exit(1)
	}
	mutex.Lock()

	var userPostListFromStorage *post_pb.UserPostList

	decoder := gob.NewDecoder(bytes.NewBufferString(responseInBytes))
	decodedDataErr := decoder.Decode(&userPostListObj.userPostList)

	if decodedDataErr != nil {
		log.Println("error occurred while decoding: ", decodedDataErr)
		os.Exit(1)
		userPostListFromStorage, errorValue = nil, decodedDataErr
	} else {
		userPostListFromStorage, errorValue = userPostListObj.userPostList, nil
	}

	if errorValue != nil {
		log.Println("Error occurred while decoding in GET Request: ", errorValue)
		os.Exit(1)
	}

	for _, userPost := range userPostListFromStorage.PostList {
		for _, userID := range userFolloweeList.FolloweeUserIDsList {
			if userPost.UserID == userID {
				var postListObj = &post_pb.Post{UserID: userID, PostText: userPost.PostText}
				followeePostList.PostList = append(followeePostList.PostList, postListObj)
			}
		}
	}
	mutex.Unlock()
	return followeePostList, nil
}

func (s *Server) ClearAllPosts(ctx context.Context, postInfo *post_pb.Post) (*post_pb.PostPbStruct, error) {

	userPostListValue := CreateUserPosts()
	userPostListObj.userPostList = userPostListValue
	_, errorVal := raftHelper.RequestForService("userPostList", userPostListObj.userPostList, "PUT")
	if errorVal != nil {
		log.Println("Error occurred while PUT Request in Initialize Metod: ", errorVal)
		os.Exit(1)
	}

	postStatus := &post_pb.PostPbStruct{
		AddPostStatus: true,
	}
	return postStatus, nil
}

func (s *Server) GetAllPosts(ctx context.Context, postInfo *post_pb.Post) (*post_pb.UserPostList, error) {
	mutex.Lock()
	upl := CreateUserPosts()
	followeePostList := upl
	mutex.Unlock()
	responseInBytes, errorValue := raftHelper.RequestForService("userPostList", userPostListObj.userPostList, "GET")
	if errorValue != nil {
		log.Println("Error occurred while PUT Request in Post Metod: ", errorValue)
		os.Exit(1)
	}
	mutex.Lock()

	var userPostListFromStorage *post_pb.UserPostList

	decoder := gob.NewDecoder(bytes.NewBufferString(responseInBytes))
	decodedDataErr := decoder.Decode(&userPostListObj.userPostList)

	if decodedDataErr != nil {
		log.Println("error occurred while decoding: ", decodedDataErr)
		os.Exit(1)
		userPostListFromStorage, errorValue = nil, decodedDataErr
	} else {
		userPostListFromStorage, errorValue = userPostListObj.userPostList, nil
	}

	if errorValue != nil {
		log.Println("Error occurred while decoding in GET Request: ", errorValue)
		os.Exit(1)
	}

	for _, userPost := range userPostListFromStorage.PostList {
		var postListObj = &post_pb.Post{UserID: 9999, PostText: userPost.PostText}
		followeePostList.PostList = append(followeePostList.PostList, postListObj)
	}

	postList := &post_pb.UserPostList{
		PostList: userPostListObj.userPostList.PostList,
	}

	mutex.Unlock()
	return postList, nil
}
