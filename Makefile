CURRENT_DIR=$(shell pwd)
DB_URL="postgres://postgres:mubina2007@db:5434/userdb?sslmode=disable"

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	./scripts/gen-proto.sh 

migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up 

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migrate-force:
	migrate -path migrations -database "$(DB_URL)" -verbose force

migrate-file:
	migrate create -ext sql -dir migrations/ -seq create_account_verification_table

g:
	go run cmd/main.go