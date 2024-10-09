CONTAINER_NAME := jsonDiff

up:
	docker compose up -d

clean:
	docker compose down

test:
	docker exec $(CONTAINER_NAME) sh -c 'go test ./...'

test-coverage:
	docker exec $(CONTAINER_NAME) sh -c ' \
		go test -coverprofile=coverage.out -coverpkg=./... ./... && \
		go tool cover -func=coverage.out && \
		go tool cover -html=coverage.out -o coverage.html'
	rm coverage.out

go_build:
	docker exec $(CONTAINER_NAME) sh -c ' \
		GOARCH=amd64 GOOS=linux go build -buildvcs=false -o bin/$(CONTAINER_NAME)-linux && \
		GOARCH=amd64 GOOS=darwin go build -buildvcs=false -o bin/$(CONTAINER_NAME)-darwin && \
		GOARCH=amd64 GOOS=windows go build -buildvcs=false -o bin/$(CONTAINER_NAME)-windows'

ecr_build:
	GOARCH=amd64 GOOS=linux go build -buildvcs=false -o bin/$(CONTAINER_NAME)-linux

build: up go_build clean
tests: up test clean
tests-coverage: up test-coverage clean
