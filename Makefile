init:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/ue4-process-manager ue4-lspm/*
	GOOS=linux GOARCH=amd64 go build -o bin/linux/ue4-process-manager ue4-lspm/*

run-mac:
	./bin/mac/ue4-process-manager -total-processes=4 -path=ls

run:
	./bin/linux/ue4-process-manager