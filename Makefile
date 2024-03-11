migrateup:
	migrate -path schema -database "postgresql://postgres:hell@localhost:5432/orders?sslmode=disable" -verbose up

migratedown:
	migrate -path schema -database "postgresql://postgres:hell@localhost:5432/orders?sslmode=disable" -verbose down

.PHONY: migrateup migratedown