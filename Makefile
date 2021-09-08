build:
	docker build -t resize-api .
rmi:
	docker rmi resize-api
run:
	docker run -d -p 9000:8080 --name resize-api resize-api:latest /main
stop:
	docker stop resize-api
rm:
	docker rm resize-api
start:
	@make build
	@make run
restart:
	@make stop
	@make rm
	@make rmi
	@make build
	@make run
down:
	@make stop
	@make rm
invoke:
	curl -H "Content-Type: image/png" --data-binary "@./tmp/original/gopher.png" -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations"