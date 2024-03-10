create_migration:
	migrate create -ext sql -dir ./db/migration -seq $(migration)

migrateup:
	migrate -path ./db/migration -database "$(db_url)" -verbose up

migratedown:
	migrate -path ./db/migration -database "$(db_url)" -verbose down


make migrateup db_url="mysql://root:root@tcp(127.0.0.1:52000)/user_task?x-tls-insecure-skip-verify=true"