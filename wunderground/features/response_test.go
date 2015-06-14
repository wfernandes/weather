package features_test

import (
	"github.com/wfernandes/weather/wunderground/features"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"encoding/json"
)

var _ = Describe("Response", func() {


	It("unmarshalles JSON correctly", func() {

		jsonStr := `{"response":{"version":"0.1","termsofService":"http://www.wunderground.com/weather/api/d/terms.html","features":{"conditions":1},"error":{"type":"querynotfound","description":"No cities match your search query"}}}`
		var condRes features.ConditionsResponse
		err := json.Unmarshal([]byte(jsonStr), &condRes)
		Expect(err).ToNot(HaveOccurred())

		Expect(condRes.Response.Version).To(Equal("0.1"))
		Expect(condRes.Response.TermsOfService).To(Equal("http://www.wunderground.com/weather/api/d/terms.html"))
		Expect(condRes.Response.Error.Type).To(Equal("querynotfound"))
		Expect(condRes.Response.Error.Description).To(Equal("No cities match your search query"))
	})
})