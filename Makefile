lint:
	golangci-lint run -c ./golangci.yml ./...

test:
	go test ./... -v --cover
	cd portal && pnpm --filter "uvid-js" --filter "dash" test

test-report:
	go test ./... -v --cover -coverprofile=coverage.out
	go tool cover -html=coverage.out

build:
	cd portal && pnpm --filter "uvid-js" --filter "dash" build
	GOFLAGS=-mod=mod go build -o bin/uvid main.go

run: 
	gin --immediate run main.go