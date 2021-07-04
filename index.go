package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/skratchdot/open-golang/open"
)

type ClassTime struct {
	Day  string `json:"day"`
	Time string `json:"time"`
}

type Class struct {
	Link  string      `json:"link"`
	Times []ClassTime `json:"times"`
}

type ClassConfig map[string]Class

//go:embed sample.json
var sampleConfigStr string

func openClassLink(config map[string]Class) {
	open.Run("https://google.com")
}

func main() {
	APPDATA_PATH := os.Getenv("APPDATA")
	if APPDATA_PATH == "" {
		APPDATA_PATH = os.Getenv("HOME")
	}

	CONFIG_PATH := APPDATA_PATH + "/auto-class-launcher-timetable.json"

	var config ClassConfig

	// Now check if file exists
	if data, err := os.ReadFile(CONFIG_PATH); err == nil {
		json.Unmarshal(data, &config)
	} else if os.IsNotExist(err) {
		json.Unmarshal([]byte(sampleConfigStr), &config)

		if os.WriteFile(CONFIG_PATH, []byte(sampleConfigStr), 0666) != nil {
			fmt.Println("Unable to create timetable on disk")
		}
	}

	ticker := time.NewTicker(10 * time.Second)

	for range ticker.C {
		openClassLink(config)
	}
}
