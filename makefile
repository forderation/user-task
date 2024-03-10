create_migration:
	migrate create -ext sql -dir ./db/migration -seq $(migration)

migrateup:
	migrate -path ./db/migration -database "$(db_url)" -verbose up

migratedown:
	migrate -path ./db/migration -database "$(db_url)" -verbose down
