package car

type Car struct {
	PlateNo          string
	Brand       string
	Model       string
	Year        int
	RentedBy    string
	IsRented    bool
}

func New(plateNo, brand, model string, year int) *Car {
	return &Car{
		PlateNo: plateNo,
		Brand: brand,
		Model:model,
		Year:year,
		IsRented: false,
	}
}
