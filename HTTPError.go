package golib_v1

import (
	"encoding/json"
	"fmt"
)
type HTTPError struct {
	Cause  error  `json:"-"`
	InfoMsg string `json:"InfoMsg"`
	InfoId string `json:"InfoId"`
	Status int    `json:"-"`
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.InfoMsg
	}
	return e.InfoMsg + " : " + e.Cause.Error()
}

// ResponseBody returns JSON response body.
func (e *HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

// ResponseHeaders returns http status code and headers.
func (e *HTTPError) ResponseHeaders() (int, map[string]string) {
	return e.Status, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

func NewHTTPError(err error, status int, infoMsg string, infoId string) error {
	return &HTTPError{
		Cause:  err,
		InfoMsg: infoMsg,
		InfoId: infoId,
		Status: status,
	}
}