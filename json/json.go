// Package json is package for functions related to JSON operation.
package json

import (
	"encoding/json"
	"errors"
	"strings"
)

var errInvalidRequest = errors.New("invalid request")

// ConvertRequest is function to convert from request body to specified struct.
func ConvertRequest(body []byte, v interface{}) error {
	s := string(body)
	s = strings.Replace(s, " ", "", -1)
	if len(s) < 10 {
		return errInvalidRequest
	}

	err := json.Unmarshal(body, &v)
	if err != nil {
		return err
	}

	return nil
}
