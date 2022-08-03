package rootHandler

import (
	"fmt"
	"log"
	"net/http"
	errors "msflib/libError"
)

type RootHandler func(http.ResponseWriter, *http.Request) error

func init() {  
    fmt.Println("rootHandler package initialized")
}

func (fn RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err == nil {
		return
	}
	
	log.Printf("An error accured: %v", err)

	clientError, ok := err.(errors.ClientError)
	if !ok {
		w.WriteHeader(500) // return 500 Internal Server Error.
		return
	}

	body, err := clientError.ResponseBody()
	if err != nil {
		log.Printf("An error accured: %v", err)
		w.WriteHeader(500)
		return
	}
	status, headers := clientError.ResponseHeaders() // Get http status code and headers.
	for k, v := range headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(status)
	w.Write(body)
}