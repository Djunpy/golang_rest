include app.env


migration_up:
	migrate -path db/migrations/ -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

migration_down:
	migrate -path db/migrations/ -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down

migration_fix:
	migrate -path db/migrations/ -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" force ${VERSION}

sqlc:
	sqlc compile
	sqlc generate