/*
	Use current user token to get information about the user (usedID), then convert userId to username to look up
	current user details in the user map (stored in the server side) then add to user's follow list if anothe friend
	user is followed. Remove from follow list if the user is un followed.
*/

package modules

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/user/user_pb"
)

func FollowUser(w http.ResponseWriter, r *http.Request) {
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

	if r.Method == "POST" {
		userIdNum, errorValue := strconv.Atoi(userCookie.Value)
		if errorValue != nil {
			panic(errorValue)
		}
		userNameToFollow := r.FormValue("user_name")
		currUserIdNum := user_pb.UserStruct{
			IdNum: int32(userIdNum),
		}

		currentUser, errVal := clientObj.UserClient.GetUserFromUserId(context.Background(), &currUserIdNum)
		if errVal != nil {
			log.Fatal("Error: ", err)
		}

		actionForFolloweeStruct := user_pb.ActionForFollowee{
			CurrentUserName:      currentUser.User.UserName,
			UserNameForOperation: userNameToFollow,
		}
		followStatus, errVal := clientObj.UserClient.FollowUser(context.Background(), &actionForFolloweeStruct)

		if errVal != nil {
			log.Fatal("Error: ", err)
		}
		log.Println("Follow Status: ", followStatus.Status)

		temp, errorValue := template.ParseFiles("../../static/html/followStatus.html")
		if errorValue != nil {
			panic(errorValue)
		}
		errorValue = temp.Execute(w, nil)
		if errorValue != nil {
			panic(errorValue)
		}

	} else {

		//List of users that the current user can potentially follow:
		userIdNum, errorValue := strconv.Atoi(userCookie.Value)
		if errorValue != nil {
			panic(errorValue)
		}
		currUserIdNum := user_pb.UserStruct{
			IdNum: int32(userIdNum),
		}
		currentUser, _ := clientObj.UserClient.GetUserFromUserId(context.Background(), &currUserIdNum)

		allUsers, _ := clientObj.UserClient.GetAllUsers(context.Background(), &currUserIdNum)

		var usersCurrCanFollow []string

		for userName := range allUsers.User {
			if currentUser.User.UserName == userName {
				continue
			}
			found := false
			for i := 0; i < len(currentUser.User.Following); i += 1 {
				if currentUser.User.Following[i] == userName {
					found = true
				}
			}

			if !found {
				usersCurrCanFollow = append(usersCurrCanFollow, userName)
			}
		}

		log.Println("Current User can follow GET: ")
		for i := 0; i < len(usersCurrCanFollow); i += 1 {
			log.Println(usersCurrCanFollow[i])
		}

		type UserFollowStruct struct {
			UserFollowList []string
		}

		UserFolloweeLists := UserFollowStruct{
			UserFollowList: make([]string, 0),
		}

		UserFolloweeLists.UserFollowList = usersCurrCanFollow
		temp, errorValue := template.ParseFiles("../../static/html/findfriends.html")
		if errorValue != nil {
			panic(errorValue)
		}
		errorValue = temp.Execute(w, UserFolloweeLists)
		if errorValue != nil {
			panic(errorValue)
		}
	}
}

func FolloweeList(w http.ResponseWriter, r *http.Request) {
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

	if r.Method == "GET" {
		//List of users that the current user can potentially follow:
		userIdNum, errorValue := strconv.Atoi(userCookie.Value)
		if errorValue != nil {
			panic(errorValue)
		}
		currUserIdNum := user_pb.UserStruct{
			IdNum: int32(userIdNum),
		}
		currentUser, _ := clientObj.UserClient.GetUserFromUserId(context.Background(), &currUserIdNum)

		log.Println("Following List: ", currentUser.User.Following)

		var usersCurrCanFollow []string
		usersCurrCanFollow = currentUser.User.Following

		for i := 0; i < len(currentUser.User.Following); i += 1 {
			if currentUser.User.Following[i] == currentUser.User.UserName {
				usersCurrCanFollow = append(usersCurrCanFollow[:i], usersCurrCanFollow[i+1:]...)
			}
		}

		log.Println("User followee List: ", usersCurrCanFollow)
		type UserFollowStruct struct {
			UserFollowList []string
		}

		UserFolloweeLists := UserFollowStruct{
			UserFollowList: make([]string, 0),
		}

		UserFolloweeLists.UserFollowList = usersCurrCanFollow
		temp, errorValue := template.ParseFiles("../../static/html/followeeList.html")
		if errorValue != nil {
			panic(errorValue)
		}
		errorValue = temp.Execute(w, UserFolloweeLists)
		if errorValue != nil {
			panic(errorValue)
		}
	}
}

func UnFollowUser(w http.ResponseWriter, r *http.Request) {
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
	if r.Method == "POST" {
		userIdNum, errorValue := strconv.Atoi(userCookie.Value)
		if errorValue != nil {
			panic(errorValue)
		}
		userNameToUnFollow := r.FormValue("user_name")
		currUserIdNum := user_pb.UserStruct{
			IdNum: int32(userIdNum),
		}

		currentUser, errVal := clientObj.UserClient.GetUserFromUserId(context.Background(), &currUserIdNum)
		if errVal != nil {
			log.Fatal("Error: ", err)
		}

		actionForFolloweeStruct := user_pb.ActionForFollowee{
			CurrentUserName:      currentUser.User.UserName,
			UserNameForOperation: userNameToUnFollow,
		}
		unFollowStatus, errVal := clientObj.UserClient.UnfollowUser(context.Background(), &actionForFolloweeStruct)
		if errVal != nil {
			log.Fatal("Error: ", err)
		}
		log.Println("UnFollow Status: ", unFollowStatus.Status)
		temp, errorValue := template.ParseFiles("../../static/html/unfollowStatus.html")
		if errorValue != nil {
			panic(errorValue)
		}
		errorValue = temp.Execute(w, nil)
		if errorValue != nil {
			panic(errorValue)
		}
	} else {
		userIdNum, errorValue := strconv.Atoi(userCookie.Value)
		if errorValue != nil {
			panic(errorValue)
		}
		currUserIdNum := user_pb.UserStruct{
			IdNum: int32(userIdNum),
		}
		currentUser, _ := clientObj.UserClient.GetUserFromUserId(context.Background(), &currUserIdNum)
		var usersCurrCanFollow []string
		usersCurrCanFollow = currentUser.User.Following

		for i := 0; i < len(currentUser.User.Following); i += 1 {
			if currentUser.User.Following[i] == currentUser.User.UserName {
				usersCurrCanFollow = append(usersCurrCanFollow[:i], usersCurrCanFollow[i+1:]...)
			}
		}

		log.Println("User followee List: ", usersCurrCanFollow)
		type UserFollowStruct struct {
			UserFollowList []string
		}

		UserFolloweeLists := UserFollowStruct{
			UserFollowList: make([]string, 0),
		}

		UserFolloweeLists.UserFollowList = usersCurrCanFollow
		temp, errorValue := template.ParseFiles("../../static/html/removefriends.html")
		if errorValue != nil {
			panic(errorValue)
		}
		errorValue = temp.Execute(w, UserFolloweeLists)
		if errorValue != nil {
			panic(errorValue)
		}
	}
}
