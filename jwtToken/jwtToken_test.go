package jwtToken

import (
	"nitinaggarwal27/XM-Golang-Exercise/model"
	"testing"
)

const (
	email string = "test@test.com"
	name  string = "test"
)

func TestGinJwtToken(t *testing.T) {

	acc := model.User{}
	//save data in account table
	acc.Name = name
	acc.Role = "admin"
	acc.Email = email

	//test csae 1 -> generate token with account only
	mapd := JwtToken(acc)

	if mapd["token"].(string) == "" {
		t.Error("test case fail")
	}
}
