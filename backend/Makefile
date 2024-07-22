format:
	go fmt ./...

linter: 
	golangci-lint run

run:
	swag init --parseDependency --parseInternal -g cmd/main.go
	go run cmd/main.go

docker:
	docker compose up

test:
	go test -v ./...

swagger:
	swag init --parseDependency --parseInternal -g cmd/main.go
