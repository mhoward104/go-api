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

type serverStatus struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Version     string `json:"Version"`
	sync.Mutex
	UpTime string `json:"UpTime"`
}

func newServerStatus() *serverStatus {
	s := &serverStatus{
		Name:        "exampleServer",
		Description: "This is a test server",
		UpTime:      "00:00:00",
		Version:     apiVersion,
	}

	return s
}

func (s *serverStatus) getInfo(w http.ResponseWriter, r *http.Request) {
	println("Getting Server Info...")
	totalUpTime := time.Now().Sub(startingTime)

	s.Lock()
	s.UpTime = time.Time{}.Add(totalUpTime).Format("15:04:05")
	s.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&s)
}

func main() {
	router := mux.NewRouter()
	s := newServerStatus()
	startingTime = time.Now()

	router.HandleFunc("/info", s.getInfo).Methods("GET")

	http.ListenAndServe(":8000", router)

}
