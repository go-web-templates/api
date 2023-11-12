package interfaces

type ValidateError struct {
	Tag     string `json:"tag"`
	Field   string `json:"field"`
	Value   any    `json:"value"`
	Message string `json:"message"`
}

type JsonValidator interface {
	Validate(data interface{}) (bool, []ValidateError)
}
