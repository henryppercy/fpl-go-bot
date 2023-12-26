ARGS = $(arg)

run:
	@echo "Running FPL Go Bot with league IDs: $(ARGS)\n"
	@go run cmd/fpl-go-bot/main.go $(ARGS)
