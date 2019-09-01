# Project Go

## Libraries

* Server HTTP: https://golang.org/pkg/net/http/
* Database: https://github.com/jmoiron/sqlx
* Database migrations: https://github.com/lopezator/migrator
* Middleware: https://github.com/justinas/alice

## Book

* Let's Go! Learn to Build Professional Web Applications With Golang: https://lets-go.alexedwards.net/

## Best practices

* Use modules: https://blog.golang.org/migrating-to-go-modules

## Links

* How do you structure your Go apps? : https://www.youtube.com/watch?v=VQym87o91f8
* Concurrency patterns in Go: https://www.youtube.com/watch?v=rDRa23k70CU
* Achieving concurrency in Go: https://medium.com/rungo/achieving-concurrency-in-go-3f84cbf870ca
* Learning Go’s Concurrency Through Illustrations: https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3
* Awesome Go: https://github.com/avelino/awesome-go
* Projects in Go: https://github.com/golang/go/wiki/Projects

### Considerations

* Every request is handled in it’s own goroutine.

### Development

```bash
go rub cmd/web/*
go run cmd/web/!(*_test).go
```

### Tests

```bash
go test -v ./cmd/web
```
