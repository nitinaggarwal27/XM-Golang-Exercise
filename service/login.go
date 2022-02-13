package service

import (
	"log"
	"net/http"
	"nitinaggarwal27/XM-Golang-Exercise/database"
	"nitinaggarwal27/XM-Golang-Exercise/jwtToken"
	"nitinaggarwal27/XM-Golang-Exercise/methods"
	"nitinaggarwal27/XM-Golang-Exercise/model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	type Login struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var data Login
	if err := c.BindJSON(&data); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Please enter required fields",
		})
		return
	}
	if !methods.ValidateEmail(data.Email) {
		// if there is some error passing bad status code
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Please enter valid email id."})
		return
	}

	// check request is from mobile or from somewhere else
	userAgent := c.Request.Header.Get("User-Agent")

	//=============================================
	// recording user activity
	activity := model.Activities{
		Email:       data.Email,
		ClientIP:    c.ClientIP(),
		ClientAgent: userAgent,
		Timestamp:   time.Now().Unix(),
	}
	//=============================================
	code, mapd := signin(strings.ToLower(data.Email), data.Password)
	if code == 200 {
		// recording user activity of login
		activity.ActivityName = "login"
	} else {
		// recording user activity of failed login
		activity.ActivityName = "failedlogin"
	}
	db := database.GetDB()
	db.Create(&activity)
	c.JSON(code, mapd)

}

func signin(email, password string) (int, map[string]interface{}) {
	mapd := make(map[string]interface{})

	//connection to db
	db := database.GetDB()

	//Checking whether registered or not
	var account []model.User
	db.Where("email=?", email).Find(&account)

	// when no account found
	if len(account) == 0 {
		mapd["error"] = true
		mapd["message"] = "Account doesn’t exist"
		return 401, mapd
	}

	// checking password with saved password
	if methods.CheckHashForPassword(account[0].Password, password) {
		mapd = jwtToken.JwtToken(account[0])
		mapd["name"] = account[0].Name
		mapd["role"] = account[0].Role
		mapd["email"] = account[0].Email
		return 200, mapd
	}
	// when password not matched
	mapd["error"] = true
	mapd["message"] = "Invalid email or password"
	return 401, mapd
}
