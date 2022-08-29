package BaseHandler

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	errors "github.com/pranavkv/golib_v1/libError"
	utils "github.com/pranavkv/golib_v1/LibUtils"
	data "github.com/pranavkv/golib_v1/LibData"
)

type RootHandler func(http.ResponseWriter, *http.Request) (data.GoLibResponse,error)

func init() {  
	// utils.Logger.Debug("Roothandler initialized")
}

func ProcessRequest(w http.ResponseWriter, r *http.Request) (data.GoLibRequest,data.GoLibResponse, error) {
	
	req := data.GoLibRequest{MsgId : utils.GetMsgID()}
	res := data.GoLibResponse{MsgId : utils.GetMsgID()}
	var resErr error
	
	body, err := ioutil.ReadAll(r.Body) // Read request body.
	utils.Logger.Infof("Request Receievd: %s", string(body))
	if err != nil {
		resErr = errors.NewHTTPError(nil, 404, "Request body read error.", "EGN001")
	}

	// Parse body as json.
	if err = json.Unmarshal([]byte(body), &req); err != nil {
		resErr =  errors.NewHTTPError(err, 400, "Bad request : invalid JSON.", "EGN002")
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

	utils.Logger.Infof("response sent: %v", string(resBody))

}