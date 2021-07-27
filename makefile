run:
	go run main.go
run-test:
	go test ./... -v
run-test-coverage:
	go test  ./...  -v -coverprofile coverage/cover.out ./
	go tool cover -html=coverage/cover.out -o coverage/cover.html
build:
	go mod tidy && go build -o bin/
docker-build-run:
	docker build -t web-service-gin:1 .;
	docker run --name=web-service-gin --rm -p 8080:8080 web-service-gin:1

docker-run:
	docker run --name=web-service-gin --rm -p 8080:8080 web-service-gin:1