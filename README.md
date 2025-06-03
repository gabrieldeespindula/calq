# 🧮 Go Calculator (with Docker)

Esta é uma calculadora de linha de comando escrita em Go, que permite realizar as quatro operações básicas: **adição, subtração, multiplicação e divisão**, com suporte a **números decimais (float64)**.

O projeto está totalmente empacotado com Docker, permitindo que você execute o código sem precisar instalar o Go localmente.

---

## 🚀 Como rodar com Docker Compose

### Pré-requisitos

- Docker instalado
- Docker Compose (v2 ou superior)

---

### 📁 Estrutura do Projeto

```

.
├── app
│   └── main.go          # Código da calculadora
├── Dockerfile           # Instruções para construir o container
└── docker-compose.yml   # Orquestração do container

````

---

### 📦 Build do projeto (opcional)

```bash
docker compose build
````

> Você só precisa fazer isso se mudar o `Dockerfile`.

---

### ▶️ Executando a calculadora

Use o seguinte comando para rodar a calculadora no terminal:

```bash
docker compose run --rm -it go
```

* `--rm` remove o container assim que o programa termina.
* `-it` permite entrada interativa via terminal (necessário para o `fmt.Scanln` funcionar corretamente).
* `go` é o nome do serviço no `docker-compose.yml`.

---

## 🧠 Funcionalidades

* Lê dois números `float64` do usuário
* Permite escolher a operação: `+`, `-`, `*`, `/`
* Trata divisão por zero com erro amigável
* Loop para múltiplos cálculos, com opção de sair (`y/n`)

---

## 📌 Exemplo de uso

```text
Type the first number: 10.5
Choose an operation: addition(+), subtraction(-), multiplication(*), division(/)
+
Type the second number: 2

The result of 10.50 + 2.00 is: 12.50
Do you want to calculate again? (y/n): n
Thank you for using the calculator!
```

---

## 🛠️ Tecnologias

* [Go 1.22](https://go.dev/doc/go1.22)
* Docker e Docker Compose

---

## 📄 Licença

Este projeto é livre para estudo e modificação.

