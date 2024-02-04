# Base command
default:
    @just --list

runall CMD *ARGS:
    -cd core && {{CMD}} {{ARGS}}
    -cd service && {{CMD}} {{ARGS}}

test:
    go test -v github.com/jgfranco17/gitglow/...

build:
    @echo "Building CLI app..."
    go mod download all
    CGO_ENABLED=0 GOOS=linux go build -o ./app service/cmd/main.go
    @echo "Build successful!"
