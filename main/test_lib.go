/*
Author: Pranav KV
Mail: pranavkvnambiar@gmail.com
*/
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	lib "git.marketsimplified.com/Platform-3.0/GoLib-Platform"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type User struct {
	Id    int
	Name  string
	Email string
}

type ApiRequest struct {
	// Url    string
	// Header http.Header
	// Model  interface{}
}

type ApiResponse struct {
}

func (res *ApiResponse) CheckError() error {
	return nil
}

func (res *ApiResponse) GetResponse() interface{} {
	return res
}

func (req ApiRequest) GetUrl() string {

	return "https://dummy.restapiexample.com/api/v1/create"
}

func (req ApiRequest) GetHeader() http.Header {
	return http.Header{}
}

func (req ApiRequest) GetModel() map[string]interface{} {
	jsonReq := map[string]interface{}{"name": "test", "salary": "123", "age": "23"}
	return jsonReq
}

func testApiCall() {
	var req ApiRequest
	//var res ApiResponse

	// req.Url = req.GetUrl()
	// req.Model = req.GetModel()
	// req.Header = req.GetHeader()

	var res map[string]interface{}

	lib.PostRequest(req, &res)

	fmt.Println("Response from client call ", res)
}

func testMySQL() {

	fmt.Println()
	out := fmt.Sprintf("database.name=%s", lib.GetString("database.name"))
	fmt.Println(out)

	lib.InitMySQL()

	var result User

	lib.OrmDB.Raw("SELECT id, name, email FROM USER_DETAILS WHERE id = ?", 11).Scan(&result)
	fmt.Println(result.Name)
	time.Sleep(10 * time.Minute)

}

func testPostGreSql() {
	fmt.Printf("db name = %s", lib.GetString("database.name"))

	lib.InitPostGreSql()

	fmt.Println("Finished")
}

func testMongoDB() {
	fmt.Printf("db name = %s", lib.MongoDB)

	lib.InitMongoDB()
	dbCollection := lib.MongoDB.Collection("employee")
	fmt.Printf(" Finished connecting with Mongo db ")

	user := User{3, "arun", "email"}
	insertResult, err := dbCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
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
	testApiCall()
	//testPostGreSql()
	//testMongoDB()
	//test_rest_service()
	// testMySQL()

}
