package models

import (
	"crypto/md5"
	"encoding/hex"
)

type UserStruct struct {
	IdNum     int
	UserName  string
	FirstName string
	LastName  string
	Password  string
	Following []UserStruct
}

type UsersMap map[string]UserStruct
type IdToUserName map[int]string

func CreateUserMap() UsersMap {
	return make(UsersMap, 0)
}

func CreateIdToUserNameMap() IdToUserName {
	return make(IdToUserName, 0)
}

func (um *UsersMap) AddToMap(username string, firstName string, lastName string, password string) (UserStruct, bool, int) {

	m := *um

	//If username already exists
	_, found := m[username]

	if found {
		return m[username], true, m[username].IdNum
	}

	incUserId := len(*um) + 1
	currUser := UserStruct{IdNum: incUserId, UserName: username, FirstName: firstName, LastName: lastName, Password: getHash(password)}

	m[username] = currUser

	return currUser, false, currUser.IdNum

}

func (uidToUname *IdToUserName) AddUidToUnameMap(userIdNum int, userName string) {
	(*uidToUname)[userIdNum] = userName
}

func (uidToUname *IdToUserName) GetUname(userIdNum int) string {
	return (*uidToUname)[userIdNum]
}

func getHash(inpPass string) string {
	hash := md5.Sum([]byte(inpPass))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func (um *UsersMap) CheckCred(userName string, password string) (UserStruct, bool) {
	m := *um
	//If username exists
	_, found := m[userName]
	hashed := getHash(password)

	if found && m[userName].Password == hashed {
		return m[userName], true
	}

	return m[userName], false
}

func (um *UsersMap) FollowUser(userName string, userNameToFollow string) {
	user, findStatus := (*um)[userNameToFollow]
	if findStatus {
		if entry, ok := (*um)[userName]; ok {
			entry.Following = append(entry.Following, user)
			(*um)[userName] = entry
		}
	}

}

func (um *UsersMap) UnFollowUser(userName string, userNameToUnFollow string) {
	user, findStatus := (*um)[userNameToUnFollow]
	if findStatus {
		// delete user from list
		if entry, ok := (*um)[userName]; ok {
			for i := 0; i < len(entry.Following); i++ {
				if entry.Following[i].UserName == user.UserName {
					entry.Following = append(entry.Following[:i], entry.Following[i+1:]...)
					(*um)[userName] = entry
					break
				}
			}
		}

	}

}
