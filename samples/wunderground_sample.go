package main
import (
	"github.com/wfernandes/weather/wunderground"
	"fmt"
)

const (
	apiKey = "API_KEY"
	zip = "ZIP_CODE"
)

func main() {

	conf := wunderground.DefaultConfig(apiKey)
	w := wunderground.New(conf)

	conditions, err := w.Conditions(zip)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nResponse:\n %+v\n", conditions.Response)
	fmt.Printf("\nConditions:\n %+v\n", conditions.Condition)
}