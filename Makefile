
psql:
	psql favmov -U adil

create_db:
	createdb --username=adil --owner=adil favmov

drop_db:
	dropdb favmov -U adil

run:
	go run main.go

migrate_up:
	migrate -path db/migration -database "postgresql://adil:123456@localhost:5432/favmov?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://adil:123456@localhost:5432/favmov?sslmode=disable" -verbose down

migrate_force:
	migrate -path db/migration -database "postgresql://adil:123456@localhost:5432/favmov?sslmode=disable" force 1
	