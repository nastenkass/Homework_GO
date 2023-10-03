package homePeople

import "fmt"

type HomePeople struct {
	Name       string
	Lastname   string
	Age        int
	Profession string
	LiveRoom   string
}

func InitializePeople() []HomePeople {
	people1 := HomePeople{
		Name:       "Настя",
		Lastname:   "Савина",
		Age:        20,
		Profession: "Мамин шарф",
		LiveRoom:   "Спальня_№1 (основная)",
	}
	people2 := HomePeople{"Анастасия", "Цой", 21, "Создатель таблеток", "Спальня_№1 (основная)"}
	people3 := HomePeople{"Леонардо", "Дикаприо", 48, "Разливатель пива", "Спальня_№2 (гостевая)"}
	people4 := HomePeople{"Железный", "Человек", 31, "Супергерой", "Гостинная"}
	people5 := HomePeople{"Акакий", "Акакиевич", 35, "Маленький человек", "Кухня"}

	return []HomePeople{people1, people2, people3, people4, people5}
}

func PrintPeople(people []HomePeople) {
	fmt.Printf("Жители дома:\n")
	for _, person := range people {
		fmt.Printf("Имя: %s\nФамилия: %s\nВозраст: %d\nПрофессия: %s\nКомната проживания: %s\n\n", person.Name, person.Lastname, person.Age, person.Profession, person.LiveRoom)
	}
}
