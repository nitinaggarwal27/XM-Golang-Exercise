package config

import (
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

// Config is a structure for configuration
type Config struct {
	Service  Service
	Database Database
	Admin    Admin
	JWT      JWT
}

// Database is a structure for cockroach database configuration
type Database struct {
	Name   string
	Engine string
	Host   string
	Port   string
	User   string
	Pass   string
	Ssl    string
}

// Service is a structure for service specific related configuration
type Service struct {
	Port           string
	Environment    string
	Build          string
	ValidLocations string
}

// Admin is a structure for admin account credentials
type Admin struct {
	Email string
	Pass  string
}

// JWT is structure for jwt token specific configuration
type JWT struct {
	PrivateKey    string
	JWTExpireTime time.Duration
}

var (
	// Conf is a global variable for configuration
	Conf Config
	// TomlFile is a global variable for toml file path
	TomlFile string
)

// ConfigurationWithEnv is a method to initialize configuration with environment variables
func ConfigurationWithEnv() {
	// database configuration
	Conf.Database.Engine = os.Getenv("DB_ENGINE")
	Conf.Database.Host = os.Getenv("DB_HOST")
	Conf.Database.Port = os.Getenv("DB_PORT")
	Conf.Database.User = os.Getenv("DB_USER")
	Conf.Database.Pass = os.Getenv("DB_PASS")
	Conf.Database.Name = os.Getenv("DB_NAME")
	Conf.Database.Ssl = "disable"
	//admin credentials
	Conf.Admin.Email = os.Getenv("ADMIN_EMAIL")
	Conf.Admin.Pass = os.Getenv("ADMIN_PASS")

	// if service port is not defined set default port
	if os.Getenv("PORT") != "" {
		Conf.Service.Port = os.Getenv("PORT")
	} else {
		Conf.Service.Port = "8000"
	}
	//service specific configuration
	Conf.Service.Environment = os.Getenv("ENVIRONMENT")
	Conf.Service.Build = os.Getenv("BUILD_IMAGE")

	//JWT Token Timeout in minutes
	Conf.JWT.PrivateKey = os.Getenv("PRIVATE_KEY")
	Conf.JWT.JWTExpireTime = time.Minute * 30
}

// ConfigurationWithToml is a method to initialize configuration with toml file
func ConfigurationWithToml(filePath string) error {
	// set varible as file path if configuration is done using toml
	TomlFile = filePath
	// parse toml file and save data config structure
	_, err := toml.DecodeFile(filePath, &Conf)
	if err != nil {
		log.Println(err)
		return err
	}

	// set constants
	//JWT Token Timeout in minutes
	Conf.JWT.JWTExpireTime = time.Minute * 30
	if Conf.Service.Port == "" {
		Conf.Service.Port = "8000"
	}
	Conf.Database.Ssl = "disable"
	Conf.Service.Build = os.Getenv("BUILD_IMAGE")
	return nil
}

// SetConfig is a method to re-intialise configuration at runtime
func SetConfig() {
	if TomlFile == "" {
		ConfigurationWithEnv()
	} else {
		ConfigurationWithToml(TomlFile)
	}
}
