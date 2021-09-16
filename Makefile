build:
  GOARCH=amd64 GOOS=linux go build -o bin/main
emulate:
	docker build -t resize-api .
	docker run -d -p 9000:8080 --name resize-api resize-api:latest /main
down:
	docker stop resize-api
	docker rm resize-api
	docker rmi resize-api
reemulate:
	@make down
	@make emulate
deploy:
	lambroll deploy
invoke:
	curl -X POST -H 'Content-type: image/png' --data-binary "@./tmp/original/gopher.png" "http://localhost:9000/2015-03-31/functions/function/invocations"
post:
	curl -X POST -H 'Content-type: image/png' --data-binary "@./tmp/original/gopher.png" https://3d8r7a230b.execute-api.ap-northeast-1.amazonaws.com/default/resize-api