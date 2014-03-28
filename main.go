package main

import (
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

type BackendFunc func(file *ini.File, event *string, text *string)

func ProwlBackend(file *ini.File, event *string, text *string) {
  notification := &gowl.Notification{
    Application: APPLICATION_NAME,
    Event:       *event,
    Description: *text,
  }

  api_key, ok := file.Get("prowl", "api_key")
  if !ok {
    log.Fatal("Prowl API Key not found")
  }

  g := gowl.New(api_key)
  g.Add(notification)
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
  flag.Parse()

  backends := make(map[string]BackendFunc)
  backends["prowl"] = ProwlBackend

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
