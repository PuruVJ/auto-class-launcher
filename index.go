package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/skratchdot/open-golang/open"
)

type ClassTime struct {
	Day  string `json:"day,omitempty"`
	Time string `json:"time,omitempty"`
}

type Class struct {
	Link  string      `json:"link,omitempty"`
	Times []ClassTime `json:"times"`
}

type ClassConfig map[string]Class

type ClassToday struct {
	name string
	time time.Time
	link string
}

//go:embed sample.json
var sampleConfigStr []byte

var todaysClassLaunched map[string]bool

func getHourAndMinutes(time string) (int, int) {
	timeArr := strings.Split(time, ":")

	hours, _ := strconv.Atoi(timeArr[0])
	minutes, _ := strconv.Atoi(timeArr[1])

	return hours, minutes
}

func getClassesToday(config ClassConfig, weekDay string) []ClassToday {
	classesToday := []ClassToday{}

	for className, classContent := range config {

		todayTime := ClassTime{}

		for _, val := range classContent.Times {
			if val.Day == weekDay {
				todayTime = val
			}
		}

		// No time found, then move on
		if (todayTime == ClassTime{}) {
			continue
		}

		hours, minutes := getHourAndMinutes(todayTime.Time)
		currentTime := time.Now()
		time := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), hours, minutes, 0, currentTime.Nanosecond(), currentTime.Location())

		classesToday = append(classesToday, ClassToday{link: classContent.Link, name: className, time: time})
	}

	if len(todaysClassLaunched) == 0 {
		for _, classContent := range classesToday {
			todaysClassLaunched[classContent.name] = time.Now().After(classContent.time)
		}
	}

	sort.Slice(classesToday, func(i, j int) bool {
		return classesToday[i].time.Before(classesToday[j].time)
	})

	return classesToday
}

func openClassLink(config ClassConfig) {
	date := time.Now()
	weekDay := strings.ToLower(date.Weekday().String())[0:3]

	// classesToday
	classesToday := getClassesToday(config, weekDay)

	// Find the upcoming class
	upcomingClass := ClassToday{}

	for _, val := range classesToday {
		if val.time.After(date) {
			upcomingClass = val
			break
		}
	}

	// Check if all classes lanched
	allClassesLaunched := true
	for _, launched := range todaysClassLaunched {
		if !launched {
			allClassesLaunched = false
		}
	}

	if (upcomingClass == ClassToday{} || allClassesLaunched) {
		fmt.Fprintln(
			color.Output,
			color.YellowString("No more classes for today 🥳🥳🥳. Feel free to close this window."),
		)
		return
	}

	launchTime := upcomingClass.time.Add(-5 * time.Minute)

	minuteStr := strconv.Itoa(upcomingClass.time.Minute())
	if upcomingClass.time.Minute() < 10 {
		minuteStr = "0" + minuteStr
	}

	if !todaysClassLaunched[upcomingClass.name] {
		fmt.Fprintf(
			color.Output,

			color.BlueString("[RUNNING] Launching next class "+color.CyanString("%s")+" @ %d:%02d\n"),
			upcomingClass.name,
			launchTime.Hour(),
			launchTime.Minute(),
		)
	}

	if launchTime.Before(time.Now()) && !todaysClassLaunched[upcomingClass.name] {
		fmt.Fprintf(color.Output, color.GreenString("[LAUNCHING] Launching %s\n"), upcomingClass.name)

		link := upcomingClass.link
		if len(link) == 0 {

			link += "https://auto-class-launcher-alarm.vercel.app/?className="
			link += url.QueryEscape(upcomingClass.name)
			link += "&timing=" + strconv.Itoa(upcomingClass.time.Hour()) + ":" + minuteStr
		}

		open.Run(link)

		todaysClassLaunched[upcomingClass.name] = true
	}
}

func main() {

	CONFIG_PATH := "./auto-class-launcher-timetable.json"

	var config ClassConfig

	// Now check if file exists
	if data, err := os.ReadFile(CONFIG_PATH); err == nil {
		json.Unmarshal(data, &config)
	} else if os.IsNotExist(err) {
		json.Unmarshal([]byte(sampleConfigStr), &config)

		if os.WriteFile(CONFIG_PATH, sampleConfigStr, 0666) != nil {
			fmt.Println("Unable to create timetable on disk")
		}
	}

	fmt.Fprint(color.Output, color.CyanString(`
  __          ________ _      _____ ____  __  __ ______  
  \ \        / /  ____| |    / ____/ __ \|  \/  |  ____| 
   \ \  /\  / /| |__  | |   | |   | |  | | \  / | |__    
    \ \/  \/ / |  __| | |   | |   | |  | | |\/| |  __|   
     \  /\  /  | |____| |___| |___| |__| | |  | | |____  
      \/  \/   |______|______\_____\____/|_|  |_|______|       
`))

	fmt.Fprintf(
		color.Output,
		"\n\nThis is the "+color.GreenString("Auto Class launcher")+" project! Opens up your class links based on your timetable 5 minutes before\n",
	)

	fmt.Println("\nThis project works based on a config file stored in your computer. Your timetable and links are there only. ")

	fmt.Fprintf(color.Output, "Your config file is stored at "+color.GreenString(CONFIG_PATH)+"\n")

	fmt.Fprint(color.Output, color.YellowString("\nPlease modify the file for your own purposes.\n"))

	fmt.Fprintf(color.Output, "\nTo read about how to modify the file and its format, go to "+color.RedString("https://github.com/PuruVJ/auto-class-launcher\n"))

	ticker := time.NewTicker(10 * time.Second)

	// Initialize todaysClassLaunched here
	todaysClassLaunched = make(map[string]bool)

	for range ticker.C {
		openClassLink(config)
	}
}
