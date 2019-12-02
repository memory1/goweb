package model

import (
	"crypto/sha512"
	"fmt"
	"encoding/base64"
	"testing"
)

func TestLoginSendsCorrectswordHash(t *testing.T){
	testDB := new(mockDB)
	testDB.returnedRow = &mockRow{}
	db = testDB

	password :="the password"
	email :="the email"
	Login(email, password)

	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	fmt.Println("hasher sum: ", hasher.Sum(nil))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	if testDB.lastArgs[1] != pwd {
		t.Errorf("Login function failed to send correct password hash to database")
	}

}