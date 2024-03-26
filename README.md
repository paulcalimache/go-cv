# Go curriculum
Curriculum Vitae generator in pdf &amp; html format based from a simple yaml config file

## Usage

Linter
```sh
golangci-lint run -v
```

Tests
```sh
go test ./...
```

Run
```sh
go run main.go --help

go run main.go generate -f examples/data.yaml
```