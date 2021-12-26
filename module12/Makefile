export VERSION=3.0.0

build:
	docker build -t jikebang/httpserver:${VERSION} .
build.push:build
	docker push jikebang/httpserver:${VERSION}

run:
	docker run -p 8080:8080 jikebang/httpserver:${VERSION}

deploy:
	kubectl apply -f k8s.yaml

del:
	kubectl delete -f k8s.yaml