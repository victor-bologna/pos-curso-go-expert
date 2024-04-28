package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

type DBConfig struct {
	DB       string `json:"db"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var config DBConfig

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	MarshalFile("config.json")
	fmt.Println(config)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("Event: ", event)
				if event.Has(fsnotify.Write) {
					MarshalFile("config.json")
					fmt.Println("Modified File: ", event.Name)
					fmt.Println(config)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error: ", err)
			}
		}
	}()
	err = watcher.Add("./config.json") // Caminho do arquivo
	if err != nil {
		panic(err)
	}
	<-make(chan struct{}) // Segura o programa
}

func MarshalFile(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		panic(err)
	}
}
