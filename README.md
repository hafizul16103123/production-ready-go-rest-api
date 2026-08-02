# production-ready-go-rest-api
# Migration
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate create -ext sql -dir migrations -seq create_users
## Run migration cmd
migrate -path migrations -database "postgres://postgres:postgresql@localhost:5432/postgres?sslmode=disable" up

migrate -path migrations -database "postgres://postgres:postgresql@localhost:5432/postgres?sslmode=disable" down
## Roll Back Undo the last migration:
migrate -path migrations -database "..." down 1

## Best Practices
✅ One migration = one logical change.
✅ Never edit an already-applied migration.
✅ Keep Up and Down migrations in sync.
✅ Commit migrations to Git.\
## Common Beginner Mistakes
❌ Editing old migration files after they have been applied.
❌ Forgetting the Down migration.
❌ Combining unrelated schema changes into one migration.
❌ Running SQL manually in production.