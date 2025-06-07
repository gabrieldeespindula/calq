build:
	docker compose run --rm dev go build -o calq ./cmd/calq

run:
	docker compose run --rm dev go run ./cmd/calq

release-local:
	docker compose run --rm releaser release --snapshot --clean