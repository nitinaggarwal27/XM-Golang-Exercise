package jwtToken

import (
	"nitinaggarwal27/XM-Golang-Exercise/model"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtToken is a method for creating claims map to be added in a token
// and also save sessions in cockroach database and redis database
func JwtToken(acs model.User) map[string]interface{} {
	// start span from parent span contex

	// intialise claims map

	claims := make(map[string]interface{})
	// populate claims map
	claims["id"] = acs.ID
	claims["name"] = acs.Name
	claims["email"] = acs.Email
	claims["sys_role"] = acs.Role

	// generate jwt token, expiration time and extra info like (expire jwt time, start and end time)

	mapd := ginJwtToken(claims)

	// check token is empty or not
	if mapd["token"].(string) == "" {

		return mapd
	}

	return mapd
}

// GinJwtToken is a method to generate new token with expiry
// with dynamic payload passed in arguments
func ginJwtToken(setClaims map[string]interface{}) map[string]interface{} {

	// intializing middleware
	mw := MwInitializer()

	// Create the token
	token := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	// extracting claims in form of map
	claims := token.Claims.(jwt.MapClaims)

	// extracting expire time
	expire := mw.TimeFunc().Add(mw.Timeout)

	// setting claims
	for key, val := range setClaims {
		claims[key] = val
	}
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = mw.TimeFunc().Unix()

	mapd := map[string]interface{}{"token": "", "expire": ""}

	// signing token
	tokenString, err := token.SignedString(mw.Key)
	if err != nil {
		return mapd
	}

	// passing map with all information
	mapd = map[string]interface{}{
		"error":  false,
		"token":  tokenString,
		"expire": expire.Format(time.RFC3339),
	}

	return mapd
}
