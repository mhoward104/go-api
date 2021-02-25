package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var (
	startingTime time.Time
)

const (
	apiVersion = "1.0.0"
)

//struct containing information about the server/api
type serverStatus struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Version     string `json:"Version"`
	sync.Mutex
	UpTime string `json:"UpTime"`
}

//constructor for serverStatus
func newServerStatus() *serverStatus {
	s := &serverStatus{
		Name:        "exampleServer",
		Description: "This is a test server",
		UpTime:      "00:00:00",
		Version:     apiVersion,
	}

	return s
}

//called with `/info`
func (s *serverStatus) getInfo(w http.ResponseWriter, r *http.Request) {
	//get uptime of the server
	totalUpTime := time.Now().Sub(startingTime)

	s.Lock()
	s.UpTime = time.Time{}.Add(totalUpTime).Format("15:04:05")
	s.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&s)
}

func main() {
	router := mux.NewRouter()
	//initialize object
	s := newServerStatus()
	//get start time
	startingTime = time.Now()

	//enpoints
	router.HandleFunc("/info", s.getInfo).Methods("GET")

	//start the server
	http.ListenAndServe(":8000", router)

}
