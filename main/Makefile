BINARY_NAME=test

run:
	go run test_lib.go


build:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/${BINARY_NAME}-linux-arm test_lib.go
	GOOS=linux GOARCH=arm64 go build -o bin/${BINARY_NAME}-linux-arm64 test_lib.go
	GOOS=freebsd GOARCH=386 go build -o bin/${BINARY_NAME}-freebsd-386 test_lib.go    
    GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -o bin/${BINARY_NAME}-windows-386 test_lib.go

all: build run

clean: 
    GOOS=linux GOARCH=arm go clean rm bin/${BINARY_NAME}-linux-arm
