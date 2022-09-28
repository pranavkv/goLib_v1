/*
Author: Pranav KV
Mail: pranavkvnmabiar@gmail.com
*/
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	lib "github.com/pranavkv/golib_v1"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type User struct {
	Id    int
	Name  string
	Email string
}

type ApiRequest struct {
}

func (req ApiRequest) GetUrl() string {

	return "url"
}

func (req ApiRequest) GetHeader() http.Header {

	var header http.Header
	return header

}

func (req ApiRequest) GetModel() interface{} {
	return &req
}

type ApiResponse struct {
}

func (res ApiResponse) CheckError() error {

	return nil
}

func (res ApiResponse) GetModel() interface{} {

	return &res
}

func testApiCall() {
	var req ApiRequest
	var res ApiResponse
	lib.PostRequest(req, res)
}

func testMySQL() {

	fmt.Println()
	out := fmt.Sprintf("database.name=%s", lib.GetString("database.name"))
	fmt.Println(out)

	lib.InitMySQL()
	fmt.Println("finished.....")

	var result User

	lib.OrmDB.Raw("SELECT id, name, email FROM USER_DETAILS WHERE id = ?", 11).Scan(&result)
	fmt.Println(result.Name)
	time.Sleep(10 * time.Minute)

}

func ProcessLogin(request lib.GoLibRequest, response *lib.GoLibResponse) error {

	lib.Logger.Info("requerst received")

	response.Data = make(map[string]interface{})
	response.Data["name"] = "pranav"
	return nil
}

func test_rest_service() {

	sm := http.NewServeMux()
	sm.HandleFunc("/login", lib.Process(ProcessLogin))
	sm.Handle("/metrics", promhttp.Handler())

	s := &http.Server{
		Addr:         ":9091",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			lib.Logger.Error(err)
		}
	}()

	// make a new channel to notify on os interrupt of server (ctrl + C)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// This blocks the code until the channel receives some message
	sig := <-sigChan
	lib.Logger.Infof("Received terminate, graceful shutdown ", sig)
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}

func main() {

	lib.InitLog("GOTEST", "pranav-PC")
	lib.InitConfig(".")

	test_rest_service()
	// testMySQL()

}
