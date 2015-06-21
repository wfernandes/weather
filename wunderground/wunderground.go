package wunderground

import (
	"fmt"
	"net/http"

	"github.com/wfernandes/weather/wunderground/validation"
	"encoding/json"
	"io/ioutil"
	"log"
	"github.com/wfernandes/weather/wunderground/features"
	"github.com/wfernandes/weather/wunderground/config"
)

const api = "http://api.wunderground.com"

type Wunderground struct {
	apiKey string
	apiHost string
	client *http.Client
}

func New(cnf *config.Config) *Wunderground {

	err := cnf.Validate()
	if err != nil {
		panic(err)
	}

	return &Wunderground{
		apiKey: cnf.ApiKey,
		apiHost: cnf.ApiHost,
		client: http.DefaultClient,
	}
}

func DefaultConfig(apiKey string) *config.Config {
	return &config.Config{
		ApiKey: apiKey,
		ApiHost: api,
	}
}


func (w *Wunderground) Conditions(zipCode string) (*features.ConditionsResponse, error) {

	err := validation.Zip(zipCode)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/%s/conditions/q/%s.json", w.ApiHost(), w.ApiKey(), zipCode)

	bodyBytes, err := w.makeRequest(url)
	if err != nil {
		log.Print("Error making request")
		return nil, err
	}
	var cond *features.ConditionsResponse
	err = json.Unmarshal(bodyBytes, &cond)
	if err != nil {
		log.Print("Error unmarshalling condition body")
	}

	return cond, nil
}

func (w *Wunderground) ApiKey() string {
	return w.apiKey
}

func (w *Wunderground) ApiHost() string {
	return w.apiHost
}

func (w *Wunderground) makeRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := w.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	bodyBytes, err := checkRespForError(resp)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func checkRespForError(resp *http.Response) ([]byte, error) {

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Error reading response body", err)
		return nil, err
	}

	var r *features.GenericResponse
	err = json.Unmarshal(bodyBytes, &r)
	if err != nil {
		log.Print("Error unmarshalling response")
		return nil, err
	}

	return bodyBytes, r.HasError()

}