package handler

import (
	"fmt"

	"car-rent/internal/model/car"
	"car-rent/internal/repository/memory"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type CarHandler struct {
	// contract.RepoContract
	RepoContract memory.RepoMemory
}

func (h *CarHandler) GetCar(ctx contractapi.TransactionContextInterface, plateNo string) (*car.Car, error) {
	return h.RepoContract.GetCar(ctx, plateNo)
}

func (h *CarHandler) GetAllCars(ctx contractapi.TransactionContextInterface) ([]*car.Car, error) {
	return h.RepoContract.GetAllCars(ctx)
}

func (h *CarHandler) RegisterCar(ctx contractapi.TransactionContextInterface, plateNo, brand, model string, year int) error {
	_, err := h.GetCar(ctx, plateNo)
	if err == nil {
		return fmt.Errorf("car with PlateNo %s is already registered", plateNo)
	}
	car := car.New(plateNo, brand, model, year)
	return h.RepoContract.PutCar(ctx, car)
}

func (h *CarHandler) RentCar(ctx contractapi.TransactionContextInterface, plateNo, renter string) error {
	car, err := h.GetCar(ctx, plateNo)
	if err != nil {
		return err
	}
	if car.IsRented {
		return fmt.Errorf("car with PlateNo %s is already rented", plateNo)
	}
	car.IsRented = true
	car.RentedBy = renter
	return h.RepoContract.PutCar(ctx, car)
}

func (h *CarHandler) ReturnCar(ctx contractapi.TransactionContextInterface, plateNo string) error {
	car, err := h.GetCar(ctx, plateNo)
	if err != nil {
		return err
	}
	if !car.IsRented {
		return fmt.Errorf("car with PlateNo %s is not rented", plateNo)
	}
	car.IsRented = false
	car.RentedBy = ""
	return h.RepoContract.PutCar(ctx, car)
}
