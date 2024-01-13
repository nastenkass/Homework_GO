package homePeople

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HomePeople struct {
	Name       string
	Lastname   string
	Age        int
	Profession string
	LiveRoom   string
}

func InitializePeopleFromDatabase(connPool *pgxpool.Pool) ([]HomePeople, error) {
	query := "SELECT name, lastname, age, profession, liveroom FROM homePeople"
	rows, err := connPool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var people []HomePeople
	for rows.Next() {
		var person HomePeople
		err := rows.Scan(&person.Name, &person.Lastname, &person.Age, &person.Profession, &person.LiveRoom)
		if err != nil {
			return nil, err
		}
		people = append(people, person)
	}

	return people, nil
}

func PrintPeople(people []HomePeople) {
	fmt.Printf("Жители дома:\n")
	for _, person := range people {
		fmt.Printf("Имя: %s\nФамилия: %s\nВозраст: %d\nПрофессия: %s\nКомната проживания: %s\n\n", person.Name, person.Lastname, person.Age, person.Profession, person.LiveRoom)
	}
}
