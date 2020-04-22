init:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/ue4-process-manager ue4-lspm/*
	GOOS=linux GOARCH=amd64 go build -o bin/linux/ue4-process-manager ue4-lspm/*

run:
	./bin/ue4-process-manager