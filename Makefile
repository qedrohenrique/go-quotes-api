.PHONY: wire
wire:
	go run github.com/google/wire/cmd/wire ./di/

.PHONY: dev
dev:
	go run main.go

.PHONY: build
build: wire
	go build -o bin/quotes-api main.go

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -rf bin/
