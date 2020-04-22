package main

import (
	"log"
	"net/http"
	"strconv"
)

// ProcessManagerServer is the HTTP server that recieves requests
// to create and terminate Unreal Engine dedicated server processes.
type ProcessManagerServer struct {
	Host       string
	Port       int
	State      map[string]string
	HTTPServer *http.ServeMux
}

// NewProcessManagerServer ;
func NewProcessManagerServer(host string, port int) *ProcessManagerServer {
	return &ProcessManagerServer{
		Host:       host,
		Port:       port,
		State:      make(map[string]string),
		HTTPServer: http.NewServeMux(),
	}
}

func (pm *ProcessManagerServer) createDedicatedServerProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println(r.Method + " /create - 405")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	// Start a new instance of the configured server binary.
}

// Start starts the ProcessManager HTTP server.
func (pm *ProcessManagerServer) Start() {
	hostString := pm.Host + ":" + strconv.Itoa(pm.Port)
	pm.HTTPServer.HandleFunc("/create", pm.createDedicatedServerProcess)
	err := http.ListenAndServe(hostString, pm.HTTPServer)
	if err != nil {
		log.Fatal(err)
	}
}
