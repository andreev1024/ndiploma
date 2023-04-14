# API

```sh
curl -F "name=Andrey" -F phone=+375291112233 -F role=teacher -F available-time=12:00 -F consult-date=15.04.2023 http://localhost:8081/consult-requests
```

# Migrations

```sh
docker compose --profile tools run migrate up
docker compose --profile tools run migrate down 1
docker compose --profile tools run create-migration
```

# Links

https://firehydrant.com/blog/develop-a-go-app-with-docker-compose/
