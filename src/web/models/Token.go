package models

import "math/rand"

type TokenMap map[string]int

// create a new object of type TokenMap
func CreateTokenObject() TokenMap {
	return make(TokenMap)
}

func GenerateRandToken() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	n := 16
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	strTok := string(b)
	return strTok
}

// add token for the user id
func (tm TokenMap) RegisterToken(userId int) string {
	genTok := GenerateRandToken()
	tm[genTok] = userId
	return genTok
}
