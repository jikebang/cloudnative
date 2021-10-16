build:
	docker build -t jikebang/httpserver:v1.0.0 .
build.push:build
	docker push jikebang/httpserver:v1.0.0

run:
	docker run -p 8080:8080 jikebang/httpserver:v1.0.0