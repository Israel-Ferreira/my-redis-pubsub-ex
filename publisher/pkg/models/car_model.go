package models

import "github.com/google/uuid"

type CarModel struct {
	Model            string  `json:"model"`
	Manufacturer     string  `json:"manufacturer"`
	Description      string  `json:"description"`
	LicenseNumber    string  `json:"license_number"`
	Year             uint    `json:"year"`
	KilometersDriven uint    `json:"kilometers_driven"`
	Price            float64 `json:"price"`
}

type Car struct {
	UUID uuid.UUID `json:"uid"`
	CarModel
}

func NewCar(model string, manufacturer string, description string, licenseNumber string, year uint, kmDriven uint, price float64) *Car {
	return &Car{
		UUID: uuid.New(),
		CarModel: CarModel{
			Model:            model,
			Manufacturer:     manufacturer,
			Description:      description,
			KilometersDriven: kmDriven,
			LicenseNumber:    licenseNumber,
			Year:             year,
			Price:            price,
		},
	}
}
