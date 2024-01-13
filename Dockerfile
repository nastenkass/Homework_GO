FROM golang:1.21

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./

# проверка зависимостей
RUN go mod download
RUN go mod verify

COPY . .
RUN go build -o ./my_home/main ./main.go
RUN chmod +x ./

CMD ["./main"]