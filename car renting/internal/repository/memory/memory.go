package memory

import (
	"car-rent/internal/model/car"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type RepoMemory map[string]*car.Car

func New() *RepoMemory {
	rm := make(RepoMemory)
	return &rm
}

func (rm *RepoMemory) GetCar(_ contractapi.TransactionContextInterface, plateNo string) (*car.Car, error) {
	car, existing := (*rm)[plateNo]
	if !existing {
		return nil, fmt.Errorf("no car with plateNo: %s", plateNo)
	}
	return car, nil
}

func (rm *RepoMemory) GetAllCars(_ contractapi.TransactionContextInterface) ([]*car.Car, error) {
	var cars []*car.Car

	for _, v := range *rm {
		cars = append(cars, v)
	}

	if len(cars) == 0 {
		return nil, fmt.Errorf("no cars found")
	}

	return cars, nil
}

func (rm *RepoMemory) PutCar(_ contractapi.TransactionContextInterface, car *car.Car) error {
	(*rm)[car.PlateNo] = car
	return nil
}
