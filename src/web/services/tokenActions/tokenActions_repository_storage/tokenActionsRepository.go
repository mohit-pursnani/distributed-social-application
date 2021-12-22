/*
	Using session based authentication in go. Idea is to generate a user cookie/token and use that to figure out
	which user is currently logged in.

	Reference: https://www.sohamkamani.com/golang/session-based-authentication/
*/

package tokenActions_methods

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"math/rand"
	"os"
	"sync"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/raftHelper"
	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/tokenActions/tokenActions_pb"
)

var mutex sync.Mutex

type Server struct{}

type TokenObj struct {
	authenticationToken *tokenActions_pb.AuthenticationToken
}

var tokenObj TokenObj

func InitializeTokenActionsService() {
	mutex.Lock()
	tokenValue := CreateToken()
	tokenObj.authenticationToken = tokenValue

	_, errorValue := raftHelper.RequestForService("tokenObj", tokenObj.authenticationToken, "PUT")

	if errorValue != nil {
		log.Println("Error occurred while registering token: ", errorValue)
		os.Exit(1)
	}

	mutex.Unlock()
}

func CreateToken() *tokenActions_pb.AuthenticationToken {

	emptyTokenMapObj := &tokenActions_pb.AuthenticationToken{
		TokenMap: make(map[string]int32),
	}
	return emptyTokenMapObj
}

func (s *Server) GetTokenMap(ctx context.Context, user *tokenActions_pb.TokenActionsPbStruct) (*tokenActions_pb.AuthenticationToken, error) {

	responseInBytes, errorValue := raftHelper.RequestForService("tokenObj", tokenObj.authenticationToken, "GET")
	if errorValue != nil {
		log.Println("Error occurred while GET Request: ", errorValue)
		os.Exit(1)
	}
	mutex.Lock()
	var tokenFromStorage *tokenActions_pb.AuthenticationToken

	decoder := gob.NewDecoder(bytes.NewBufferString(responseInBytes))
	decodedDataErr := decoder.Decode(&tokenObj.authenticationToken)

	if decodedDataErr != nil {
		log.Println("error occurred while decoding: ", decodedDataErr)
		os.Exit(1)
		tokenFromStorage, errorValue = nil, decodedDataErr
	} else {
		tokenFromStorage, errorValue = tokenObj.authenticationToken, nil
	}

	if errorValue != nil {
		log.Println("Error occurred while decoding in GET Request: ", errorValue)
		os.Exit(1)
	}
	mutex.Unlock()
	return tokenFromStorage, nil
}

//Reference: stackoverflow
func GenerateRandToken() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	n := 16
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	tokenValue := string(b)
	return tokenValue
}

func (s *Server) ClearAllTokens(ctx context.Context, token *tokenActions_pb.TokenActionsPbStruct) (*tokenActions_pb.TokenActionsPbStruct, error) {
	statusObj := &tokenActions_pb.TokenActionsPbStruct{
		TokenStatus: true,
	}

	tokenValue := CreateToken()
	tokenObj.authenticationToken = tokenValue

	_, errorValue := raftHelper.RequestForService("tokenObj", tokenObj.authenticationToken, "PUT")

	if errorValue != nil {
		log.Println("Error occurred while registering token: ", errorValue)
		os.Exit(1)
	}

	return statusObj, nil
}
