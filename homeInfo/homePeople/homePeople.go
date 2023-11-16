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
	return []HomePeople{
		{"Настя", "Савина", 20, "Вязальщица", "Спальня_№1 (основная)"},
		{"Анастасия", "Цой", 21, "Создатель таблеток", "Спальня_№1 (основная)"},
		{"Леонардо", "Дикаприо", 48, "Менеджер", "Спальня_№2 (гостевая)"},
		{"Железный", "Человек", 31, "Супергерой", "Гостинная"},
		{"Акакий", "Акакиевич", 35, "Маленький человек", "Кухня"},
	}
}

func PrintPeople(people []HomePeople) {
	fmt.Printf("Жители дома:\n")
	for _, person := range people {
		fmt.Printf("Имя: %s\nФамилия: %s\nВозраст: %d\nПрофессия: %s\nКомната проживания: %s\n\n", person.Name, person.Lastname, person.Age, person.Profession, person.LiveRoom)
	}
}
