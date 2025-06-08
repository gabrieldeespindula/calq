build:
	docker compose run --rm dev go build -o calq ./cmd/calq

run:
	docker compose run --rm dev go run ./cmd/calq

release-local:
	docker compose run --rm releaser release --snapshot --clean

test:
	docker compose run --rm dev go test -coverprofile=coverage.out ./...

coverage: test
	docker compose run --rm dev go tool cover -html=coverage.out -o coverage.html
	open coverage.html || xdg-open coverage.html || start coverage.html
