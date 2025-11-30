run:
	go run .

dev:
	air

lint:
	golangci-lint run 

fix:
	golangci-lint run --fix