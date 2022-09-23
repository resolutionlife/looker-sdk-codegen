package rtl

import (
	"fmt"
	"strings"
)

type ResponseError struct {
	StatusCode int
	Body       []byte
	Err        error
}

func (e ResponseError) Error() string {
	return e.Err.Error()
}

type Error struct {
	Message          string `json:"message"`           // Error details
	DocumentationUrl string `json:"documentation_url"` // Documentation link
}

func (e Error) Error() string {
	return e.Message
}

type ValidationError struct {
	Message          string                   `json:"message"`           // Error details
	Errors           *[]ValidationErrorDetail `json:"errors,omitempty"`  // Error detail array
	DocumentationUrl string                   `json:"documentation_url"` // Documentation link
}

type ValidationErrorDetail struct {
	Field            *string `json:"field,omitempty"`   // Field with error
	Code             *string `json:"code,omitempty"`    // Error code
	Message          *string `json:"message,omitempty"` // Error info message
	DocumentationUrl string  `json:"documentation_url"` // Documentation link
}

func (e ValidationError) Error() string {
	if e.Errors == nil {
		return e.Message
	}

	var errSlice []string
	for _, m := range *e.Errors {
		// need to check field and message are not nil
		errSlice = append(errSlice, fmt.Sprintf("an error has occured with field %s. %s\n", *m.Field, *m.Message))
	}
	return strings.Join(errSlice, ",")
}
