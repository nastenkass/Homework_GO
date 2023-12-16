FROM golang:1.21

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
COPY go.sum ./

# проверка зависимостей
RUN go mod download
RUN go mod verify

COPY . .
RUN go build -o ./my_home/main ./main.go
RUN chmod +x ./  # предоставление прав

CMD ["./main"]