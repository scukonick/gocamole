package main

import (
	"bufio"
	"github.com/scukonick/go-fastcgi-client"
	"net/http"
)

// Pool presents one of pool in the
// list of monitoring pools
type Pool struct {
	Socket string `json:"socket"`
	Uri    string `json:"uri"`

	Status    *PoolStatus            `json:"-"`
	Available bool                   `json:"-"`
	fcgi      *fcgiclient.FCGIClient `json:"-"`
}

func NewPool(socket string, uri string) *Pool {
	pool := new(Pool)
	pool.Socket = socket
	pool.Uri = uri
	pool.Available = false
	return pool
}

func (pool *Pool) connect() error {
	fcgi, err := fcgiclient.New(pool.Socket)
	if err != nil {
		//log.Printf("ERROR: Ooops, could not connect to fcgi server: %v\n", err)
		return err
	}
	pool.fcgi = fcgi
	return nil
}

func (pool *Pool) askStatus() error {
	request, err := http.NewRequest("GET", pool.Uri, nil)
	if err != nil {
		//log.Printf("ERROR: Could not create request: %v\n", err)
		return err
	}

	http_response, err := pool.fcgi.DoHTTPRequest(request, "/tmp/test.php")
	if err != nil {
		//log.Printf("ERROR: Could not execute http request: %v\n", err)
		return err
	}
	r := bufio.NewReader(http_response.Body)
	pool.Status, err = NewPoolStatus(r)
	if err != nil {
		//log.Printf("ERROR: Could not read pool status from json: %v\n", err)
		return err
	}
	pool.Available = true
	return nil
}

func (pool *Pool) UpdateStatus() error {
	err := pool.connect()
	if err != nil {
		return err
	}
	err = pool.askStatus()
	//log.Printf("Pool availability: %v", pool.Available)
	return err
}
