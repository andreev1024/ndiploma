My [niece](https://github.com/Troublemaker06) diploma. My part is the backend (GO coding + DB + Docker).

# How to run locally

-   run in the project directory

```sh
docker compose up # the web server will start (shell #1)
docker compose --profile tools run migrate up # migration will execute (shell #2)
```

-   open `localhost:8081` in the browser.

# API

```sh
curl -F name=Andrey -F phone=+375291112233 -F role=teacher -F available-time=12:00 -F consult-date=15.04.2023 http://localhost:8081/consult-requests
```

# Migrations

```sh
docker compose --profile tools run migrate up
docker compose --profile tools run migrate down 1
docker compose --profile tools run create-migration
```

# Links

https://firehydrant.com/blog/develop-a-go-app-with-docker-compose/
