postgres:
	chmod u+x ./scripts/postgres.sh
	./scripts/postgres.sh

run-postgres:
	docker run --name postgres_magfile -p 5437:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13.2-alpine

stop-postgres:
	docker stop postgres_magfile

createdb:
	docker exec -it postgres_magfile createdb --username=root --owner=root magfile_server

dropdb:
	docker exec -it postgres_magfile dropdb --username=root magfile_server

accessdb:
	docker exec -it postgres_magfile psql -U root magfile_server

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5437/magfile_server?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5437/magfile_server?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

reset-db:
	make stop-postgres;make postgres;sleep 5;make createdb;make migrateup;sleep .5;echo "\nRest DB done.";

.PHONY: postgres run-postgres stop-postgres createdb dropdb accessdb migrateup migratedown sqlc test reset-db
