package model

import (
	"nitinaggarwal27/XM-Golang-Exercise/config"
	"nitinaggarwal27/XM-Golang-Exercise/methods"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required" gorm:"not null;unique;"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

// Activities is a structure to record user activties
type Activities struct {
	gorm.Model
	Email        string `json:"email"`
	ActivityName string `json:"activity_name"`
	ClientIP     string `json:"client_ip"`
	ClientAgent  string `json:"client_agent"`
	Timestamp    int64  `json:"timestamp"`
}

// InitAdminAccount is a function used to create admin account
func InitAdminAccount() User {

	// fetching info from env variables
	adminEmail := config.Conf.Admin.Email
	if adminEmail == "" {
		adminEmail = "admin@xenonstack.com"
	}
	adminPass := config.Conf.Admin.Pass
	if adminPass == "" {
		adminPass = "admin"
	}
	// return struct with details of admin
	return User{
		Name:     "admin",
		Password: methods.HashForNewPassword(adminPass),
		Email:    adminEmail,
		Role:     "admin",
	}
}
