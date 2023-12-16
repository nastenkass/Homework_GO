package buildHouse

import (
	"my_home/homeInfo/homeAppliances"
	"my_home/homeInfo/homeFurniture"
	"my_home/homeInfo/homePeople"
	"my_home/homeInfo/homeSizeAndRooms"
)

func BuildProject() {
	rooms := homeSizeAndRooms.InitializeRooms()
	furniture := homeFurniture.InitializeFurniture()
	appliance := homeAppliances.InitializeAppliances()
	people := homePeople.InitializePeople()

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
}

//проверка..
