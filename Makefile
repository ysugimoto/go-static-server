all: darwin linux windows

darwin:
	GOOS=darwin GOARC=amd64 go build -o build/darwin/go-static-server main.go
linux:
	GOOS=linux GOARC=amd64 go build -o build/linux/go-static-server main.go
windows:
	GOOS=windows GOARC=amd64 go build -o build/windows/go-static-server.exe main.go
