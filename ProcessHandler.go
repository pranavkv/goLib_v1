/*
Author: Pranav KV
Mail: pranavkvnambiar@gmail.com
*/
package golib_v1

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var mySigningKey = []byte(GetString("jwt.sign.key"))

func setHeader(ctx *gin.Context) {

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Expose-Headers", "Set-Cookie")
	ctx.Header("Access-Control-Allow-Headers", "cache-control, content-type, DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Set-Cookie,origin,accept")
}

func Process[Req GoLibRequest, Resp GoLibResponse](processFunc func(request GoLibRequest, response *GoLibResponse) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		request := GoLibRequest{MsgId: GetMsgID()}
		response := GoLibResponse{MsgId: GetMsgID()}

		if ctx.Request.Method == "OPTIONS" {
			setHeader(ctx)
		}

		ctx.Header("Content-Type", "application/json")
		if err := ctx.ShouldBindJSON(&request); err != nil {
			Logger.Errorf("Decoding body failed: %v", err)
			resErr := NewHTTPError(err, 400, "Bad request : invalid JSON.", "EGN002")
			response.Error = resErr.Error()
			ctx.BindJSON(response)
			return
		}

		resErr := processFunc(request, &response)
		if resErr != nil {
			Logger.Errorf("Unable to process request: %v", resErr)
			response.Error = resErr.Error()
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.Status(http.StatusOK)
		}

		ctx.BindJSON(response)
	}
}

func ValidateSession[Req GoLibRequest, Resp GoLibResponse](processFunc func(request GoLibRequest, response *GoLibResponse) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		request := GoLibRequest{MsgId: GetMsgID()}
		response := GoLibResponse{MsgId: GetMsgID()}

		ctx.Header("Content-Type", "application/json")

		token, err := jwt.Parse(ctx.GetHeader("Authorization"), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid token")
			}
			return mySigningKey, nil
		})

		if err != nil || !token.Valid {
			response.Error = NewHTTPError(nil, 401, "UnAuthorized request received.", "EGN003").Error()
			ctx.BindJSON(response)
			return
		}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			Logger.Errorf("Decoding body failed: %v", err)
			resErr := NewHTTPError(err, 400, "Bad request : invalid JSON.", "EGN002")
			response.Error = resErr.Error()
			ctx.BindJSON(response)
			return
		}

		resErr := processFunc(request, &response)
		if resErr != nil {
			Logger.Errorf("Unable to process request: %v", resErr)
			response.Error = resErr.Error()
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.Status(http.StatusOK)
		}

		ctx.BindJSON(response)
	}

}
