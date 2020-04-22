package main

import "log"

func main() {
	configuration, err := NewProcessManagerServerConfiguration("config/configuration.json")
	if err != nil {
		log.Fatal(err)
	}
	server := NewProcessManagerServer("0.0.0.0", 8086, configuration)
	server.Start()
}
