build:
    docker compose run dev go build -o calq ./cmd/calq

run:
	docker compose run dev go run ./cmd/calq