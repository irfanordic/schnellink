# schnellink

A URL shortener written in Go with [Gin](https://github.com/gin-gonic/gin), built as a learning project for backend and system design.

## Running locally

Requires Go 1.26+.

```bash
go run .
```

The server listens on `:8081`.

## API

### `GET /ping`

Health check.

```bash
curl localhost:8081/ping
# {"message":"pong"}
```

### `POST /shorten`

Generates a short code for a URL. `url` is required.

```bash
curl -X POST localhost:8081/shorten -d '{"url":"https://example.com"}'
# {"original_url":"https://example.com","short_url":"xaypYr"}
```

An empty or invalid payload returns `400`.

## Before committing

```bash
gofmt -w .
go vet ./...
go build ./...
```

## Roadmap

- [x] HTTP server with health check
- [x] Request validation on `/shorten`
- [ ] Cryptographically secure short codes (`crypto/rand`)
- [ ] `LinkStore` interface with an in-memory implementation
- [ ] Collision handling: insert, and retry on duplicate
- [ ] `GET /:code` redirect
- [ ] Tests (`go test -race ./...`)
- [ ] Persistent storage (Postgres) with a unique constraint on `short_code`
- [ ] Config from environment variables, graceful shutdown
- [ ] Docker and Docker Compose
- [ ] Redis cache, click analytics, rate limiting
- [ ] Kubernetes deployment
