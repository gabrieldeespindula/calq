FROM golang:1.22

WORKDIR /app

COPY ./app ./

# Go exige um módulo, mesmo para coisas simples
RUN go mod init playground && \
    go mod tidy

CMD ["go", "run", "main.go"]
