run:
	@go run cmd/main.go

test:
	@go test -timeout 30s -coverprofile=./coverage github.com/eric-batista/neowire-api-task-manager-poc/domain/models -v