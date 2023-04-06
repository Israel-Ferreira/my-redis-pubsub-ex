package models

type CarDTO struct {
	CarModel
}

func (cdto CarDTO) ToCar() *Car {
	return NewCar(
		cdto.Model,
		cdto.Manufacturer,
		cdto.Description,
		cdto.LicenseNumber,
		cdto.Year,
		cdto.KilometersDriven,
		cdto.Price,
	)
}
