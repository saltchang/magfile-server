postgres:
	chmod u+x ./scripts/postgres.sh
	./scripts/postgres.sh

run-postgres:
	docker run --name postgres13.2 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13.2-alpine

stop-postgres:
	docker stop postgres13.2

createdb:
	docker exec -it postgres13.2 createdb --username=root --owner=root magfile_server

dropdb:
	docker exec -it postgres13.2 dropdb --username=root magfile_server

accessdb:
	docker exec -it postgres13.2 psql -U root magfile_server

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/magfile_server?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/magfile_server?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

reset-db:
	make stop-postgres;make postgres;sleep 5;make createdb;make migrateup;sleep .5;echo "\nRest DB done.";

.PHONY: postgres run-postgres stop-postgres createdb dropdb migrateup migratedown sqlc test reset-db
