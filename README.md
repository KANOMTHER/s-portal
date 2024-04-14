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

## To access API documentation
The API documentation is generated using [Swaggo](https://github.com/swaggo/swag?tab=readme-ov-file). Follow these steps to view the API documentation:
1. run your application
2. Open your web browser and navigate to http://localhost:3000/swagger/index.html#/.
3. You will see Swagger 2.0 API documents similar to the image below:
![image](https://github.com/KANOMTHER/s-portal/assets/89908219/0dffadce-ab9f-47ba-b23e-61a43783b2e5)

## To generate API documentation
To generate API documentation using Swaggo, use the following command:
```bash
swag --parseDependency --parseInternal
```
For additional options and configurations, refer to the [Swaggo documentation](https://github.com/swaggo/swag?tab=readme-ov-file#swag-cli).

For an example of using Swaggo with Gin, check out the [Swaggo example with Gin](https://github.com/swaggo/swag/tree/master/example/celler).

## To format swag
The Swag Comments can be automatically formatted, similar to go fmt.
Usuage:
```bash
swag fmt
```
This command will format Swag comments in your codebase according to predefined standards.
