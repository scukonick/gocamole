package main

import (
	//"bytes"
	"bufio"
	"encoding/json"
	"github.com/scukonick/go-fastcgi-client"
	"io"
	"log"
	"net/http"
	//"strings"
)

// PoolStatus presents PHP-FPM status information
// about one pool
type PoolStatus struct {
	Pool               string `json:"pool"`
	Manager            string `json:"process manager"`
	StartTime          int    `json:"start time"`
	StartSince         int    `json:"start since"`
	AcceptedConn       int    `json:"accepted conn"`
	ListenQueue        int    `json:"listen queue"`
	MaxListenQueue     int    `json:"max listen queue"`
	ListenQueueLen     int    `json:"listen queue len"`
	IdleProcesses      int    `json:"idle processes"`
	ActiveProcesses    int    `json:"active processes"`
	TotalProcesses     int    `json:"total processes"`
	MaxActiveProcesses int    `json:"max active processes"`
	MaxChildrenReached int    `json:"max children reached"`
	SlowRequests       int    `json:"slow requests"`
}

func NewPoolStatus(r io.Reader) (*PoolStatus, error) {
	poolStatus := new(PoolStatus)

	decoder := json.NewDecoder(r)
	err := decoder.Decode(poolStatus)
	return poolStatus, err
}

func main() {
	log.Printf("Ahahaha, %v", "Lol")
	fcgi, err := fcgiclient.New("127.0.0.1", 9000)
	if err != nil {
		log.Fatalf("Ooops, could not connect to fscgi server: %v", err)

	}

	request, err := http.NewRequest("GET", "http://google.com/php-status?json", nil)
	request.Header["Host"] = append(request.Header["Host"], "google.com")
	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	http_response, err := fcgi.DoHTTPRequest(request, "/tmp/test.php")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Printf("Resp code: %v", http_response.ResponseCode)
	r := bufio.NewReader(http_response.Body)

	poolStatus, err := NewPoolStatus(r)
	if err == nil {
		log.Printf("Pool name: %v", poolStatus.Pool)
	}

}
