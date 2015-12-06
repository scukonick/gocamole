package main

import (
	"strings"
	"testing"
)

func TestNewConfig(t *testing.T) {
	json_string := `{  
  "pools":[  
    {  
      "socket":"127.0.0.1:9000",
      "uri":"/php-status?json"
    },
    {  
      "socket":"127.0.0.1:9001",
      "uri":"/php-status?json"
    }
  ]
}`

	r := strings.NewReader(json_string)
	config := NewConfig(r)
	if config.Pools[0].Uri != "/php-status?json" {
		t.Error("Config uri was parsed incoorrectly")
	}
	if config.Pools[0].Socket != "127.0.0.1:9000" {
		t.Error("Config socket was parsed incoorrectly")
	}
	if config.Pools[1].Socket != "127.0.0.1:9001" {
		t.Error("Config socket was parsed incoorrectly")
	}
	if len(config.Pools) != 2 {
		t.Error("Incorrect pools count")
	}
}
