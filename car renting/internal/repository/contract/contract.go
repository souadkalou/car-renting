package contract

import (
	"car-rent/internal/model/car"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type RepoContract struct {
	contractapi.Contract
}

func (r *RepoContract) GetCar(ctx contractapi.TransactionContextInterface, plateNo string) (*car.Car, error) {
	carJSON, err := ctx.GetStub().GetState(plateNo)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: error %v", err)
	}
	if carJSON == nil {
		return nil, fmt.Errorf("car with PlateNo %s does not exist", plateNo)
	}

	var car car.Car
	err = json.Unmarshal(carJSON, &car)
	if err != nil {
		return nil, err
	}

	return &car, nil
}

func (r *RepoContract) GetAllCars(ctx contractapi.TransactionContextInterface) ([]*car.Car, error) {
	carIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: error %v", err)
	}
	defer carIterator.Close()

	var cars []*car.Car

	for carIterator.HasNext() {
		result, err := carIterator.Next()
		if err != nil {
			return nil, err
		}
		var car car.Car
		err = json.Unmarshal(result.Value, &car)
		if err != nil {
			return nil, err
		}

		cars = append(cars, &car)
	}

	return cars, nil
}

func (r *RepoContract) PutCar(ctx contractapi.TransactionContextInterface, car *car.Car) error {
	carJSON, err := json.Marshal(car)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(car.PlateNo, carJSON)
}
