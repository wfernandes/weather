package features

import "fmt"

type Response struct {
	Version        string       `json:"version"`
	TermsOfService string       `json:"termsofService"`
	Error          *ErrResponse `json:"error,omitempty"`
}

type ErrResponse struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type GenericResponse struct {
	Response *Response `json:"response"`
}

func (r *GenericResponse) HasError() error {
	if r.Response == nil || r.Response.Error == nil {
		return nil
	}

	return fmt.Errorf("%s: %s", r.Response.Error.Type, r.Response.Error.Description)
}
