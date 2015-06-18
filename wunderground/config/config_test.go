package config_test

import (
	"github.com/wfernandes/weather/wunderground/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	It("requires an api key and api host", func() {
		conf := &config.Config{}
		err := conf.Validate()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Required: API Key, API Host\n"))
	})

})
