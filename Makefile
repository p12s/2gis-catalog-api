.PHONY:
.SILENT:
.DEFAULT_GOAL := run

# golang development
build:
	go mod download && go build cmd/main.go
run:
	go run cmd/main.go
test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage
create-migration:
	migrate create -ext sql -dir ./schema -seq init
test.coverage:
	go tool cover -func=cover.out
swag:
	swag init -g cmd/main.go
lint:
	golangci-lint run

# docker-compose

init: dc-clear dc-pull dc-build dc-up # команда для первоначального запуска в системе
dc-down:
	docker-compose --env-file ./.env down # будет останавливать все контейнеры с данным именем, даже если что-то по-ошибке удалено из docker-compose
dc-clear:
	docker-compose --env-file ./.env down -v --remove-orphans # полная очистка с удалением volumes
dc-pull:
	docker-compose --env-file ./.env pull
dc-build:
	docker-compose --env-file ./.env build
dc-up:
	docker-compose --env-file ./.env up -d
