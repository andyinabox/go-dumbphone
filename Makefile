default: cli

setup:
	sh ./bin/setup.sh

test: bindata
	go test ./...

bindata:
	go-bindata -pkg data -o ./bin/data/data.go ./bin/data/...

cli: bindata
	go build -o ./bin ./cmd/dumbp

api:
	go build -o ./bin ./cmd/api
