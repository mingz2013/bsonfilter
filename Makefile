build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/bsonfilter.linux cmd/main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/bsonfilter.darwin cmd/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/bsonfilter.windows cmd/main.go

#run:
#	./bin/bsonfilter.darwin -i backup/data/data/data.bson -o tmp.bson -q '{"_id": 10001}'
#
run:
	go run cmd/main.go -i backup/data/data/data.bson -o new/data/data/tmp.bson -q '{"_id": 1}'
