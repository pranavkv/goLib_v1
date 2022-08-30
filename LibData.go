package golib_v1

type GoLibRequest struct {
	Data         map[string]interface{} `json:"data"`
	AppId        string                 `json:"appID"`
	MsgId        string                 `json:"msgID"`
	Access_token string                 `json:"access_token"`
}

type GoLibResponse struct {
	Data         map[string]interface{} `json:"data"`
	AppId        string                 `json:"appID"`
	MsgId        string                 `json:"msgID"`
	Error		 error					`json:"error"`
}
