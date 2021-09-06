run:
	docker run -d -p 9000:8080 --name resize-api resize-api:latest /main
stop:
	docker stop resize-api
rm:
	docker stop resize-api
invoke:
	curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'