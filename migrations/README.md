# Migrations

Generated using [migrate](https://github.com/golang-migrate/migrate#migrate).

## Create new migrations

To create new migration files just run this comand on the project root:

```bash
migrate create -ext sql -dir migrations -seq [migration_name]
```
