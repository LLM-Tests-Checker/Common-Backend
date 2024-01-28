install-all:
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

generate-all:
	mkdir -p internal/generated/schema
	oapi-codegen -package dto -generate chi-server,types,spec api/schema.yaml > internal/generated/schema/dto.gen.go

launch-dev-env:
	sh ./dev/mongodb/start.sh
	sh ./dev/kafka/start.sh

stop-dev-env:
	sh ./dev/mongodb/stop.sh
	sh ./dev/kafka/stop.sh

test:
	 go test ./...
lint:
	golangci-lint run

launch-server:
	go run ./cmd/server/main.go
launch-consumer:
	go run ./cmd/consumer/main.go

launch-worker:
	go run ./cmd/worker/main.go
