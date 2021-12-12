export GO111MODULE=on
export GOARCH=amd64
export GOOS=linux

.PHONY: build
build: bin/main

bin/main: main.go pkg/**/*.go config/aws.go go.mod go.sum
	go build -o bin/main

.PHONY: emulate
emulate:
	docker build -t lambdraw .
	docker run -d -p 9000:8080 --name lambdraw lambdraw:latest /main

.PHONY: down
down:
	docker stop lambdraw
	docker rm lambdraw
	docker rmi lambdraw

.PHONY: reemulate
reemulate:
	@make down
	@make emulate

.PHONY: deploy
deploy: bin/main
	lambroll deploy

.PHONY: invoke
invoke:
	curl -X POST -H 'Content-type: image/png' --data-binary "@./pkg/handler/testdata/gopher.png" "http://localhost:9000/2015-03-31/functions/function/invocations"

.PHONY: execute
execute:
	curl -X POST -H 'Content-type: image/png' --data-binary "@./pkg/handler/testdata/gopher.png" https://3d8r7a230b.execute-api.ap-northeast-1.amazonaws.com/default/lambdraw

.PHONY: execute-node
execute-node:
	@node execute/node/main.mjs

.PHONY: execute-php
execute-php:
	@php execute/php/main.php
