package wunderground_test
import (
	"github.com/gorilla/mux"
	"net/http/httptest"
	"net/http"
)


const (
	InvalidApiKeyResponse = `{"response":{"version":"0.1","termsofService":"http://www.wunderground.com/weather/api/d/terms.html","features":{},"error":{"type":"keynotfound","description":"this key does not exist"}}}`

	InvalidZipCodeResponse = `{"response":{"version":"0.1","termsofService":"http://www.wunderground.com/weather/api/d/terms.html","features":{},"error":{"type":"querynotfound","description":"No cities match your search query"}}}`
)


type MockWunderground struct {
	Router *mux.Router
	Server *httptest.Server
	ReceivedReq *http.Request
	FakeResponseBody string
}

type handler func(resp http.ResponseWriter, req *http.Request)

func NewMockWunderground() *MockWunderground {

	w := &MockWunderground{}
	w.Router = mux.NewRouter()
	w.Router.HandleFunc("/api/{apiKey}/conditions/q/{zip}.json", w.fakeHandler)
	w.Server = httptest.NewServer(w.Router)

	return w

}

func (w *MockWunderground) fakeHandler(resp http.ResponseWriter, req *http.Request) {

	w.ReceivedReq = req
	resp.Write([]byte(w.FakeResponseBody))
}

func (w *MockWunderground) Close() {
	w.Server.Close()
}