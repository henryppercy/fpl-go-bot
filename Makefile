run:
	@echo "Running FPL Go Bot\n"
	@go run cmd/fpl-go-bot/main.go

fmt:
	gofmt -w -s .