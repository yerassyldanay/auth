postgres:
	docker pull postgres:11-alpine && docker run --name auth_postgres -p 0.0.0.0:8101:5432 -e POSTGRES_PASSWORD=simple -e POSTGRES_USER=simple -e POSTGRES_USER=simple -d postgres:11-alpine

postgres_delete:
	docker stop auth_postgres && docker rm auth_postgres

postgres_logs:
	docker logs auth_postgres -f --tail=30

migrate_up:
	migrate -path ./database/postgres -database postgres://simple:simple@localhost:8101/simple?sslmode=disable -verbose up

migrate_down:
	migrate -path ./database/postgres -database postgres://simple:simple@localhost:8101/simple?sslmode=disable -verbose down

generate:
	sqlc generate

test:
	go test -v ./model/sqlc/

.PHONY: postgres postgres_delete postgres_logs postgres_up postgres_down generate test
