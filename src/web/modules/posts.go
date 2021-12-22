package modules

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/post/post_pb"
	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/user/user_pb"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userCookie, err := r.Cookie("user")
		if err != nil || userCookie.Value == "" {
			temp, err := template.ParseFiles("../../static/html/login.html")
			if err != nil {
				panic(err)
			}
			err = temp.Execute(w, nil)
			if err != nil {
				panic(err)
			}
			return
		}
		userIdNum, errorValue := strconv.Atoi(userCookie.Value)
		log.Println("user cookie: " + userCookie.Value)
		if errorValue != nil {
			panic(errorValue)
		}
		text := r.FormValue("post_text")
		post := post_pb.Post{
			UserID: int32(userIdNum), PostText: text,
		}
		postList, err := clientObj.PostClient.AddPost(context.Background(), &post)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		log.Println("User post list: ", postList)
		temp, err := template.ParseFiles("../../static/html/createpost.html")
		if err != nil {
			panic(err)
		}
		err = temp.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	} else {
		temp, err := template.ParseFiles("../../static/html/createpost.html")
		if err != nil {
			panic(err)
		}
		err = temp.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	}
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	userCookie, err := r.Cookie("user")
	if err != nil || userCookie.Value == "" {
		temp, err := template.ParseFiles("../../static/html/login.html")
		if err != nil {
			panic(err)
		}
		err = temp.Execute(w, nil)
		if err != nil {
			panic(err)
		}
		return
	}
	userIdNum, errorValue := strconv.Atoi(userCookie.Value)
	if errorValue != nil {
		panic(errorValue)
	}
	log.Println("r method: " + r.Method)
	if r.Method == "GET" {
		currUserIdNum := user_pb.UserStruct{
			IdNum: int32(userIdNum),
		}

		currentUser, errVal := clientObj.UserClient.GetUserFromUserId(context.Background(), &currUserIdNum)
		if errVal != nil {
			log.Fatal("Error: ", err)
		}
		log.Println("Login user following ", currentUser.User.Following)
		var followeesUserIDArray []int32
		// add user id of followees in an array
		for _, user := range currentUser.User.Following {
			followeeUserName := user_pb.UserStruct{
				UserName: user,
			}
			followeeUser, _ := clientObj.UserClient.GetUserFromUserName(context.Background(), &followeeUserName)
			followeesUserIDArray = append(followeesUserIDArray, followeeUser.User.IdNum)
		}

		getFolloweePosts := post_pb.PostPbStruct{
			FolloweeUserIDsList: followeesUserIDArray,
		}
		followeePosts, _ := clientObj.PostClient.GetFolloweePosts(context.Background(), &getFolloweePosts)

		type Post struct {
			FirstName string
			LastName  string
			PostText  string
		}
		type FolloweePostStruct struct {
			FolloweePostsArray []Post
		}

		FolloweePostsInfo := FolloweePostStruct{
			FolloweePostsArray: make([]Post, 0),
		}
		for _, followeePost := range followeePosts.PostList {
			followeeUserID := user_pb.UserStruct{
				IdNum: int32(followeePost.UserID),
			}
			followeeUserName, _ := clientObj.UserClient.GetUserNameByUserId(context.Background(), &followeeUserID)

			followeeUserNameStruct := user_pb.UserStruct{
				UserName: followeeUserName.UserName,
			}
			followeeUser, _ := clientObj.UserClient.GetUserFromUserName(context.Background(), &followeeUserNameStruct)
			post := Post{
				FirstName: followeeUser.User.FirstName,
				LastName:  followeeUser.User.LastName,
				PostText:  followeePost.PostText,
			}
			log.Println("Posts struct: ", post)
			FolloweePostsInfo.FolloweePostsArray = append(FolloweePostsInfo.FolloweePostsArray, post)
		}
		log.Println("Followee Post Array: ", FolloweePostsInfo.FolloweePostsArray)
		temp, _ := template.ParseFiles("../../static/html/feed.html")
		temp.Execute(w, FolloweePostsInfo)
	}
}
