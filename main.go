package main

import (
	"flag"
	"log"
	"os"

	"nitinaggarwal27/XM-Golang-Exercise/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// setup for reading flags for deciding whether to do configuration with toml or env variables
	conf := flag.String("conf", "environment", "set configuration from toml file or environment variables")

	file := flag.String("file", "", "set path of toml file")

	flag.Parse()

	if *conf == "environment" {
		log.Println("environment")
		config.ConfigurationWithEnv()
	} else if *conf == "toml" {
		log.Println("toml")
		if *file == "" {
			log.Println("Please pass toml file path")
			os.Exit(1)
		} else {
			err := config.ConfigurationWithToml(*file)
			if err != nil {
				log.Println("Error in parsing toml file")
				log.Println(err)
				os.Exit(1)
			}
		}
	} else {
		log.Println("Please pass valid arguments, conf can be set as toml or environment")
		os.Exit(1)
	}

	//create database
	// err := database.CreateDatabase()
	// if err != nil {
	// 	return
	// }
	// //database config
	// dbConfig := config.DBConfig()
	// //number of ideal connections
	// var ideal int
	// idealStr := config.Conf.Database.Ideal
	// if idealStr == "" {
	// 	ideal = 50
	// } else {
	// 	ideal, _ = strconv.Atoi(idealStr)
	// }
	// connecting db using connection string
	// db, err := gorm.Open("postgres", dbConfig)
	// // fmt.Println(dbConfig)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// // close db instance whenever whole work completed
	// defer db.Close()
	// db.DB().SetMaxIdleConns(ideal)
	// db.DB().SetConnMaxLifetime(1 * time.Hour)
	// config.DB = db

	//create auth-team database tables
	// database.CreateDBTablesIfNotExists()

	//wg.Wait()
	// initialize gin router
	router := gin.Default()

	//allowing CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AddAllowMethods("DELETE")
	router.Use(cors.New(corsConfig))

	// service specific routes
	// api.Routes(router)
	// run the service
	router.Run(":" + config.Conf.Service.Port)
}
