package main

import (
	"strings"
	"testing"
)

func TestNewPoolStatus(t *testing.T) {
	json_string := "{\"pool\":\"www\",\"process manager\":\"dynamic\",\"start time\":1447967095,\"start since\":462727,\"accepted conn\":155,\"listen queue\":0,\"max listen queue\"    :0,\"listen queue len\":128,\"idle processes\":1,\"active processes\":1,\"total processes\":2,\"max active processes\":1,\"max children reached\":0,\"slow requests\":0}"

	r := strings.NewReader(json_string)

	poolStatus, err := NewPoolStatus(r)

	if err != nil {
		t.Error("Error should be nil when creating new PoolStatus from correct json")
	}
	if poolStatus.Pool != "www" {
		t.Error("Pool name is wrong")
	}
	if poolStatus.ListenQueueLen != 128 {
		t.Error("Pool listen queue len is wrong")
	}
}

func TestNewPoolStatusIncorrectJson(t *testing.T) {
	json_string := "{\"sadqwdwq: 123}"
	r := strings.NewReader(json_string)

	_, err := NewPoolStatus(r)
	if err == nil {
		t.Error("If json string is incorrect should return error")
	}
}
