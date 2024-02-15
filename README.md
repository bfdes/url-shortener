# url-shortener

![GitHub Actions](https://github.com/bfdes/url-shortener/workflows/Test/badge.svg)
[![Codecov](https://codecov.io/gh/bfdes/url-shortener/branch/master/graph/badge.svg)](https://codecov.io/gh/bfdes/url-shortener)

URL shortener and redirect service [designed for](https://www.notion.so/URL-shortening-8272c692648143698859d9f3524a8b5e#a2becb53582444cfb3e9cce1dd8978ba) low-latency, read-heavy use.

## Usage

### Requirements

- [Go](https://golang.org/) 1.22.*
- [Docker Engine](https://docs.docker.com/engine/) 20.*
- [Docker Compose](https://docs.docker.com/compose/) 2.*

Run the following command within the repository root to start container dependencies in the background:

```shell
docker compose up --detach cache database
```

Then, when the databases are ready to accept connections, start the server with `go run .`.

### Shorten a URL

```shell
curl http://localhost:8080/api/links \
  --request POST \
  --data '{"url": "http://example.com"}'
# {"url": "http://example.com", "slug": "<SLUG>" }
```

### Redirect a URL

```shell
curl http://localhost:8080/<SLUG>
```

### Testing

Run unit and integration tests with `go test` after starting container dependencies.

[GitHub Actions](https://github.com/bfdes/url-shortener/actions) will run tests for every code push.

## Deployment

This URL shortener is unsuited for production use; it does not support logging or metric collection.
