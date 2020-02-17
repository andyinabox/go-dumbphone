default: test build-cli

setup:
	sh ./bin/setup.sh

test: bindata
	go test ./...

bindata:
	go-bindata -pkg data -o ./bin/data/data.go ./bin/data/...

build-cli: bindata
	go build -o ./bin ./cmd/dumbp

build-api:
	go build -o ./bin ./cmd/api

install: build-cli
	sudo cp bin/dumbp /usr/local/bin/dumbp
