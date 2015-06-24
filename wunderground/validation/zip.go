package validation

import (
	"errors"
	"regexp"
)

const (
	ErrInvalidZip = "Invalid ZipCode"
)

func Zip(zipCode string) error {

	zipCodeRegEx := regexp.MustCompile("^[0-9]{5}$")
	if !zipCodeRegEx.MatchString(zipCode) {
		return errors.New(ErrInvalidZip)
	}

	return nil
}
