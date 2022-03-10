package main

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"gotextme/telegram"
)

func getConfig() *ini.File {
	// Construct a path to the expected location of the user's .pgpass file
	usr, _ := user.Current()
	homeDir := usr.HomeDir
	path := filepath.Join(homeDir, ".gotextme.ini")

	// Check if the configuration file exists
	if _, err := os.Stat(path); err != nil {
		log.Println(err)
		log.Fatalf("FATAL: Can not access the configuration file: '%s'", path)
	}

	// Parse the configuration INI file
	config, err := ini.Load(path)
	if err != nil {
		log.Printf("Error parsing the INI file: '%s'\n", path)
		log.Fatalf("ERROR: '%v'", err)
	}

	return config
}

func main() {
	config := getConfig()

	telegram.SendMessage(
		"This is a more sophisticated test from Golang!",
		config,
	)
}
