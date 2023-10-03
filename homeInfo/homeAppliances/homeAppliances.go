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
	stove := HomeAppliances{
		Name:     "Кухонная плита",
		Height:   0.85,
		Width:    0.6,
		Long:     0.6,
		RoomName: "Кухня",
	}
	fridge := HomeAppliances{"Холодильник", 1.817, 0.629, 0.6, "Кухня"}
	dishwasher := HomeAppliances{"Посудомойка", 0.64, 0.55, 0.599, "Кухня"}
	washingMachine := HomeAppliances{"Стиральная машина", 0.852, 0.59, 0.6, "Санузел"}
	conditioner := HomeAppliances{"Кондиционер", 0.33, 0.232, 1.132, "Спальня_№1 (основная)"}
	heater := HomeAppliances{"Обогреватель", 0.137, 0.94, 1.5, "Спальня_№2 (гостевая)"}
	electronicFireplace := HomeAppliances{"Электронный камин", 1.045, 0.39, 1.18, "Гостинная"}
	return []HomeAppliances{stove, fridge, dishwasher, washingMachine, conditioner, heater, electronicFireplace}
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
