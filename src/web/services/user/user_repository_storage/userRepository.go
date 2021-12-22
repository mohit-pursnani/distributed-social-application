/*
	Data structure (map) to store user name to user details and user id to username implemented.
	All methods to add and get user details from data structure implemented
*/

package user_methods

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"sync"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/raftHelper"
	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/user/user_pb"
)

var mutex sync.Mutex

type Server struct{}

type userObj struct {
	um *user_pb.UsersMap
}

type userIdToNameListObj struct {
	uIdToName *user_pb.IdToUserName
}

var userList userObj
var userIDToNameList userIdToNameListObj

func InitializeUserService() {
	ul := CreateUserMap()
	userList.um = ul
	uidtnl := CreateIdToUserNameMap()
	userIDToNameList.uIdToName = uidtnl

	_, errorValue := raftHelper.RequestForService("userList", userList.um, "PUT")
	if errorValue != nil {
		log.Println("Error occurred while PUT Request in Post Metod: ", errorValue)
		os.Exit(1)
	}
	_, errorValue = raftHelper.RequestForService("userIDToNameList", userIDToNameList.uIdToName, "PUT")
	if errorValue != nil {
		log.Println("Error occurred while PUT Request in Post Metod: ", errorValue)
		os.Exit(1)
	}
}

func CreateUserMap() *user_pb.UsersMap {
	return &user_pb.UsersMap{
		User: make(map[string]*user_pb.UserStruct, 0),
	}
}

func CreateIdToUserNameMap() *user_pb.IdToUserName {
	return &user_pb.IdToUserName{
		UserIDToNameMap: make(map[int32]string, 0),
	}
}

func AddIdToUserName(uId int, uName string) {
	userIDToNameListFromStorage := GetUserIDListFromStorage()
	userIDToNameListFromStorage.UserIDToNameMap[int32(uId)] = uName
	_, errorValue := raftHelper.RequestForService("userIDToNameList", userIDToNameListFromStorage, "PUT")
	if errorValue != nil {
		log.Println("Error occurred while PUT Request in Post Metod: ", errorValue)
		os.Exit(1)
	}
	log.Println("userIdToUserNameMap after Adding user: ", userIDToNameListFromStorage.UserIDToNameMap)
}

func (s *Server) CheckCred(ctx context.Context, user *user_pb.UserStruct) (*user_pb.UserStatus, error) {
	mutex.Lock()
	userListFromStorage := GetUserListFromStorage()
	//If username exists
	_, found := userListFromStorage.User[user.UserName]
	hashedPassword := GetHash(user.Password)
	// user credential is correct
	userCrdentialStatus := &user_pb.UserStatus{
		Status: true,
	}

	if found && userListFromStorage.User[user.UserName].Password == hashedPassword {
		mutex.Unlock()
		return userCrdentialStatus, nil
	}

	userCrdentialStatus.Status = false
	mutex.Unlock()
	return userCrdentialStatus, nil
}

func (s *Server) GetUserNameByUserId(ctx context.Context, user *user_pb.UserStruct) (*user_pb.UserStruct, error) {
	mutex.Lock()
	userIDToNameListFromStorage := GetUserIDListFromStorage()
	mutex.Unlock()
	return &user_pb.UserStruct{
		UserName: userIDToNameListFromStorage.UserIDToNameMap[user.IdNum],
	}, nil
}

func (s *Server) GetUserFromUserName(ctx context.Context, user *user_pb.UserStruct) (*user_pb.UserStructStruct, error) {
	mutex.Lock()
	userListFromStorage := GetUserListFromStorage()
	mutex.Unlock()
	return &user_pb.UserStructStruct{
		User: userListFromStorage.User[user.UserName],
	}, nil
}

func GetHash(inpPass string) string {
	hash := md5.Sum([]byte(inpPass))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func (s *Server) GetUserFromUserId(ctx context.Context, user *user_pb.UserStruct) (*user_pb.UserStructStruct, error) {
	mutex.Lock()
	userIdNum := user.IdNum
	userName := GetUserFromUserName(userIdNum)
	userListFromStorage := GetUserListFromStorage()
	currUserStruct := userListFromStorage.User[userName]
	retUserStruct := user_pb.UserStructStruct{
		User: currUserStruct,
	}
	log.Println("User list from get call: ", currUserStruct.Following)
	mutex.Unlock()
	return &retUserStruct, nil
}

func GetUserFromUserName(idNum int32) string {
	userIDListFromStorage := GetUserIDListFromStorage()
	return userIDListFromStorage.UserIDToNameMap[idNum]
}

func (s *Server) ClearAllUsers(ctx context.Context, user *user_pb.UserStruct) (*user_pb.UsersMap, error) {
	mutex.Lock()
	ul := CreateUserMap()
	userList.um = ul
	uidtnl := CreateIdToUserNameMap()
	userIDToNameList.uIdToName = uidtnl

	_, errorValue := raftHelper.RequestForService("userList", userList.um, "PUT")
	if errorValue != nil {
		log.Println("Error occurred while PUT Request in Post Metod: ", errorValue)
		os.Exit(1)
	}
	_, errorValue = raftHelper.RequestForService("userIDToNameList", userIDToNameList.uIdToName, "PUT")
	if errorValue != nil {
		log.Println("Error occurred while PUT Request in Post Metod: ", errorValue)
		os.Exit(1)
	}

	userListFromStorage := GetUserListFromStorage()
	retUserMap := user_pb.UsersMap{
		User: userListFromStorage.User,
	}
	mutex.Unlock()
	return &retUserMap, nil
}
