# Practical Microservices in Go

I'm working through the book [Practical Microservices](https://pragprog.com/titles/egmicro/practical-microservices/) which uses Node as the teaching lanugage.

I'm not interested in practicing Node. I'm interested in practicing Go. So, I'm doing this in Go.

## Commands

Start database

```sh
docker compose up
```

Connect to database

```sh
psql postgres://postgres:password@localhost:5432/practical_microservices
```

Migrate database

```sh
migrate -path 'db/migrations' -database 'postgres://postgres:password@localhost:5432/practical_microservices?sslmode=disable' up
```

Seed database

```sh
psql postgres://postgres:password@localhost:5432/practical_microservices -f db/seeder.sql
```

Create a migration

```sh
migrate create -ext sql -dir db/migrations create_videos_table
```

Run app

```
APP_NAME='Video Tutorials' DB='postgres://postgres:password\localhost:5432/practical_microservices' MESSAGE_STORE_DB='postgres://postgres:password\localhost:5433/message_store?search_path=message_store' PORT=3000 ./video-tutorials
``

Run app and watch (requires modd)

```

modd

```


```
