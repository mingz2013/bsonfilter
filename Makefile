build:
	#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/bsonfilter.linux
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/bsonfilter.darwin
	#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/bsonfilter.windows

#run:
#	./bin/bsonfilter.darwin -i backup/data/data/data.bson -o tmp.bson -q '{"_id": 10001}'
#
run:
	go run main.go -i backup/data/data/data.bson -o tmp.bson -q '{"_id": 10001}'
