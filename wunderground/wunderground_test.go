package wunderground_test

import (
	"github.com/wfernandes/weather/wunderground"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"time"
	"fmt"
	"errors"
)

var _ = Describe("Wunderground", func() {

	var (
		apiKey string
		client *http.Client
	)

	BeforeEach(func(){
		apiKey = "some_api_key"
	})

	It("requires apiKey and client for new wunderground service", func() {

		apiKey := "some_api_key"
		client = &http.Client{
			Timeout: time.Minute,
		}

		w := wunderground.New(apiKey, client)
		Expect(w).ToNot(BeNil())
		Expect(w.Client()).To(Equal(client))
	})

	It("sets a default client if one is not provided", func() {
		w := wunderground.New(apiKey, nil)
		Expect(w).ToNot(BeNil())
		Expect(w.Client()).To(Equal(http.DefaultClient))
	})

	Context("Conditions", func(){
		It("builds the correct url for the wunderground api", func() {
			fc := &fakeClient{
				ActualReqCh: make(chan *http.Request, 1),
			}
			w := wunderground.New(apiKey, fc)

			zipCode := "80516"
			w.Conditions(zipCode)

			var actualReq *http.Request
			Eventually(fc.ActualReqCh).Should(Receive(&actualReq))
			Expect(actualReq.Method).To(Equal("GET"))
			Expect(actualReq.URL.Scheme).To(Equal("http"))
			Expect(actualReq.URL.Host).To(Equal("api.wunderground.com"))
			path := fmt.Sprintf("/api/%s/conditions/q/%s.json", apiKey, zipCode)
			Expect(actualReq.URL.Path).To(Equal(path))
		})

		XIt("returns an error for a non existant zip code", func() {

			fakeErrorResp := buildErrorResp()
			fc := &fakeClient{}
			w := wunderground.New(apiKey, nil)
			err := w.Conditions("00000")
			Expect(err).To(HaveOccurred())

			_ = fakeErrorResp
			_ = fc

		})

	})


	XIt("makes a real request for weather conditions", func() {
		w := wunderground.New("0e6bd1b46e3d9869", nil)

		zipCode := "00000"
		w.Conditions(zipCode)

	})
})

type fakeClient struct {
	ActualReqCh chan *http.Request
	ReturnResp *http.Response
}

func (f fakeClient) Do(req *http.Request) (*http.Response, error) {
	f.ActualReqCh <- req
	return f.ReturnResp, errors.New("error making request")
}

func buildErrorResp() *http.Response {

	jsonResp := `{"response":{"version":"0.1","termsofService":"http://www.wunderground.com/weather/api/d/terms.html","features":{},"error":{"type":"querynotfound","description":"No cities match your search query"}}}`

	_ = jsonResp
	return nil
}