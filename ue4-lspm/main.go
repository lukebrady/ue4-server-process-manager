package main

import (
	"flag"
	"log"
)

func main() {
	configurationFile := flag.String("config-file", "", "Path to a LSPM configuration file.")
	totalProcesses := flag.Int("total-processes", 3, "The amount of concurrent processes that can be managed.")
	path := flag.String("path", "", "Path to the executable.")
	flag.Parse()

	var configuration *ProcessManagerServerConfiguration
	var err error
	if len(*configurationFile) > 1 {
		configuration, err = NewProcessManagerServerConfigurationFromFile("config/configuration.json")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		configuration = NewProcessManagerServerConfiguration(*path, *totalProcesses)
	}
	server := NewProcessManagerServer("0.0.0.0", 8086, configuration)
	server.Start()
}
