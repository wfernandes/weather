package features_test

import (
	"github.com/wfernandes/weather/wunderground/features"

	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {

	It("unmarshalles JSON correctly", func() {

		jsonStr := `{"version":"0.1","termsofService":"http://www.wunderground.com/weather/api/d/terms.html","features":{"conditions":1},"error":{"type":"querynotfound","description":"No cities match your search query"}}`

		var condRes *features.Response
		err := json.Unmarshal([]byte(jsonStr), &condRes)
		Expect(err).ToNot(HaveOccurred())

		Expect(condRes.Version).To(Equal("0.1"))
		Expect(condRes.TermsOfService).To(Equal("http://www.wunderground.com/weather/api/d/terms.html"))
		Expect(condRes.Error.Type).To(Equal("querynotfound"))
		Expect(condRes.Error.Description).To(Equal("No cities match your search query"))
	})
})
