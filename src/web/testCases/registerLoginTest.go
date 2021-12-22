package main

import (
	"context"
	"log"
	"strconv"
	"sync"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/user/user_pb"
)

func RegisterTest() {

	log.Println("xxxxxxxxxxRunning Register Test casexxxxxxxxxx")

	user1 := user_pb.UserStruct{
		UserName:  "santhosh",
		Password:  "password123",
		FirstName: "Santhosh",
		LastName:  "Bala Hari Krishnan",
	}

	user2 := user_pb.UserStruct{
		UserName:  "mohit",
		Password:  "password123",
		FirstName: "Mohit",
		LastName:  "Pursnani",
	}

	user3 := user_pb.UserStruct{
		UserName:  "longtao",
		Password:  "password123",
		FirstName: "Longtao",
		LastName:  "Lyu",
	}

	/*var usersToRegister []user_pb.UserStruct

	usersToRegister = append(usersToRegister, user1)
	usersToRegister = append(usersToRegister, user2)
	usersToRegister = append(usersToRegister, user3)*/
	_, errVal1 := clientObj.UserClient.AddToUsersMap(context.Background(), &user1)
	if errVal1 != nil {
		log.Fatal("Test Case to register failed", errVal1)
	}
	_, errVal2 := clientObj.UserClient.AddToUsersMap(context.Background(), &user2)
	if errVal2 != nil {
		log.Fatal("Test Case to register failed", errVal2)
	}
	_, errVal3 := clientObj.UserClient.AddToUsersMap(context.Background(), &user3)

	if errVal3 != nil {
		log.Fatal("Test Case to register failed", errVal3)
	}

	userMap, _ := clientObj.UserClient.GetAllUsers(context.Background(), &user1)

	log.Println("Users from server side: ")
	for _, v := range userMap.User {
		log.Println()
		log.Println("UserName: ", v.UserName)
		log.Println("FirstName: ", v.FirstName)
		log.Println("LastName: ", v.LastName)
		log.Println("Password: ", v.Password)
		log.Println()
	}

	if len(userMap.User) == 3 {
		log.Println("All testcases passed for register!.... 3 users were added and read successfully")
	} else {
		log.Println("Register Test case failed.... 3 users were not added properly")
	}
}

func BatchRegister() {
	log.Println("xxxxxxxxxxRunning batch Concurrent Register Test casexxxxxxxxxx")
	passedTestCases := true
	wg := sync.WaitGroup{}
	for i := 0; i < 300; i += 1 {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			useri := user_pb.UserStruct{
				UserName:  strconv.Itoa(i),
				Password:  strconv.Itoa(i),
				FirstName: strconv.Itoa(i),
				LastName:  strconv.Itoa(i),
			}

			_, errVal1 := clientObj.UserClient.AddToUsersMap(context.Background(), &useri)
			if errVal1 != nil {
				log.Fatal("Test Case to register failed", errVal1)
				passedTestCases = false
			}
		}(i)
	}

	wg.Wait()

	useri := user_pb.UserStruct{
		UserName:  strconv.Itoa(1),
		Password:  strconv.Itoa(1),
		FirstName: strconv.Itoa(1),
		LastName:  strconv.Itoa(1),
	}
	userMap, _ := clientObj.UserClient.GetAllUsers(context.Background(), &useri)

	if !passedTestCases {
		log.Println("Register Test case failed.... 300 users were not added properly")
	} else if len(userMap.User) == 303 { //3 users already present from initial registration
		log.Println("All testcases passed for register!.... 300 users were added and read successfully")
	} else {
		log.Println("Register Test case failed.... 300 users were not added properly")
	}
}

func LoginTest() {
	testCasesPassed := true
	log.Println("xxxxxxxxxxRunning Login Test casexxxxxxxxxxxxxxxx")

	//Test 1: Entering correct credentials for user 1
	user1correct := user_pb.UserStruct{
		UserName: "santhosh",
		Password: "password123",
	}

	responseStatus1, _ := clientObj.UserClient.CheckCred(context.Background(), &user1correct)

	if !responseStatus1.Status {
		//The username and password was entered correctly. So, this should return true, if it returns false the test case failed
		log.Println("Test case 1 failed")
		testCasesPassed = false
	}

	//Test 2: Enterring incorrect credentials for user 1
	user1incorrect := user_pb.UserStruct{
		UserName: "santhosh",
		Password: "incorrectPassword",
	}

	responseStatus2, _ := clientObj.UserClient.CheckCred(context.Background(), &user1incorrect)

	if responseStatus2.Status {
		//The username and password was entered incorrect. So, this should return false, if it returns true the test case failed
		log.Println("Test case 2 failed")
		testCasesPassed = false
	}

	//Test 3: Entering correct credentials for user 2
	user2correct := user_pb.UserStruct{
		UserName: "mohit",
		Password: "password123",
	}

	responseStatus3, _ := clientObj.UserClient.CheckCred(context.Background(), &user2correct)

	if !responseStatus3.Status {
		//The username and password was entered correctly. So, this should return true, if it returns false the test case failed
		log.Println("Test case 3 failed")
		testCasesPassed = false
	}

	//Test 4: Enterring incorrect credentials for user 2
	user2incorrect := user_pb.UserStruct{
		UserName: "mohit",
		Password: "incorrectPassword",
	}

	responseStatus4, _ := clientObj.UserClient.CheckCred(context.Background(), &user2incorrect)

	if responseStatus4.Status {
		//The username and password was entered incorrect. So, this should return false, if it returns true the test case failed
		log.Println("Test case 4 failed")
		testCasesPassed = false
	}

	//Test 5: Entering correct credentials for user 3
	user3correct := user_pb.UserStruct{
		UserName: "mohit",
		Password: "password123",
	}

	responseStatus5, _ := clientObj.UserClient.CheckCred(context.Background(), &user3correct)

	if !responseStatus5.Status {
		//The username and password was entered correctly. So, this should return true, if it returns false the test case failed
		log.Println("Test case 5 failed")
		testCasesPassed = false
	}

	//Test 6: Enterring incorrect credentials for user 3
	user3incorrect := user_pb.UserStruct{
		UserName: "mohit",
		Password: "incorrectPassword",
	}

	responseStatus6, _ := clientObj.UserClient.CheckCred(context.Background(), &user3incorrect)

	if responseStatus6.Status {
		//The username and password was entered incorrect. So, this should return false, if it returns true the test case failed
		log.Println("Test case 6 failed")
		testCasesPassed = false
	}

	if !testCasesPassed {
		log.Println("Login Test cases failed.... The credential status were not returned correctly")
	} else {
		log.Println("Login Test cases Passed!")
	}
}

func BatchLogin() {
	log.Println("xxxxxxxxxxRunning batch Concurrent Login Test casexxxxxxxxxx")
	passedTestCases := true
	wg := sync.WaitGroup{}
	for i := 0; i < 300; i += 1 {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			usercorrect := user_pb.UserStruct{
				UserName: strconv.Itoa(i),
				Password: strconv.Itoa(i),
			}

			responseStatus, _ := clientObj.UserClient.CheckCred(context.Background(), &usercorrect)

			if !responseStatus.Status {
				passedTestCases = false
			}
		}(i)
	}

	wg.Wait()

	if !passedTestCases {
		log.Println("Login Test case failed....")
	} else {
		log.Println("Login Test passed!")
	}
}

func RegisterTestNewSetOfUsers() {

	user1 := user_pb.UserStruct{
		UserName:  "santhosh",
		Password:  "password123",
		FirstName: "Santhosh",
		LastName:  "Bala Hari Krishnan",
	}

	user2 := user_pb.UserStruct{
		UserName:  "mohit",
		Password:  "password123",
		FirstName: "Mohit",
		LastName:  "Pursnani",
	}

	user3 := user_pb.UserStruct{
		UserName:  "longtao",
		Password:  "password123",
		FirstName: "Longtao",
		LastName:  "Lyu",
	}
	_, errVal1 := clientObj.UserClient.AddToUsersMap(context.Background(), &user1)
	if errVal1 != nil {
		log.Fatal("Test Case to register failed", errVal1)
	}
	_, errVal2 := clientObj.UserClient.AddToUsersMap(context.Background(), &user2)
	if errVal2 != nil {
		log.Fatal("Test Case to register failed", errVal2)
	}
	_, errVal3 := clientObj.UserClient.AddToUsersMap(context.Background(), &user3)

	if errVal3 != nil {
		log.Fatal("Test Case to register failed", errVal3)
	}

	_, _ = clientObj.UserClient.GetAllUsers(context.Background(), &user1)

}
