/*
	Using session based authentication in go. Idea is to generate a user cookie/token and use that to figure out
	which user is currently logged in.

	Reference: https://www.sohamkamani.com/golang/session-based-authentication/
*/

package modules

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/tokenActions/tokenActions_pb"
	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/user/user_pb"
)

func LoginFunctionality(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("../../static/html/login.html")
		if err != nil {
			panic(err)
		}
		err = temp.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	} else {
		username := r.FormValue("user_name")
		password := r.FormValue("password")
		user := user_pb.UserStruct{
			UserName: username,
			Password: password,
		}
		responseStatus, err := clientObj.UserClient.CheckCred(context.Background(), &user)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		log.Println("Response Stauts check credential: ", responseStatus.Status)
		if responseStatus.Status {
			log.Println("user found")
			userDetails, errorValue := clientObj.UserClient.GetUserFromUserName(context.Background(), &user)
			if errorValue != nil {
				log.Fatal("Error: ", err)
			}
			authUser := tokenActions_pb.TokenActionsPbStruct{
				UserID: userDetails.User.IdNum,
			}
			log.Println("User ID", authUser.UserID)

			token, _ := clientObj.TokenClient.RegisterToken(context.Background(), &authUser)

			log.Println("Token", token.Key)

			tokenMap, _ := clientObj.TokenClient.GetTokenMap(context.Background(), &authUser)
			log.Println("Token Map", tokenMap.TokenMap)

			tokenCookie := &http.Cookie{Name: "ctok", Value: token.Key, Expires: time.Now().Add(time.Hour), Path: "/"}
			userCookie := &http.Cookie{Name: "user", Value: strconv.Itoa(int(userDetails.User.IdNum)), Expires: time.Now().Add(time.Hour), Path: "/"}

			log.Println("user cookie while setting : " + userCookie.Value)
			http.SetCookie(w, tokenCookie)
			http.SetCookie(w, userCookie)

			http.Redirect(w, r, "/feed/", http.StatusFound)
		} else {
			log.Println("User Not found")
			temp, errorValue := template.ParseFiles("../../static/html/loginStatus.html")
			type LoginStatusStruct struct {
				LoginStatus bool
			}
			AuthLoginStatus := LoginStatusStruct{
				LoginStatus: false,
			}
			AuthLoginStatus.LoginStatus = false

			if errorValue != nil {
				panic(errorValue)
			}
			errorValue = temp.Execute(w, AuthLoginStatus)
			if errorValue != nil {
				panic(errorValue)
			}
		}
	}
}

func LogoutFunctionality(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		token, err := r.Cookie("user")
		if err != nil || token.Value == "" {
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
		log.Println("Deleting token: " + token.Value)

		toDeleteToken := tokenActions_pb.TokenActionsPbStruct{
			Key: token.Value,
		}

		status, _ := clientObj.TokenClient.DeleteToken(context.Background(), &toDeleteToken)

		log.Print("Deleted Token Status: ", status.TokenStatus)

		emptyTokenCookie := &http.Cookie{Name: "empctok", Value: "", Expires: time.Now().Add(-1 * time.Hour), MaxAge: -9999, Path: "/"}
		emptyUserCookie := &http.Cookie{Name: "empuser", Value: "", Expires: time.Now().Add(-1 * time.Hour), MaxAge: -9999, Path: "/"}

		http.SetCookie(w, emptyTokenCookie) //Now has empty token cookie
		http.SetCookie(w, emptyUserCookie)  //Now has empty user cookie

		// go to logout page
	}
	temp, errorValue := template.ParseFiles("../../static/html/logout.html")
	if errorValue != nil {
		panic(errorValue)
	}
	errorValue = temp.Execute(w, nil)
	if errorValue != nil {
		panic(errorValue)
	}
}

func RegisterFunctionality(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		firstName := r.FormValue("first_name")
		lastName := r.FormValue("last_name")
		username := r.FormValue("user_name")
		password := r.FormValue("Password")

		user := user_pb.UserStruct{
			UserName:  username,
			Password:  password,
			FirstName: firstName,
			LastName:  lastName,
		}
		//user.Following = append(user.Following, user.UserName)
		responseStatus, err := clientObj.UserClient.AddToUsersMap(context.Background(), &user)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		log.Println("Response Stauts check credential: ", responseStatus.Status)

		actionForFolloweeStruct := user_pb.ActionForFollowee{
			CurrentUserName:      user.UserName,
			UserNameForOperation: user.UserName,
		}
		followStatus, errVal := clientObj.UserClient.FollowUser(context.Background(), &actionForFolloweeStruct)
		if errVal != nil {
			log.Fatal("Error: ", err)
		}
		log.Println("Follow Status: ", followStatus.Status)

		temp, err := template.ParseFiles("../../static/html/registrationstatus.html")
		if err != nil {
			panic(err)
		}
		err = temp.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	} else {
		temp, err := template.ParseFiles("../../static/html/register.html")
		if err != nil {
			panic(err)
		}
		err = temp.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	}
}
