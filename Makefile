.PHONY: build
build:
	go build ./main.go

.PHONY: run
run:
	go run ./main.go

# .PHONY: test
# test:
# 	go test -v -race -timeout 30s ./internal/app/...

.PHONY: israce
race:
	go run ./main.go -race
