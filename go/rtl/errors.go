package rtl

import (
	"fmt"
)

// ResponseError is a struct used to expose the StatusCode and Body of the error response.
type ResponseError struct {
	Status     string
	StatusCode int
	Body       []byte
}

// Error ensures that the ResponseError type complies to the error interface.
func (e ResponseError) Error() string {
	return fmt.Sprintf("response error. status=%s. error=%s", e.Status, string(e.Body))
}
