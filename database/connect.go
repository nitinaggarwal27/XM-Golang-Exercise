package database

import (
	"fmt"
	"nitinaggarwal27/XM-Golang-Exercise/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

//GetDB : get current connection
func GetDB() *gorm.DB {
	return db
}

func SetupConnection() {
	switch config.Conf.Database.Engine {
	case "sqlite":
		SQLiteDBConnect()
	case "postgres":
		PostgresDBConnect()
	case "mysql":
		MySQLDBConnect()
	default:
		SQLiteDBConnect()
	}
	if err != nil {
		panic(err)
	}
	CreateDatabaseTables(db)
}

//SQLiteDBConnect : Create Connection to database
func SQLiteDBConnect() {
	//Connect to database
	db, err = gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", config.Conf.Database.Name)), &gorm.Config{})
}

//PostgresDBConnect : Create Connection to database
func PostgresDBConnect() {
	//Connect to database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.User,
		config.Conf.Database.Pass,
		config.Conf.Database.Name, config.Conf.Database.Ssl)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

//MySQLDBConnect : Create Connection to database
func MySQLDBConnect() {
	//Connect to database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Conf.Database.User,
		config.Conf.Database.Pass,
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.Name)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
