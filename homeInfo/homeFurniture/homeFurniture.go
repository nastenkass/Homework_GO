package homeFurniture

import "fmt"

type HomeFurniture struct {
	TypeFurniture bool
	Name          string
	Height        float64
	Width         float64
	Long          float64
	Color         string
	RoomName      string
}

func InitializeFurniture() []HomeFurniture {
	return []HomeFurniture{
		{true, "Шкаф для одежды", 2.1, 0.6, 1.8, "Серый", "Спальня_№1 (основная)"},
		{true, "Стеллаж для книг", 1.8, 0.4, 1.2, "Белый", "Спальня_№1 (основная)"},
		{false, "Рабочий стол", 0.8, 0.8, 1.2, "Черный", "Спальня_№1 (основная)"},
		{false, "Игровое кресло", 1.4, 0.6, 0.6, "Черно-белый", "Спальня_№1 (основная)"},
		{false, "Диван", 0.82, 0.93, 1.74, "Серый", "Спальня_№1 (основная)"},
	}
}

func CalculatingAreaFurnitureForRoom(roomName string, furniture ...HomeFurniture) float64 {
	totalArea := 0.0
	for _, furnitureObj := range furniture {
		if furnitureObj.RoomName == roomName {
			totalArea += furnitureObj.AreaFurniture()
		}
	}
	return totalArea
}

func AreaOccupiedFurniture(roomName string, furniture ...HomeFurniture) {
	totalAreaFurniture := CalculatingAreaFurnitureForRoom(roomName, furniture...)
	fmt.Printf("Занимаемая площадь мебели в комнате «%s»: %.2f m²\n", roomName, totalAreaFurniture)
}

func (obj HomeFurniture) AreaFurniture() float64 {
	return obj.Width * obj.Long
}
