package main

import (
	"flag"
	"log"
	"os"

	"nitinaggarwal27/XM-Golang-Exercise/config"
	"nitinaggarwal27/XM-Golang-Exercise/database"
	"nitinaggarwal27/XM-Golang-Exercise/router"
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

	//setup database
	database.SetupConnection()

	// initialize and run gin router
	router.Routes().Run(":" + config.Conf.Service.Port)
}
