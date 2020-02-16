default: cli

bindata:
	go-bindata -pkg data -o ./data/data.go ./bin/data/...

cli: bindata
	go build -o ./bin ./cmd/dumbp

api:
	go build -o ./bin ./cmd/api