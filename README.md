# calq

[![Release](https://img.shields.io/github/v/release/gabrieldeespindula/calq)](https://github.com/gabrieldeespindula/calq/releases)

A simple terminal calculator written in Go.  
It evaluates basic mathematical expressions (addition, subtraction, multiplication, division) with operator precedence and supports using the last result in new expressions.

---

## Features

- Parses and evaluates simple math expressions
- Interactive terminal usage
- Operator precedence handling (`*` and `/` before `+` and `-`)
- Supports reusing the last result in new calculations

---

## Usage

### Prebuilt Executable (Release)

Download the latest release from the [Releases](https://github.com/gabrieldeespindula/calq/releases) page, extract it, and run:

```bash
./calq
```

Type your expressions and see the results.
Press `Ctrl+C` to exit.

---

### Development

To develop, build, or run the project locally using Docker and Makefile (no Go installation required):

1. Clone the repository:

```bash
git clone https://github.com/gabrieldeespindula/calq.git
cd calq
```

2. To build the binary:

```bash
make build
```

3. To run the project without building:

```bash
make run
```

These commands run Go inside a Docker container to provide a clean and reproducible environment.

---

## Project Structure

* `cmd/calq/main.go`: application entrypoint
* `internal/calculator/`: calculator logic (tokenization, evaluation, operations)
* `Dockerfile`: Docker image for development environment
* `docker-compose.yml`: Docker service config for development
* `Makefile`: shortcuts for building and running using Docker

---

## Technologies

* Go 1.22
* Docker (for isolated development environment)
* Makefile (command automation)

---

## Usage Examples

```plaintext
Welcome to calq! Type your expressions (Ctrl+C to exit):
> 3 + 4 * 2
11
> + 5
16
> 100 / 4 - 7
18
```

## Local Release (Snapshot)

Você pode gerar uma release local (snapshot) para testar os artefatos do GoReleaser sem publicar no GitHub. Isso ajuda a validar o build e a estrutura dos binários antes de uma release real.

### Requisitos

* Docker
* `docker-compose`

### Rodando manualmente

```bash
make release-local
```

Isso usará o container `goreleaser/goreleaser` definido no `docker-compose.yml` para gerar os binários localmente dentro da pasta `dist/`.

---

## Contributing

Contributions are welcome!
Please open issues for bugs or feature requests, and submit pull requests for review.

---

## License

MIT License © Gabriel de Espindula

---

## Contact

Gabriel de Espindula – [@gabrieldeespindula](https://github.com/gabrieldeespindula)

## About

This project was created as a personal learning exercise to deepen my understanding of Go, building from scratch a simple calculator with expression parsing and evaluation. While libraries exist that solve this problem more comprehensively, this project focuses on educational purposes.
