/*
Author: Pranav KV
Mail: pranavkvnmabiar@gmail.com
*/
package golib_v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte(GetString("jwt.sign.key"))

func Process[Req GoLibRequest, Resp GoLibResponse](processFunc func(request GoLibRequest, response *GoLibResponse) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		request := GoLibRequest{MsgId: GetMsgID()}
		response := GoLibResponse{MsgId: GetMsgID()}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			Logger.Errorf("Decoding body failed: %v", err)
			resErr := NewHTTPError(err, 400, "Bad request : invalid JSON.", "EGN002")
			response.Error = resErr
			json.NewEncoder(w).Encode(response)
			return
		}

		resErr := processFunc(request, &response)
		if resErr != nil {
			Logger.Errorf("Unable to process request: %v", resErr)
			response.Error = resErr
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(200)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Encoding response failed: %v", err)
		}
	}

}

func ValidateSession[Req GoLibRequest, Resp GoLibResponse](processFunc func(request GoLibRequest, response *GoLibResponse) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		request := GoLibRequest{MsgId: GetMsgID()}
		response := GoLibResponse{MsgId: GetMsgID()}

		token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid token")
			}
			return mySigningKey, nil
		})

		if err != nil || !token.Valid {
			response.Error = NewHTTPError(nil, 401, "UnAuthorized request received.", "EGN003")
			json.NewEncoder(w).Encode(response)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			Logger.Errorf("Decoding body failed: %v", err)
			resErr := NewHTTPError(err, 400, "Bad request : invalid JSON.", "EGN002")
			response.Error = resErr
			json.NewEncoder(w).Encode(response)
			return
		}

		resErr := processFunc(request, &response)
		if resErr != nil {
			Logger.Errorf("Unable to process request: %v", resErr)
			response.Error = resErr
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(200)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Encoding response failed: %v", err)
		}
	}

}
