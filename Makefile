export GO111MODULE=on
export GOARCH=amd64
export GOOS=linux

.PHONY: build
build: bin/main

bin/main: main.go pkg/**/*.go config/aws.go go.mod go.sum
	go build -o bin/main

.PHONY: emulate
emulate:
	docker build -t resize-api .
	docker run -d -p 9000:8080 --name resize-api resize-api:latest /main

.PHONY: down
down:
	docker stop resize-api
	docker rm resize-api
	docker rmi resize-api

.PHONY: reemulate
reemulate:
	@make down
	@make emulate

.PHONY: deploy
deploy:
	lambroll deploy

.PHONY: invoke
invoke:
	curl -X POST -H 'Content-type: image/png' --data-binary "@./tmp/original/gopher.png" "http://localhost:9000/2015-03-31/functions/function/invocations"

.PHONY: execute
post:
	curl -X POST -H 'Content-type: image/png' --data-binary "@./tmp/original/gopher.png" https://3d8r7a230b.execute-api.ap-northeast-1.amazonaws.com/default/resize-api

.PHONY: execute-node
execute-node:
	@node execute/node/main.mjs

.PHONY: execute-php
execute-php:
	@php execute/php/main.php