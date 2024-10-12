.PHONY: build
build:
	go build -o ./bin/api-demo ./cmd/api-demo
	docker build -t api-demo .

.PHONY: run
run:
	docker run -d -p 8080:8080 api-demo:latest

.PHONY: integration-test
integration-test:
	go test -v ./test/integration

.PHONY: test
test:
	go test -v ./test/unit
