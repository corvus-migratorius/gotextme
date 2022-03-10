package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/ini.v1"

	"gotextme/telegram"
)

// Load from disk and return a configuration in the INI format
func getConfig() *ini.File {
	// Construct a path to the expected location of the user's configuration file
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
	if len(os.Args) < 2 {
		fmt.Println("Please supply a message to be sent.")
		os.Exit(1)
	}

	config := getConfig()

	telegram.SendMessage(
		os.Args[1],
		config,
	)
}
