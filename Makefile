default: format lint vet unit

run:
	go run *.go

format:
	gofmt -l -w -s .

lint:
	golangci-lint run

vet:
	go vet ./...

unit:
	go test -race ./...
