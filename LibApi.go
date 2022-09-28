/*
Author: Pranav KV
Mail: pranavkvnambiar@gmail.com
*/
package golib_v1

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ResponseObject interface {
	CheckError() error
	GetModel() interface{}
}

type RequestObject interface {
	GetUrl() string
	GetHeader() http.Header
	GetModel() interface{}
}

func PostRequest(req RequestObject, res ResponseObject) error {

	var resErr error
	client := &http.Client{}
	reqBody, err := json.Marshal(res.GetModel())
	if err != nil {
		Logger.Error("Encoding request failed: %v", err)
		resErr = NewHTTPError(err, 400, "Unable to process the request", "EGN004")
	}
	reqUrl := req.GetUrl()
	httpReq, err := http.NewRequest("POST", reqUrl, strings.NewReader(string(reqBody)))
	if err != nil {
		Logger.Error("Bad Request: %v", err)
		resErr = NewHTTPError(err, 400, "Unable to process the request", "EGN004")
	}

	if req.GetHeader() != nil {
		httpReq.Header = req.GetHeader()
	}

	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(httpReq)
	if err != nil {
		Logger.Error("Unable to connect to: %v", err.Error())
		resErr = NewHTTPError(err, 400, "Unable to process the request", "EGN004")
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		Logger.Error("Invalid API response: %v", err.Error())
		resErr = NewHTTPError(err, 400, "Unable to process the request", "EGN004")
		return resErr
	}

	return resErr

}

// func PostRequest[T Responser](url string, reqBody []byte, headerMap http.Header) error {

// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", url, strings.NewReader(string(reqBody)))
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}

// 	if headerMap != nil {
// 		req.Header = headerMap
// 	}

// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("Content-Type", "application/json")
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}

// 	defer resp.Body.Close()
// 	var respInfo T
// 	err = json.NewDecoder(resp.Body).Decode(&respInfo)
// 	if err != nil {
// 		return err
// 	}

// 	if err := respInfo.CheckError(); err != nil {
// 		return err
// 	}

// 	return nil

// }

// func GetRequest[T Responser](url string, headerMap http.Header) error {

// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}

// 	if headerMap != nil {
// 		req.Header = headerMap
// 	}

// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("Content-Type", "application/json")
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}

// 	defer resp.Body.Close()
// 	var respInfo T
// 	err = json.NewDecoder(resp.Body).Decode(&respInfo)
// 	if err != nil {
// 		return err
// 	}

// 	if err := respInfo.CheckError(); err != nil {
// 		return err
// 	}

// 	return nil

// }
