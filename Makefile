all: build run

delete: 
	rm -rf bin

build:
	$(MAKE) delete
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/spindle-linux main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/spindle-windows.exe main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/spindle-darwin main.go   

run:
	./bin/spindle server --addr=0.0.0.0 --port=8081

ensure-compile-daemon:
	@which go > /dev/null || (echo "Error: Go is not installed or not in PATH" && exit 1)
	@which CompileDaemon > /dev/null || (echo "Installing CompileDaemon..." && go install github.com/githubnemo/CompileDaemon@latest)

serve:
	CompileDaemon -build="go build -o ./bin/spindle main.go" -command="./bin/spindle server --addr=0.0.0.0 --port=8081"