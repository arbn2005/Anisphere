include .env

.PHONY: run build migrate-up migrate-down docker-up docker-down clean

run:
	go run .

build:
	go build -o bin/anisphere .

docker-up:
	sudo docker compose up -d

docker-down:
	sudo docker compose down

migrate-up:
	migrate \
		-path migrations \
		-database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" \
		up

migrate-down:
	migrate \
		-path migrations \
		-database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" \
		down 1

clean:
	rm -rf bin


