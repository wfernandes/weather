package wunderground

import (
	"fmt"
	"net/http"

	"github.com/wfernandes/weather/wunderground/validation"
)

const apiHost = "http://api.wunderground.com"

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Wunderground struct {
	apiKey string
	client HttpClient
}

func New(apiKey string, client HttpClient) *Wunderground {

	if client == nil {
		client = http.DefaultClient
	}

	return &Wunderground{
		apiKey: apiKey,
		client: client,
	}
}

func (w *Wunderground) Conditions(zipCode string) error {

	err := validation.Zip(zipCode)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("/api/%s/conditions/q/%s.json", w.apiKey, zipCode)
	url := apiHost + path
	_, err = w.makeRequest(url)
	if err != nil {
		return err
	}

	return nil
}

func (w *Wunderground) Client() HttpClient {
	return w.client
}

func (w *Wunderground) makeRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return w.client.Do(req)
}

//
//println("there was no error...huh....")
//var data map[string]interface{}
//decoder := json.NewDecoder(resp.Body)
//err = decoder.Decode(&data)
//if err != nil {
//log.Panic("Error decoding json response")
//}
//
//for k, v := range data {
//if k == "current_observation" {
//for k2, v2 := range v.(map[string]interface{}) {
//fmt.Printf("\n%s : %v\n", k2, v2)
//}
//}
//}
