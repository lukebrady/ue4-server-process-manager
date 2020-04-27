package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// ProcessManagerServer is the HTTP server that recieves requests
// to create and terminate Unreal Engine dedicated server processes.
type ProcessManagerServer struct {
	Host          string
	Port          int
	Processes     int
	State         map[string]*os.Process
	Configuration *ProcessManagerServerConfiguration
	HTTPServer    *http.ServeMux
}

// NewProcessManagerServer creates a pointer to a ProcessManagerServer object.
func NewProcessManagerServer(host string, port int, configuration *ProcessManagerServerConfiguration) *ProcessManagerServer {
	return &ProcessManagerServer{
		Host:          host,
		Port:          port,
		Processes:     0,
		Configuration: configuration,
		State:         make(map[string]*os.Process),
		HTTPServer:    http.NewServeMux(),
	}
}

func (pm *ProcessManagerServer) routeCreateDedicatedServerProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println(r.Method + " /create - 405")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	if pm.Processes < pm.Configuration.TotalProcesses {
		// Start a new instance of the configured server binary.
		process, err := createServerProcess(pm.Configuration.DedicatedServerPath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		pm.State["test"] = process
		fmt.Println(process.Pid)
		pm.Processes++
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

// Start starts the ProcessManager HTTP server.
func (pm *ProcessManagerServer) Start() {
	hostString := pm.Host + ":" + strconv.Itoa(pm.Port)
	pm.HTTPServer.HandleFunc("/create", pm.routeCreateDedicatedServerProcess)
	err := http.ListenAndServe(hostString, pm.HTTPServer)
	if err != nil {
		log.Fatal(err)
	}
}
