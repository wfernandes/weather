package main

import (
	"fmt"
	"github.com/wfernandes/weather/wunderground"
)

const (
	apiKey = "API_KEY"
	zip    = "ZIP_CODE"
)

func main() {

	conf := wunderground.DefaultConfig(apiKey)
	w := wunderground.New(conf)

	printConditions(w)
	printHourly(w)
}

func printConditions(w *wunderground.Wunderground) {
	conditions, err := w.Conditions(zip)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nResponse:\n %+v\n", conditions.Response)
	fmt.Printf("\nConditions:\n %+v\n", conditions.Condition)
}

func printHourly(w *wunderground.Wunderground) {
	hourly, err := w.Hourly(zip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nResponse:\n %+v\n", hourly.Response)
	fmt.Printf("\nHourly:\n %+v\n", hourly.Hourly)
}
