BIN="svc"
BIN_WINDOWS="svc.exe"
BUILD_FLAGS="-s -w"

build:
	go build -o $(BIN)

compile:
	env GOOS=linux GOARCH=amd64 go build -o $(BIN) -ldflags $(BUILD_FLAGS)
	env GOOS=windows GOARCH=amd64 go build -o $(BIN_WINDOWS) -ldflags $(BUILD_FLAGS)

clean:
	go clean

run: 	build
	./$(BIN)

all:	build