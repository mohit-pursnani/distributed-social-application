package user_methods

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"os"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/raftHelper"
	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/user/user_pb"
)

func (s *Server) AddToUsersMap(ctx context.Context, user *user_pb.UserStruct) (*user_pb.UserStatus, error) {
	mutex.Lock()
	userListFromStorage := GetUserListFromStorage()
	//If username already exists
	_, found := userListFromStorage.User[user.UserName]
	// user already present
	userAddStatus := &user_pb.UserStatus{
		Status: false,
	}

	if found {
		mutex.Unlock()
		return userAddStatus, nil
	}

	var incrementedUserID int32 = int32(len(userListFromStorage.User) + 1)
	//currUser := UserStruct{IdNum: incUserId, UserName: username, FirstName: firstName, LastName: lastName, Password: getHash(password)}
	currentUser := &user_pb.UserStruct{
		IdNum:     incrementedUserID,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  GetHash(user.Password),
	}
	//log.Println("UserName before map: ", user.UserName)
	userListFromStorage.User[user.UserName] = currentUser
	log.Println("UserMap after Adding user: ", userListFromStorage.User)
	// user is added
	userAddStatus.Status = true

	_, errorValue := raftHelper.RequestForService("userList", userListFromStorage, "PUT")
	if errorValue != nil {
		log.Println("Error occurred while PUT Request in Post Metod: ", errorValue)
		os.Exit(1)
	}

	//if status is true => user is added => add user to uid to uname map:
	if userAddStatus.Status {
		AddIdToUserName(int(currentUser.IdNum), currentUser.UserName)
	}
	mutex.Unlock()
	return userAddStatus, nil
}

func (s *Server) FollowUser(ctx context.Context, followee *user_pb.ActionForFollowee) (*user_pb.UserStatus, error) {
	mutex.Lock()
	userListFromStorage := GetUserListFromStorage()
	followeeUser, findStatus := userListFromStorage.User[followee.UserNameForOperation]
	userActionStatus := &user_pb.UserStatus{
		Status: true,
	}
	if findStatus {
		if entry, ok := userListFromStorage.User[followee.CurrentUserName]; ok {

			alreadyPresent := false
			for i := 0; i < len(entry.Following); i += 1 {
				if entry.Following[i] == followeeUser.UserName {
					alreadyPresent = true
				}
			}
			if !alreadyPresent {
				entry.Following = append(entry.Following, followeeUser.UserName)
			}

			userListFromStorage.User[followee.CurrentUserName] = entry
			log.Println("User Name after appending: ", entry.UserName)
			_, errorValue := raftHelper.RequestForService("userList", userListFromStorage, "PUT")
			if errorValue != nil {
				log.Println("Error occurred while PUT Request in FollowUser Method: ", errorValue)
				os.Exit(1)
			}

			log.Println("User Followee list from storage: ", userListFromStorage)
			mutex.Unlock()
			return userActionStatus, nil
		}
	}
	userActionStatus.Status = false
	mutex.Unlock()
	return userActionStatus, nil
}

func (s *Server) UnfollowUser(ctx context.Context, unFollow *user_pb.ActionForFollowee) (*user_pb.UserStatus, error) {
	mutex.Lock()
	userListFromStorage := GetUserListFromStorage()
	unFollowUser, findStatus := userListFromStorage.User[unFollow.UserNameForOperation]
	userActionStatus := &user_pb.UserStatus{
		Status: true,
	}
	if findStatus {
		// delete user from list
		if entry, ok := userListFromStorage.User[unFollow.CurrentUserName]; ok {
			for i := 0; i < len(entry.Following); i++ {
				if entry.Following[i] == unFollowUser.UserName {
					entry.Following = append(entry.Following[:i], entry.Following[i+1:]...)
					userListFromStorage.User[unFollow.CurrentUserName] = entry
					break
				}
			}
			_, errorValue := raftHelper.RequestForService("userList", userListFromStorage, "PUT")
			if errorValue != nil {
				log.Println("Error occurred while PUT Request in FollowUser Metod: ", errorValue)
				os.Exit(1)
			}
			log.Println("User Un Follow list from storage: ", userListFromStorage)
		} else {
			userActionStatus.Status = false
			mutex.Unlock()
			return userActionStatus, nil
		}
	} else {
		userActionStatus.Status = false
		mutex.Unlock()
		return userActionStatus, nil
	}
	mutex.Unlock()
	return userActionStatus, nil
}

func GetUserListFromStorage() *user_pb.UsersMap {
	responseInBytes, errorValue := raftHelper.RequestForService("userList", userList.um, "GET")
	if errorValue != nil {
		log.Println("Error occurred while PUT Request in Post Metod: ", errorValue)
		os.Exit(1)
	}

	var userListFromStorage *user_pb.UsersMap

	decoder := gob.NewDecoder(bytes.NewBufferString(responseInBytes))
	decodedDataErr := decoder.Decode(&userList.um)

	if decodedDataErr != nil {
		log.Println("error occurred while decoding: ", decodedDataErr)
		os.Exit(1)
		userListFromStorage, errorValue = nil, decodedDataErr
	} else {
		userListFromStorage, errorValue = userList.um, nil
	}

	if errorValue != nil {
		log.Println("Error occurred while decoding in GET Request: ", errorValue)
		os.Exit(1)
	}
	return userListFromStorage
}

func GetUserIDListFromStorage() *user_pb.IdToUserName {
	responseInBytes, errorValue := raftHelper.RequestForService("userIDToNameList", userIDToNameList.uIdToName, "GET")
	if errorValue != nil {
		log.Println("Error occurred while PUT Request in Post Metod: ", errorValue)
		os.Exit(1)
	}

	var userIDListFromStorage *user_pb.IdToUserName

	decoder := gob.NewDecoder(bytes.NewBufferString(responseInBytes))
	decodedDataErr := decoder.Decode(&userIDToNameList.uIdToName)

	if decodedDataErr != nil {
		log.Println("error occurred while decoding: ", decodedDataErr)
		os.Exit(1)
		userIDListFromStorage, errorValue = nil, decodedDataErr
	} else {
		userIDListFromStorage, errorValue = userIDToNameList.uIdToName, nil
	}

	if errorValue != nil {
		log.Println("Error occurred while decoding in GET Request: ", errorValue)
		os.Exit(1)
	}

	return userIDListFromStorage
}

func (s *Server) GetAllUsers(ctx context.Context, user *user_pb.UserStruct) (*user_pb.UsersMap, error) {
	mutex.Lock()
	userListFromStorage := GetUserListFromStorage()
	retUserMap := user_pb.UsersMap{
		User: userListFromStorage.User,
	}
	mutex.Unlock()
	return &retUserMap, nil
}
