package main

func main() {
	server := NewProcessManagerServer("0.0.0.0", 8086)
	server.Start()
}
