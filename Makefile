build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/bsonfilter.linux
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/bsonfilter.darwin
