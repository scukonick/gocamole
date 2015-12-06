package main

import (
	"flag"
	"fmt"
	"log"
)

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.json", "Path to configuration file")
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		log.Fatalf("Should be on positional argument (metric)")
	}

	metric := args[0]
	config := NewConfigFromFile(configFile)

	for i, _ := range config.Pools {
		pool := &config.Pools[i]
		err := pool.UpdateStatus()
		if err != nil {
			log.Printf("ERROR: Could not update pool status: %v\n", err)
		}
	}
	result := NewResult(config.Pools)
	switch metric {
	case "available":
		fmt.Printf("%v\n", Btoi(result.Available))
	case "total_processes":
		fmt.Printf("%v\n", result.TotalProcesses)
	case "active_processes":
		fmt.Printf("%v\n", result.ActiveProcesses)
	case "accepted_conn":
		fmt.Printf("%v\n", result.AcceptedConn)
	default:
		log.Fatalf("Unknown metric")
	}
}
