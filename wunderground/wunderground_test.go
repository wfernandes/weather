package wunderground_test

import (
	"github.com/wfernandes/weather/wunderground"

	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/wfernandes/weather/wunderground/config"
)

var _ = Describe("Wunderground", func() {

	var (
		apiKey string
		conf   *config.Config
	)

	BeforeEach(func() {
		apiKey = "some_api_key"
		conf = wunderground.DefaultConfig(apiKey)
	})

	Context("Default Config", func() {
		It("sets the api key and defaults for the HttpClient and api host", func() {
			Expect(conf.ApiKey).To(Equal(apiKey))
			Expect(conf.ApiHost).To(Equal("http://api.wunderground.com"))
		})
	})

	Context("Initialization", func() {
		It("panics if no api key is set in the config", func() {
			conf.ApiKey = ""
			Expect(func() { wunderground.New(conf) }).To(Panic())
		})

		It("panics if no api host is set in the config", func() {
			conf.ApiHost = ""
			Expect(func() { wunderground.New(conf) }).To(Panic())
		})
	})

	Context("Conditions", func() {

		var (
			fakeServer *MockWunderground
			w          *wunderground.Wunderground
		)

		BeforeEach(func() {
			fakeServer = NewMockWunderground()
			conf.ApiHost = fakeServer.Server.URL

			w = wunderground.New(conf)
		})

		AfterEach(func() {
			fakeServer.Close()
		})

		It("builds the correct url for the wunderground api", func() {
			zipCode := "80516"
			fakeServer.FakeResponseBody = `{"response":{}}`
			_, err := w.Conditions("80516")
			Expect(err).ToNot(HaveOccurred())

			Expect(fakeServer.ReceivedReq.Method).To(Equal("GET"))
			path := fmt.Sprintf("/api/%s/conditions/q/%s.json", apiKey, zipCode)
			Expect(fakeServer.ReceivedReq.URL.Path).To(Equal(path))

		})

		It("returns an error for invalid api key", func() {
			fakeServer.FakeResponseBody = InvalidApiKeyResponse
			_, err := w.Conditions("80303")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("keynotfound: this key does not exist"))
		})

		It("returns an error for bad zip code", func() {
			fakeServer.FakeResponseBody = InvalidZipCodeResponse
			_, err := w.Conditions("00000")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("querynotfound: No cities match your search query"))
		})

		It("returns conditions for successful request", func() {
			fakeServer.FakeResponseBody = ValidResponse
			cond, err := w.Conditions("80516")

			Expect(err).ToNot(HaveOccurred())
			Expect(cond).ToNot(BeNil())
		})

	})

})
