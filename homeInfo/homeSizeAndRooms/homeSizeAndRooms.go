package homeSizeAndRooms

import "fmt"

type HomeSizeAndRooms struct {
	Height   float64
	Long     float64
	Width    float64
	RoomName string
}

func InitializeRooms() []HomeSizeAndRooms {
	kitchenRoom := HomeSizeAndRooms{
		Height:   2.8,
		Long:     3.2,
		Width:    4.3,
		RoomName: "Кухня",
	}
	livingRoom := HomeSizeAndRooms{2.8, 5.8, 4.2, "Гостинная"}
	bedRoom := HomeSizeAndRooms{2.8, 3.4, 5.4, "Спальня_№1 (основная)"}
	guestRoom := HomeSizeAndRooms{2.8, 3.8, 3.5, "Спальня_№2 (гостевая)"}
	bathRoom := HomeSizeAndRooms{2.8, 4.8, 2, "Санузел"}

	return []HomeSizeAndRooms{kitchenRoom, livingRoom, bedRoom, guestRoom, bathRoom}
}

func (obj HomeSizeAndRooms) RoomArea() float64 {
	return obj.Long * obj.Width
}

func (obj HomeSizeAndRooms) RoomVolume() float64 {
	return obj.RoomArea() * obj.Height
}

func (obj HomeSizeAndRooms) CountWindow() int {
	sizeWindow := 1.96
	naturalLightPercent := 40.0
	return int(((obj.RoomArea() * naturalLightPercent) / sizeWindow) / 100)
}

func (obj HomeSizeAndRooms) InfoAboutRoom() {
	fmt.Println(obj.RoomName)
	fmt.Printf("Площадь комнаты - %.2f m² \n", obj.RoomArea())
	fmt.Printf("Объем комнаты - %.2f m³ \n", obj.RoomVolume())
	fmt.Printf("Количество окон - %v \n", obj.CountWindow())
}

func PrintInfoAboutRooms(rooms []HomeSizeAndRooms, roomName string) {
	for _, room := range rooms {
		if room.RoomName == roomName {
			room.InfoAboutRoom()
			break
		}
	}
}

func calculatingAreaHouse(rooms ...HomeSizeAndRooms) float64 {
	totalArea := 0.0
	for _, room := range rooms {
		totalArea += room.RoomArea()
	}
	return totalArea
}

func AreaHouse(rooms ...HomeSizeAndRooms) {
	totalArea := calculatingAreaHouse(rooms...)
	fmt.Printf("Общая площадь дома: %.2f m²\n\n", totalArea)
}
