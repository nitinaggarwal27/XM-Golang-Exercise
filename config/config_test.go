package config

import (
	"os"
	"testing"
)

func TestSetConfig(t *testing.T) {

	// test case 1-> call configuration with environment variables
	TomlFile = ""
	os.Unsetenv("PORT")
	SetConfig()
	Conf = Config{}

	// test case 2 -> invalid toml file
	TomlFile = "../README.md"
	SetConfig()
	Conf = Config{}

	// test case 3 -> valid toml file
	TomlFile = "../example.toml"
	SetConfig()
}
