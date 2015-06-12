package validation_test

import (
	"github.com/wfernandes/weather/wunderground/validation"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Zip", func() {

	It("returns error when empty", func() {

		err := validation.Zip("")
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal(validation.ErrInvalidZip))
	})

	It("returns error for invalid zip code", func() {
		err := validation.Zip("123")
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal(validation.ErrInvalidZip))

		err = validation.Zip("1234#")
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal(validation.ErrInvalidZip))

		err = validation.Zip("12345")
		Expect(err).ToNot(HaveOccurred())
	})
})
