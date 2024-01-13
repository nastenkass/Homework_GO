package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"my_home/homeInfo/buildHouse"
	"os"
)

func main() {
	urlExample := "postgres://my_home:1234@my_home-db:5432/testdb"
	connPool, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Не удалось установить соединение: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Соединение установлено")

	buildHouse.BuildProject(connPool)

	fmt.Println("Готово")
}
