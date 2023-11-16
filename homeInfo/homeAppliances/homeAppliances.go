package homeAppliances

import "fmt"

type HomeAppliances struct {
	Name     string
	Height   float64
	Width    float64
	Long     float64
	RoomName string
}

func InitializeAppliances() []HomeAppliances {
	return []HomeAppliances{
		{"Кухонная плита", 0.85, 0.6, 0.6, "Кухня"},
		{"Холодильник", 1.817, 0.629, 0.6, "Кухня"},
		{"Посудомойка", 0.64, 0.55, 0.599, "Кухня"},
		{"Стиральная машина", 0.852, 0.59, 0.6, "Санузел"},
		{"Кондиционер", 0.33, 0.232, 1.132, "Спальня_№1 (основная)"},
		{"Обогреватель", 0.137, 0.94, 1.5, "Спальня_№2 (гостевая)"},
		{"Электронный камин", 1.045, 0.39, 1.18, "Гостинная"},
	}
}

func CalculatingAreaAppliancesForRoom(roomName string, appliances ...HomeAppliances) float64 {
	totalArea := 0.0
	for _, appliancesObj := range appliances {
		if appliancesObj.RoomName == roomName {
			totalArea += appliancesObj.AreaAppliances()
		}
	}
	return totalArea
}

func AreaOccupiedAppliances(roomName string, appliance ...HomeAppliances) {
	totalAreaAppliances := CalculatingAreaAppliancesForRoom(roomName, appliance...)
	fmt.Printf("Занимаемая площадь бытовой техники в комнате «%s»: %.2f m²\n\n", roomName, totalAreaAppliances)
}

func (obj HomeAppliances) AreaAppliances() float64 {
	return obj.Width * obj.Long
}
