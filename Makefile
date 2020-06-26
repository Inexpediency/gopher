.PHONY: build
build:
	go build ./main.go

.PHONY: run
run:
	go run ./main.go

.PHONY: test
test:
	go test -timeout 30s ./...

.PHONY: israce
race:
	go run ./main.go -race
