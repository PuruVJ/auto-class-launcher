package main

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

func get_time_from_hh_mm(hh_mm_str string) time.Time {
	timeArr := strings.Split(hh_mm_str, ":")

	hours, _ := strconv.Atoi(timeArr[0])
	minutes, _ := strconv.Atoi(timeArr[1])

	currentTime := time.Now()
	time := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), hours, minutes, 0, currentTime.Nanosecond(), currentTime.Location())

	return time
}

func GetClassesToday(config ClassConfig, weekDay string) []ClassToday {
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

		time := get_time_from_hh_mm(todayTime.Time)

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
