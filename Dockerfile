FROM golang:1.23
WORKDIR /app
COPY . .
RUN [ ! -f go.mod ] && go mod init go-kubernetes || true
RUN go mod tidy || true
RUN go build -o main .
CMD ["./main"]