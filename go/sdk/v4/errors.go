package v4

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/looker-open-source/sdk-codegen/go/rtl"
)

func parseErr(err error) error {
	re, ok := err.(rtl.ResponseError)
	if !ok {
		return err
	}

	if len(re.Body) == 0 {
		return fmt.Errorf("response error. status=%d. error parsing body", re.StatusCode)
	}

	var e error
	switch re.StatusCode {
	case http.StatusUnprocessableEntity:
		// Status 422 returns a json payload of type ValidationError
		e = new(ValidationError)
	default:
		// All other status codes return a json payload of type Error
		e = new(Error)
	}
	if err := json.Unmarshal(re.Body, &e); err != nil {
		return fmt.Errorf("error unmarshalling body with status: %d, body:%s, error:%s", re.StatusCode, re.Body, err.Error())
	}
	re.Err = e

	return re
}

func (e Error) Error() string {
	return e.Message
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
