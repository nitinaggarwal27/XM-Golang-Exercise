package database

import (
	"nitinaggarwal27/XM-Golang-Exercise/model"

	"gorm.io/gorm"
)

func CreateDatabaseTables(db *gorm.DB) {
	if !(db.Migrator().HasTable(model.User{})) {
		db.Migrator().CreateTable(model.User{})

		//creating admin account
		adminAcc := model.InitAdminAccount()
		db.Create(&adminAcc)
	}

	if !(db.Migrator().HasTable(model.Company{})) {
		db.Migrator().CreateTable(model.Company{})
	}
	// Database migration
	db.AutoMigrate(
		&model.Company{},
		&model.User{},
	)
}
