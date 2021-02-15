postgres:
	docker run --name postgres13.2 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13.2-alpine

createdb:
	docker exec -it postgres13.2 createdb --username=root --owner=root magfile_server

dropdb:
	docker exec -it postgres13.2 dropdb magfile_server

accessdb:
	docker exec -it postgres13.2 psql -U root magfile_server

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/magfile_server?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/magfile_server?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
