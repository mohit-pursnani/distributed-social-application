package main

import (
	"log"
	"net/http"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/modules"
)

func main() {
	log.Println("Connecting to server on port 8000")
	modules.Initialize()
	http.HandleFunc("/", modules.LoginFunctionality)
	http.HandleFunc("/login/", modules.LoginFunctionality)
	http.HandleFunc("/register/", modules.RegisterFunctionality)
	http.HandleFunc("/logout/", modules.LogoutFunctionality)
	http.HandleFunc("/createpost/", modules.CreatePost)
	http.HandleFunc("/findfriends/", modules.FollowUser)
	http.HandleFunc("/unfollowuser/", modules.UnFollowUser)
	http.HandleFunc("/followee/", modules.FolloweeList)
	http.HandleFunc("/feed/", modules.GetPosts)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Failed to connect to server:", err)
	}
}
