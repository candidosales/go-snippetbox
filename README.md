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
* Standard Go Project Layout: https://github.com/golang-standards/project-layout

### Considerations

* All configuration variable have to be a environment variable
* Every request is handled in it’s own goroutine
* Function name start with capital letter

### Development

Generate certificates

```bash
go run $GOROOT/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

```bash
go rub cmd/snippetbox/*
go run cmd/snippetbox/!(*_test).go
```

Note: If you get an event not found error when running this command, you probably need to enable extended globbing in your bash terminal first. You can do this by running:

```bash
shopt -s extglob
```

## Tests

Test to app

```bash
go test -v ./cmd/snippetbox
```

Test database

```bash
go test -v ./pkg/models/mysql
```

### Integration Tests

You need to create a test database

```sql

CREATE DATABASE test_snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci ;

CREATE USER 'test_web'@'localhost';
GRANT CREATE, DROP, ALTER, INDEX, SELECT, INSERT, UPDATE, DELETE ON test_snippetbox.* TO 'test_web'@'localhost';ALTER USER 'test_web'@'localhost' IDENTIFIED BY 'pass';

## Reference

* https://github.com/cullenjett/snippetbox


## Libraries

* Test: https://github.com/stretchr/testify
* Middleware chaining: https://github.com/justinas/alice