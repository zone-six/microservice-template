# DB

## Migrations

Using go migrate

```bash
migrate create -ext sql -dir db/migrations -seq <migration name>
```

## Connect to Postgres

Open terminal

```bash
docker exec -it <container-name> bash
psql -U postgres
CREATE DATABASE <db-name>;
\l #To confirm db was created;
\c <db-name> #To connect to <db-name>;
```

## TODO
- See about using go embed instead of bindata