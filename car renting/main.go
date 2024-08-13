package main

import (
	"car-rent/internal/handler" // Update import path
	"car-rent/internal/model/car"
	"car-rent/internal/repository/memory"
	"fmt"
)

// func mainChaincode() {
// 	chaincode, err := contractapi.NewChaincode(new(handler.CarHandler))

// 	if err != nil {
// 		fmt.Printf("error creating car chaincode: %v", err)
// 		return
// 	}

// 	if err := chaincode.Start(); err != nil {
// 		fmt.Printf("error starting car chaincode: %v", err)
// 	}
// }

func main() {
	//mainChaincode()
	handler := handler.CarHandler{
		RepoContract: *memory.New(),
	}
	err := handler.RegisterCar(nil, "abc111", "brandtest", "modeltest", 2022)
	if err != nil {
		panic(err)
	}
	var car *car.Car
	car, err = handler.GetCar(nil, "abc111")
	if err != nil {
		panic(err)

	}
	fmt.Printf("Car %#v\n", car)
}
