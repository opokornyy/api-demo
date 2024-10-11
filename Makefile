.PHONY: build
build:
	go build -o ./bin/api-demo ./cmd/api-demo
	docker build -t api-demo .

.PHONY: run
run:
	docker run -d -p 8080:8080 api-demo:latest

# TODO: add test target
# TODO: add migrate target