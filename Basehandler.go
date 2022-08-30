package golib_v1

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

type RootHandler func(http.ResponseWriter, *http.Request) (GoLibResponse,error)

func init() {  
	// utils.Logger.Debug("Roothandler initialized")
}

func SessionService(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        if r.Header["Authorization"] != nil {

            token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("There was an error")
                }
                return mySigningKey, nil
            })

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            if token.Valid {
                endpoint(w, r)
            }
        } else {

            fmt.Fprintf(w, "Not Authorized")
        }
    })
}

func ProcessRequest(w http.ResponseWriter, r *http.Request) (GoLibRequest,GoLibResponse, error) {
	
	req := GoLibRequest{MsgId : GetMsgID()}
	res := GoLibResponse{MsgId : GetMsgID()}
	var resErr error
	
	body, err := ioutil.ReadAll(r.Body) // Read request body.
	Logger.Infof("Request Receievd: %s", string(body))
	if err != nil {
		resErr = NewHTTPError(nil, 404, "Request body read error.", "EGN001")
	}

	// Parse body as json.
	if err = json.Unmarshal([]byte(body), &req); err != nil {
		resErr =  NewHTTPError(err, 400, "Bad request : invalid JSON.", "EGN002")
	}

	return req,res,resErr

}

func (fn RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	res,err := fn(w,r)
	if(err != nil) {
		res.Error = err

	}
	
	resBody, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(500)
		} else {
			w.WriteHeader(200)
			w.Write(resBody)
		}

	Logger.Infof("response sent: %v", string(resBody))

}