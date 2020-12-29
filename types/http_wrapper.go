package types

type ErrorResponse struct {
	Message string `json:"message"`
	// TranslatableMessage []map[string]string `json:"translatableMessage"`
	StatusCode string `json:"statusCode"`
	Expected   string `json:"expected"`
	Type       string `json:"type"`
}
