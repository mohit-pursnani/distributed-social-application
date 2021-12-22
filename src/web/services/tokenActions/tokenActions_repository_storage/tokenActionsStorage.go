package tokenActions_methods

import (
	"context"
	"log"
	"os"

	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/raftHelper"
	"github.com/os3224/final-project-0b5a2e16-santhoshbhk-mohit-pursnani/web/services/tokenActions/tokenActions_pb"
)

func (s *Server) RegisterToken(ctx context.Context, user *tokenActions_pb.TokenActionsPbStruct) (*tokenActions_pb.TokenActionsPbStruct, error) {

	mutex.Lock()
	tokenKey := GenerateRandToken()

	tokenObj.authenticationToken.TokenMap[tokenKey] = user.UserID

	log.Println("Token Map after registration: ", tokenObj.authenticationToken.TokenMap)

	tokKeyVal := &tokenActions_pb.TokenActionsPbStruct{
		Key: tokenKey,
	}
	mutex.Unlock()

	_, errorValue := raftHelper.RequestForService("tokenObj", tokenObj.authenticationToken, "PUT")

	if errorValue != nil {
		log.Println("Error occurred while registering token: ", errorValue)
		os.Exit(1)
	}

	return tokKeyVal, nil
}

func (s *Server) DeleteToken(ctx context.Context, token *tokenActions_pb.TokenActionsPbStruct) (*tokenActions_pb.TokenActionsPbStruct, error) {

	mutex.Lock()
	tokenKeyToDelete := token.Key
	delete(tokenObj.authenticationToken.TokenMap, tokenKeyToDelete)

	statusObj := &tokenActions_pb.TokenActionsPbStruct{
		TokenStatus: true,
	}
	mutex.Unlock()

	_, errorValue := raftHelper.RequestForService("tokenObj", tokenObj.authenticationToken, "PUT")

	if errorValue != nil {
		log.Println("Error occurred while registering token: ", errorValue)
		os.Exit(1)
	}
	log.Println("Token Map Object after token deletion", tokenObj.authenticationToken.TokenMap)

	return statusObj, nil
}
