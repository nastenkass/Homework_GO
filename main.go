package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"my_home/homeInfo/buildHouse"
	"os"
)

func main() {
	urlExample := "postgres://my_home:1234@db:5432/testdb"
	_, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stderr, "ABLE to connect to database\n")
		fmt.Println("ready")
	}
	buildHouse.BuildProject()
}
