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


exec := go run cmd/main.go -i backup/data/data/data.bson -o tmp.bson -v -q


# $$ 为转义字符$
.PHONY: run
run:
	$(exec) '{"_id": 1}'
	$(exec) '{"_id": {"$$eq": 1}}'
	$(exec) '{"_id": {"$$gt": 1}}'
	$(exec) '{"_id": {"$$gte": 1}}'
	$(exec) '{"_id": {"$$in": [1,2,3]}}'
	$(exec) '{"_id": {"$$lt": 3}}'
	$(exec) '{"_id": {"$$lte": 3}}'
	$(exec) '{"_id": {"$$ne": 1}}'
	$(exec) '{"_id": {"$$in": [1]}}'
	$(exec) '{"$$and": [{"_id": 1}, {"a": {"$$exists": 0}}]}'
	$(exec) '{"$$or": [{"_id": 1}, {"a": {"$$exists": 0}}]}'
	$(exec) '{"$$not": {"$$or": [{"_id": 1}, {"a": {"$$exists": 0}}]}}'
	$(exec) '{"$$nor": [{"_id": 1}, {"a": {"$$exists": 0}}]}'
	$(exec) '{"_id": {"$$exists": 1}}'
