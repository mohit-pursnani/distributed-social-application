package main

import (
	"context"
	"log"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/user/user_pb"
)

func FollowTest() {
	log.Println("xxxxxxxxxxRunning Follow Test casexxxxxxxxxx")
	passedTestCases := true

	/*currUserIdNum := user_pb.UserStruct{
		IdNum: int32(9999),
	}
	_, _ = clientObj.UserClient.ClearAllUsers(context.Background(), &currUserIdNum)

	RegisterTestNewSetOfUsers()*/

	//Test 1: Santhosh follows longtao
	actionForFolloweeStruct1 := user_pb.ActionForFollowee{
		CurrentUserName:      "santhosh",
		UserNameForOperation: "longtao",
	}
	followStatus1, errVal1 := clientObj.UserClient.FollowUser(context.Background(), &actionForFolloweeStruct1)

	if errVal1 != nil || !followStatus1.Status {
		passedTestCases = false
		log.Fatal("Error: ", errVal1)
		log.Fatal("Test case 1 failed...")
	}

	//Test 2: longtao follows Santhosh
	actionForFolloweeStruct2 := user_pb.ActionForFollowee{
		CurrentUserName:      "longtao",
		UserNameForOperation: "santhosh",
	}
	followStatus2, errVal2 := clientObj.UserClient.FollowUser(context.Background(), &actionForFolloweeStruct2)

	if errVal2 != nil || !followStatus2.Status {
		passedTestCases = false
		log.Fatal("Error: ", errVal2)
		log.Fatal("Test case 2 failed...")
	}

	//Test 3: Santhosh follows mohit
	actionForFolloweeStruct3 := user_pb.ActionForFollowee{
		CurrentUserName:      "santhosh",
		UserNameForOperation: "mohit",
	}
	followStatus3, errVal3 := clientObj.UserClient.FollowUser(context.Background(), &actionForFolloweeStruct3)

	if errVal3 != nil || !followStatus3.Status {
		passedTestCases = false
		log.Fatal("Error: ", errVal3)
		log.Fatal("Test case 3 failed...")
	}

	//Test 4: mohit follows Santhosh
	actionForFolloweeStruct4 := user_pb.ActionForFollowee{
		CurrentUserName:      "mohit",
		UserNameForOperation: "santhosh",
	}
	followStatus4, errVal4 := clientObj.UserClient.FollowUser(context.Background(), &actionForFolloweeStruct4)

	if errVal4 != nil || !followStatus4.Status {
		passedTestCases = false
		log.Fatal("Error: ", errVal4)
		log.Fatal("Test case 4 failed...")
	}

	//Test 5: mohit follows longtao
	actionForFolloweeStruct5 := user_pb.ActionForFollowee{
		CurrentUserName:      "mohit",
		UserNameForOperation: "longtao",
	}
	followStatus5, errVal5 := clientObj.UserClient.FollowUser(context.Background(), &actionForFolloweeStruct5)

	if errVal5 != nil || !followStatus5.Status {
		passedTestCases = false
		log.Fatal("Error: ", errVal5)
		log.Fatal("Test case 5 failed...")
	}

	//Test 6: longtao follows mohit
	actionForFolloweeStruct6 := user_pb.ActionForFollowee{
		CurrentUserName:      "longtao",
		UserNameForOperation: "mohit",
	}
	followStatus6, errVal6 := clientObj.UserClient.FollowUser(context.Background(), &actionForFolloweeStruct6)

	if errVal6 != nil || !followStatus6.Status {
		passedTestCases = false
		log.Fatal("Error: ", errVal6)
		log.Fatal("Test case 6 failed...")
	}

	//Check if each user is following the other two users
	currUserIdNum1 := user_pb.UserStruct{
		IdNum: int32(9999),
	}
	allUsers, _ := clientObj.UserClient.GetAllUsers(context.Background(), &currUserIdNum1)
	log.Println(allUsers.User["santhosh"].Following)
	log.Println(allUsers.User["longtao"].Following)
	log.Println(allUsers.User["mohit"].Following)
	for k, v := range allUsers.User {
		if k == "santhosh" {
			if len(v.Following) != 2 {
				passedTestCases = false
			}
			if v.Following[0] != "longtao" || v.Following[1] != "mohit" {
				passedTestCases = false
			}
		}
		if k == "longtao" {
			if len(v.Following) != 2 {
				passedTestCases = false
			}
			if v.Following[0] != "santhosh" || v.Following[1] != "mohit" {
				passedTestCases = false
			}
		}
		if k == "mohit" {
			if len(v.Following) != 2 {
				passedTestCases = false
			}
			if v.Following[0] != "santhosh" || v.Following[1] != "longtao" {
				passedTestCases = false
			}
		}
	}

	if !passedTestCases {
		log.Println("Follow Test cases failed....")
	} else {
		log.Println("Follow Test cases Passed!")
	}

}

func UnFollowTest() {
	log.Println("xxxxxxxxxxRunning UnFollow Test casexxxxxxxxxx")
	passedTestCases := true
	//Test 1: santhosh unfollows mohit
	actionForFolloweeStruct1 := user_pb.ActionForFollowee{
		CurrentUserName:      "santhosh",
		UserNameForOperation: "mohit",
	}
	unFollowStatus1, errVal1 := clientObj.UserClient.UnfollowUser(context.Background(), &actionForFolloweeStruct1)

	if errVal1 != nil || !unFollowStatus1.Status {
		passedTestCases = false
		log.Fatal("Error: ", errVal1)
		log.Fatal("Test case 1 failed...")
	}
	//Test 2: longtao unfollows santhosh
	actionForFolloweeStruct2 := user_pb.ActionForFollowee{
		CurrentUserName:      "longtao",
		UserNameForOperation: "santhosh",
	}
	unFollowStatus2, errVal2 := clientObj.UserClient.UnfollowUser(context.Background(), &actionForFolloweeStruct2)

	if errVal2 != nil || !unFollowStatus2.Status {
		passedTestCases = false
		log.Fatal("Error: ", errVal2)
		log.Fatal("Test case 2 failed...")
	}

	//Test 3: mohit un follows longtao
	actionForFolloweeStruct3 := user_pb.ActionForFollowee{
		CurrentUserName:      "mohit",
		UserNameForOperation: "longtao",
	}
	unFollowStatus3, errVal3 := clientObj.UserClient.UnfollowUser(context.Background(), &actionForFolloweeStruct3)

	if errVal3 != nil || !unFollowStatus3.Status {
		passedTestCases = false
		log.Fatal("Error: ", errVal3)
		log.Fatal("Test case 3 failed...")
	}

	currUserIdNum1 := user_pb.UserStruct{
		IdNum: int32(9999),
	}

	allUsers, _ := clientObj.UserClient.GetAllUsers(context.Background(), &currUserIdNum1)
	log.Println(allUsers.User["santhosh"].Following)
	log.Println(allUsers.User["longtao"].Following)
	log.Println(allUsers.User["mohit"].Following)

	for k, v := range allUsers.User {
		if k == "santhosh" {
			if len(v.Following) != 1 {
				passedTestCases = false
			}
			if v.Following[0] != "longtao" {
				passedTestCases = false
			}
		}
		if k == "longtao" {
			if len(v.Following) != 1 {
				passedTestCases = false
			}
			if v.Following[0] != "mohit" {
				passedTestCases = false
			}
		}
		if k == "mohit" {
			if len(v.Following) != 1 {
				passedTestCases = false
			}
			if v.Following[0] != "santhosh" {
				passedTestCases = false
			}
		}
	}

	if !passedTestCases {
		log.Println("UnFollow Test cases failed....")
	} else {
		log.Println("UnFollow Test cases Passed!")
	}

}

func FollowBack() {

	//Follow back the users after testing unfollow
	actionForFolloweeStruct3f := user_pb.ActionForFollowee{
		CurrentUserName:      "santhosh",
		UserNameForOperation: "mohit",
	}
	followStatus3f, errVal3f := clientObj.UserClient.FollowUser(context.Background(), &actionForFolloweeStruct3f)

	if errVal3f != nil || !followStatus3f.Status {
		log.Fatal("Error: ", errVal3f)
	}

	actionForFolloweeStruct4f := user_pb.ActionForFollowee{
		CurrentUserName:      "longtao",
		UserNameForOperation: "santhosh",
	}
	followStatus4f, errVal4f := clientObj.UserClient.FollowUser(context.Background(), &actionForFolloweeStruct4f)

	if errVal4f != nil || !followStatus4f.Status {
		log.Fatal("Error: ", errVal4f)
	}

	actionForFolloweeStruct5f := user_pb.ActionForFollowee{
		CurrentUserName:      "mohit",
		UserNameForOperation: "longtao",
	}
	followStatus5f, errVal5f := clientObj.UserClient.FollowUser(context.Background(), &actionForFolloweeStruct5f)

	if errVal5f != nil || !followStatus5f.Status {
		log.Fatal("Error: ", errVal5f)
	}
}
