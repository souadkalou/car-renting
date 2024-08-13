package repository

import (
	"car-rent/internal/model/car"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Repository interface {
	GetCar(contractapi.TransactionContextInterface, string) (*car.Car, error)
	GetAllCars(contractapi.TransactionContextInterface) ([]*car.Car, error)
	PutCar(contractapi.TransactionContextInterface, *car.Car) error
}
