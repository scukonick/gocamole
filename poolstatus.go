package main

import (
	"encoding/json"
	"io"
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
