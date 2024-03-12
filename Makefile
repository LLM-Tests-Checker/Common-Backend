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
	go run ./cmd/consumer/llm_result/main.go

launch-worker:
	go run ./cmd/worker/launch_llm_check/main.go
	go run ./cmd/worker/drop_old_in_progress_llm_check/main.go

build-all-docker-images:
	docker build -t common-backend-server:latest -f ./docker/server/Dockerfile .

