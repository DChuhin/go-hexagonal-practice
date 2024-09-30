.PHONY: test build swagger-generate

test:
	go test ./...

build: test
	docker build -t iot-application-service-go .

# See https://github.com/swaggo/swag for details. Requires manual deleting LeftDelim and RightDelim in docs.go
swagger-generate:
	swag init -g cmd/main.go -o ./internal/docs --parseInternal