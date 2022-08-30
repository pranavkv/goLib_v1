package golib_v1

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strings"
   )

   type CommonResponse struct{
	ErrCode int
	ErrMsg string
   }	

//    type LoginResponse struct{
// 	CommonResponse
// 	Result struct{
// 		Token  string
// 		Expire int
// 	}
// 	}

   func (resp CommonResponse) CheckError() error {
	if resp.ErrCode == 0 && resp.ErrMsg == "" {
		return nil
	}
    return fmt.Errorf("[%d]%s", resp.ErrCode,resp.ErrMsg)
   }

   type Responser interface {
	CheckError() error
   }

   func PostRequest[T Responser](url string, reqBody []byte, headerMap http.Header) (error) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(reqBody)))
	if err != nil {
	 fmt.Print(err.Error())
	}

	if headerMap != nil {
		req.Header = headerMap
	}
	
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
	 fmt.Print(err.Error())
	}

   	defer resp.Body.Close()
	   var respInfo T
	   err = json.NewDecoder(resp.Body).Decode(&respInfo)
	   if err!=nil {
		   return err
	   }
   
	   if err:=respInfo.CheckError(); err!=nil {
		   return err
	   }
   
	   return nil

	}

   func GetRequest[T Responser](url string, headerMap http.Header) (error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
	 fmt.Print(err.Error())
	}

	if headerMap != nil {
		req.Header = headerMap
	}
	
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
	 fmt.Print(err.Error())
	}

   	defer resp.Body.Close()
	   var respInfo T
	   err = json.NewDecoder(resp.Body).Decode(&respInfo)
	   if err!=nil {
		   return err
	   }
   
	   if err:=respInfo.CheckError(); err!=nil {
		   return err
	   }
   
	   return nil

	}

