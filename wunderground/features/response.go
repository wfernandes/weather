package features

type Response struct {
	Version        string       `json:"version"`
	TermsOfService string       `json:"termsofService"`
	Error          *ErrResponse `json:"error,omitempty"`
}

type ErrResponse struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
