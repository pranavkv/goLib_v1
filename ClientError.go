package golib_v1

type ClientError interface {

	Error() string
	
	ResponseBody() ([]byte, error)

	ResponseHeaders() (int, map[string]string)
}
