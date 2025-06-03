# Go Playground com Docker Compose

Este projeto é um ambiente simples para aprender Go dentro de um container Docker, com hot reload básico para testar código de forma rápida.

## O que ele faz?

* Um programa Go que lê dois números e uma operação (+, -, \*, /) do usuário.
* Realiza o cálculo solicitado e exibe o resultado.
* Trata divisão por zero e operações inválidas com mensagens apropriadas.

## Estrutura

* `Dockerfile`: cria a imagem Go, define diretório de trabalho e comando padrão para rodar o app.
* `docker-compose.yml`: configura serviço `go`, monta volume para sincronizar código local com container.
* `app/main.go`: código-fonte Go do programa interativo.

## Como rodar

1. Clone o projeto e entre na pasta:

   ```bash
   git clone <repo-url>
   cd <pasta-do-projeto>
   ```

2. Execute o container interativo para rodar o app:

   ```bash
   docker compose run --rm -it go
   ```

3. O programa pedirá que você digite dois números e escolha uma operação. Exemplo:

   ```
   Digite o primeiro número: 10
   Escolha uma operação: soma(+), subtração(-), multiplicação(*), divisão(/)
   +
   Digite o segundo número: 5
   Resultado: 15
   ```

## Desenvolvimento

* Você pode editar o código localmente em `app/main.go`.
* Ao rodar o container com `docker compose run --rm -it go`, seu código será executado dentro do container.

## Comandos úteis

* Para abrir um shell no container e rodar comandos Go manualmente:

  ```bash
  docker compose run --rm -it go bash
  ```

* Para reconstruir a imagem (se mudar o Dockerfile):

  ```bash
  docker compose build
  ```
