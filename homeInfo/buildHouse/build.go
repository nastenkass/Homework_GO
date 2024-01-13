package buildHouse

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"my_home/homeInfo/homeAppliances"
	"my_home/homeInfo/homeFurniture"
	"my_home/homeInfo/homePeople"
	"my_home/homeInfo/homeSizeAndRooms"
	"os"
)

func BuildProject(connPool *pgxpool.Pool) error {
	rooms := homeSizeAndRooms.InitializeRooms()
	furniture := homeFurniture.InitializeFurniture()
	appliance := homeAppliances.InitializeAppliances()

	people, err := homePeople.InitializePeopleFromDatabase(connPool)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка получения данных о людях из базы данных: %v\n", err)
		os.Exit(1)
	}

	roomNames := []string{
		"Спальня_№1 (основная)",
		"Кухня",
		"Спальня_№2 (гостевая)",
		"Гостинная",
		"Санузел",
	}

	for _, roomName := range roomNames {
		homeSizeAndRooms.PrintInfoAboutRooms(rooms, roomName)
		homeFurniture.AreaOccupiedFurniture(roomName, furniture...)
		homeAppliances.AreaOccupiedAppliances(roomName, appliance...)
	}

	homeSizeAndRooms.AreaHouse(rooms...)
	homePeople.PrintPeople(people)

	return nil
}
