postgres:
    docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
    docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
    docker exec -it postgres12 dropdb simple_bank

migrateup:
    migrate -path db/migrations -database "postgres://default:lZsIkJCjEV97@ep-round-voice-37130748-pooler.ap-southeast-1.aws.neon.tech:5432/verceldb?sslmode=require" -verbose up

migratedown:
    migrate -path db/migrations -database "postgres://default:lZsIkJCjEV97@ep-round-voice-37130748-pooler.ap-southeast-1.aws.neon.tech:5432/verceldb?sslmode=require" -verbose down

migratedown1:
    migrate -path db/migrations -database "postgres://default:lZsIkJCjEV97@ep-round-voice-37130748-pooler.ap-southeast-1.aws.neon.tech:5432/verceldb?sslmode=require" -verbose down 1

sqlc:
    sqlc generate

test:
    go test -v -cover ./...

server:
    go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown migratedown1 sqlc test server