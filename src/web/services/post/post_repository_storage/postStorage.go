package post_methods

import (
	"context"
	"log"
	"os"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/raftHelper"
	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/post/post_pb"
)

func (s *Server) AddPost(ctx context.Context, postInfo *post_pb.Post) (*post_pb.UserPostList, error) {
	mutex.Lock()
	post := &post_pb.Post{
		UserID:   postInfo.UserID,
		PostText: postInfo.PostText,
	}
	userPostListObj.userPostList.PostList = append(userPostListObj.userPostList.PostList, post)
	log.Println("User post list", userPostListObj.userPostList)
	postList := &post_pb.UserPostList{
		PostList: userPostListObj.userPostList.PostList,
	}
	_, errorValue := raftHelper.RequestForService("userPostList", userPostListObj.userPostList, "PUT")
	if errorValue != nil {
		log.Println("Error occurred while PUT Request in Post Metod: ", errorValue)
		os.Exit(1)
	}
	mutex.Unlock()
	return postList, nil
}
