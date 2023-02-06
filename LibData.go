/*
Author: Pranav KV
Mail: pranavkvnambiar@gmail.com
*/
package golib_v1

type GoLibRequest struct {
	Data  map[string]interface{} `json:"data"`
	AppId string                 `json:"appID"`
	MsgId string                 `json:"msgID"`
}

type GoLibResponse struct {
	Data    map[string]interface{} `json:"data"`
	AppId   string                 `json:"appID"`
	MsgId   string                 `json:"msgID"`
	Error   string                 `json:"error"`
	InfoId  string                 `json:"infoID"`
	InfoMsg string                 `json:"infoMsg"`
}
