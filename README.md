# s-portal
A backend of website using Gin framework.

## To create local database
```bash
docker compose --env-file ".env.local" -f "db.local.compose.yml" up -d --build
```

## To develope our backend
- Create .env.local file and fill it. You could see the example from .env.example
- Simply start devcontainer (Don't forget to install Docker)
- Run dev command
```bash
go run main.go
```
  or using air
```bash
air
```
