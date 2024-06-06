CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd

DB_URL=postgres://postgres:abdulaziz1221@localhost:5432/n9?sslmode=disable

run:
	go run cmd/main.go

migrate_up:
	migrate -path migrations -database ${DB_URL} -verbose up

migrate_down:
	migrate -path migrations -database ${DB_URL} -verbose down

migrate_force:
	migrate -path migrations -database ${DB_URL} -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq users

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs force 1
