service-salt:
	go run ./generate_salt_service/cmd/app/main.go

service-user:
	go run ./user_service/cmd/app/main.go

build:
	docker compose build

run:
	docker compose up -d