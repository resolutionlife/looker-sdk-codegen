package v4

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/looker-open-source/sdk-codegen/go/rtl"
	"github.com/pkg/errors"
)

var (
	// ErrNotFound occurs when a requested resource cannot be found.
	ErrNotFound = errors.New("not found")
)

// Error ensures that the Error type complies to the error interface.
func (e Error) Error() string {
	return e.Message
}

// Error ensures that the ValidationError type complies to the error interface.
func (e ValidationError) Error() string {
	if e.Errors == nil {
		return e.Message
	}

	var errSlice []string
	for _, m := range *e.Errors {
		if m.Field != nil && m.Message != nil {
			errSlice = append(errSlice, fmt.Sprintf("'%s' (%s)", *m.Field, *m.Message))
		}
	}
	return fmt.Sprintf("validation error on fields: %s", strings.Join(errSlice, ", "))
}

// deseraliseError decodeds the error message depending on the error response status.
func deserialiseError(err error) error {
	if err == nil {
		return nil
	}

	re, ok := err.(rtl.ResponseError)
	if !ok {
		return err
	}

	switch re.StatusCode {
	case http.StatusUnprocessableEntity:
		e := new(ValidationError)
		if err := json.Unmarshal(re.Body, &e); err != nil {
			break
		}
		return e
	case http.StatusNotFound:
		e := new(Error)
		if err := json.Unmarshal(re.Body, &e); err != nil {
			break
		}
		return errors.Wrap(ErrNotFound, e.Message)
	default:
		e := new(Error)
		if err := json.Unmarshal(re.Body, &e); err != nil {
			break
		}
		return e
	}

	return fmt.Errorf("error unmarshalling body with status: %d, body:%s, error:%s", re.StatusCode, re.Body, err.Error())
}
