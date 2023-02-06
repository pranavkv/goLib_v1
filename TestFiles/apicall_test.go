package testfiles

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"testing"
	"time"

	lib "git.marketsimplified.com/Platform-3.0/GoLib-Platform"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type User struct {
	Id    int
	Name  string
	Email string
}

type ApiResponse struct {
}

type ApiRequest struct {
	Url    string
	Header http.Header
	Model  interface{}
}

func (req ApiRequest) GetHeader() http.Header {

	return http.Header{}

}

func (req ApiRequest) GetUrl() string {

	return "https://onlineaccount.cityunionbank.in/cubservices/App/Config"
}

func (req ApiRequest) GetModel() interface{} {
	jsonReq := `{
		"request": {
			"data": {},
			"appID": {{APP_ID}}
		}
	}`

	return jsonReq
}

func (res ApiResponse) CheckError() error {

	return nil
}

func (res ApiResponse) GetModel() interface{} {

	return &res
}
func TestGetMethods(t *testing.T) {

	var request ApiRequest
	var response ApiResponse

	request.Url = request.GetUrl()
	request.Model = request.GetModel()
	request.Header = request.GetHeader()

	//responseRecorder := httptest.NewRecorder()
	err := lib.GetRequest(request, response)

	if err != nil {
		t.Error(err)
	}

	// if status := responseRecorder.Code; status != http.StatusOK {
	// 	t.Errorf("Handler returns wrong status code : got %v expected %v ", status, http.StatusOK)
	// }

	//expected := `{ "key" : "val" }`

	// if responseRecorder.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected response : got %v expected  %v",
	// 		responseRecorder.Body.String(), expected)
	// }

}

func TestPostMethods(t *testing.T) {

}

func TestMysql(t *testing.T) {

	fmt.Sprintf("database.name=%s", lib.GetString("database.name"))

	lib.InitMySQL()
	fmt.Println("finished.....")

	var result User

	lib.OrmDB.Raw("SELECT id, name, email FROM USER_DETAILS WHERE id = ?", 11).Scan(&result)
	fmt.Println(result.Name)
	time.Sleep(10 * time.Minute)
}

func TestPostgres(t *testing.T) {
	fmt.Printf("db name = %s", lib.GetString("database.name"))

	lib.InitPostGreSql()

}

func TestMongoDB(t *testing.T) {
	fmt.Printf("db name = %s", lib.MongoDB)

	lib.InitMongoDB()
	dbCollection := lib.MongoDB.Collection("employee")
	fmt.Printf(" Finished connecting with Mongo db ")

	user := User{3, "arun", "email"}
	insertResult, err := dbCollection.InsertOne(context.TODO(), user)
	if err != nil {
		t.Error(err)
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
