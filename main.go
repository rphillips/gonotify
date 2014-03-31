package main

import (
	"bitbucket.org/kisom/gopush/pushover"
	"flag"
	"fmt"
	"github.com/stevenleeg/gowl"
	"github.com/vaughan0/go-ini"
	"log"
	"os"
	"os/user"
	"path"
)

const (
	CONFIG_FILENAME  = ".gonotify"
	APPLICATION_NAME = "gonotify"
)

var (
	VERSION = "0.1.3"
)

type BackendFunc func(file *ini.File, event *string, text *string)

func ProwlBackend(file *ini.File, event *string, text *string) {
	notification := &gowl.Notification{
		Application: APPLICATION_NAME,
		Event:       *event,
		Description: *text,
	}

	api_key, ok := file.Get("prowl", "api_key")
	if !ok {
		log.Fatal("Prowl API key not found")
	}

	g := gowl.New(api_key)
	g.Add(notification)
}

func PushoverBackend(file *ini.File, event *string, text *string) {
	api_key, ok := file.Get("pushover", "api_key")
	if !ok {
		log.Fatal("Pushover API key not found")
	}

	user_key, ok := file.Get("pushover", "user_key")
	if !ok {
		log.Fatal("Pushover User key not found")
	}

	identity := pushover.Authenticate(api_key, user_key)
	sent := pushover.Notify(identity, fmt.Sprintf("%s:%s", event, text))
	if !sent {
		fmt.Println("[!] notification failed.")
		os.Exit(1)
	}
}

func getBackend(config *ini.File) (string, bool) {
	return config.Get("gonotify", "backend")
}

func validateConfig(config *ini.File) {
	_, ok := getBackend(config)
	if !ok {
		log.Fatal("backend missing from config")
	}
}

func main() {

	event := flag.String("event", "event", "String to send to notification service")
	text := flag.String("text", "done", "String to send to notification service")
	version := flag.Bool("v", false, "Display version")
	flag.Parse()

	if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	backends := make(map[string]BackendFunc)
	backends["prowl"] = ProwlBackend
	backends["pushover"] = PushoverBackend

	// Get User's home directory
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configFilePath := path.Join(usr.HomeDir, CONFIG_FILENAME)

	// Load INI File
	config, err := ini.LoadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	validateConfig(&config)

	backendName, _ := getBackend(&config)
	fun := backends[backendName]
	if fun == nil {
		fmt.Println("backend not supported")
	}
	fun(&config, event, text)

	os.Exit(0)
}
