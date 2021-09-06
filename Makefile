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
restart:
	@make stop
	@make rm
	@make rmi
	@make build
	@make run
invoke:
	curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'