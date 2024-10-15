DB_URL=postgresql://postgres:12345678@localhost:5433/postgres?sslmode=disable
mup:
	goose -dir app/db/migration postgres "$(DB_URL)" up
mdown:
	goose -dir app/db/migration postgres "$(DB_URL)" down
	 
sqlc:
	sqlc generate
	
server:
	gin -p 8081 -i run main.go

postgres:
	docker run -d  --name postgres1  -p 5433:5432 -e POSTGRES_PASSWORD=12345678  -e PGDATA=/var/lib/postgresql/data/pgdata  -v postgres_volume:/var/lib/postgresql/data  postgres:13-alpine

postgres1:
	docker run -d  --name postgres1  -p 5433:5432 -e POSTGRES_PASSWORD=12345678  -e PGDATA=/var/lib/postgresql/data/pgdata  -v postgres_volume:/var/lib/postgresql/data  postgres:15-alpine

.PHONY: mup mdown postgres 