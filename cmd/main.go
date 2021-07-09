package main

import (
	"fmt"
	"github.com/p12s/2gis-catalog-api"
)

func main() {
	company := &common.Company{
		Id: 1,
		Name: "Yandex",
		Phones: []string{"22-22-22", "33-33-33"},
		Building: common.Building{
			Address: common.Address{
				City: "Moscow",
				Street: "Lenina",
				House: 1,
			},
			Coordinates: common.Coordinates{
				Longitude: 38.8951,
				Latitude: -77.0364,
			},
		},
		Rubric: []common.Rubric{
			{
				Name: "Интернет-услуги",
			},
			{
				Name: "Услуги доставки",
			},
			{
				Name: "Облачные услуги",
			},
			{
				Name: "Транспортные услуги",
			},
			{
				Name: "Финансовые услуги",
			},
		},
	}
	fmt.Println(company)
}