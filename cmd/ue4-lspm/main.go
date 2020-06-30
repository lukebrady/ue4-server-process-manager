package main

import (
	"flag"
	"log"

	lspm "github.com/lukebrady/ue4-server-process-manager/internal/ue4lspm"
)

func main() {
	configurationFile := flag.String("config-file", "", "Path to a LSPM configuration file.")
	totalProcesses := flag.Int("total-processes", 3, "The amount of concurrent processes that can be managed.")
	path := flag.String("path", "", "Path to the executable.")
	flag.Parse()

	var configuration *lspm.ProcessManagerServerConfiguration
	var err error
	if len(*configurationFile) > 1 {
		configuration, err = lspm.NewProcessManagerServerConfigurationFromFile("config/configuration.json")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		configuration = lspm.NewProcessManagerServerConfiguration(*path, *totalProcesses)
	}
	server := lspm.NewProcessManagerServer("0.0.0.0", 8086, configuration)
	server.Start()
}
