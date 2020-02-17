default: cli

bindata: install-bindata
	go-bindata -pkg data -o ./data/data.go ./bin/data/...

cli: bindata
	go build -o ./bin ./cmd/dumbp

api:
	go build -o ./bin ./cmd/api

install-bindata:
	go get -u github.com/go-bindata/go-bindata/...