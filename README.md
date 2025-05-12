# Url shortener with Go

Simple web server for a url shortener

## Requirements

- docker (with compose)
- Go (1.24)
- Just

## How to run

1. Start Valkey through docker compose
```sh
docker compose up -d
```

2. Start the server with Just
```sh
just run
```

## Usage

To create send the url to get a short version of it:

```sh
curl http://localhost:8000 \
-d '{"url": "https://google.com"}'
```

It should return a json with "Id", this Id can be used in a get request

```sh
curl -v http://localhost:8000/:id
```
