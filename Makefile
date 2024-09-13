run:
	@go run cmd/main.go -path .

test:
	@go test -timeout 30s -coverprofile=./coverage ./... -v