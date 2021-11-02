.PHONY: help
help:
	@echo '                                                                          '
	@echo 'Makefile for bsonfilter                                                   '
	@echo '                                                                          '
	@echo 'Usage:                                                                    '
	@echo '   make help                           show help                          '
	@echo '                                                                          '
	@echo '   make run                            测试执行                            '
	@echo '   make build                          构建二进制                           '
	@echo '                                                                          '


.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/bsonfilter.linux cmd/main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/bsonfilter.darwin cmd/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/bsonfilter.windows cmd/main.go

#run:
#	./bin/bsonfilter.darwin -i backup/data/data/data.bson -o tmp.bson -q '{"_id": 10001}'
#
.PHONY: run
run:
	go run cmd/main.go -i backup/data/data/data.bson -o backup/data_new/data/tmp.bson -q '{"_id": 1}' -v
