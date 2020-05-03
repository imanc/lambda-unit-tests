PROJECT_NAME := lambda-unit-testing
CUR_DIR := $$(pwd)

build-in-docker: clean
	docker run -ti --rm \
	-w="/home/dev" \
	-v $(CUR_DIR):/home/dev \
	-e GOOS=linux \
	-e GOARCH=amd64 \
	golang:1.14.2 \
	go build -o bin/$(PROJECT_NAME) ./cmd

build: clean
	GOOS=linux GOARCH=amd64 go build -o bin/$(PROJECT_NAME) ./cmd


test-in-docker: clean
	docker run -ti --rm \
	-w="/home/dev" \
	-v $(CUR_DIR):/home/dev \
	-e GOOS=linux \
	-e GOARCH=amd64 \
	golang:1.14.2 \
	go test ./...

test: 
	go test ./...

clean:
	rm -rf bin