package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("Not implemented")
	ErrTruckNotFound  = errors.New("Truck Not Found")
)

type Truck struct {
	Id string
}

func (t *Truck) LoadCargo() error {
	return nil

}

func (t *Truck) UnloadCargo() error {
	return nil
}

func processTruck(truck Truck) error {
	fmt.Printf("Processing truck : %s\n", truck.Id)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	return nil

}

func main() {
	trucks := []Truck{
		{Id: "Truck 1"},
		{Id: "Truck 2"},
		{Id: "Truck 3"},
	}
	for _, truck := range trucks {

		fmt.Printf("Truck %s has arrived \n", truck.Id)

		// processing errors
		if err := processTruck(truck); err != nil {

			if errors.Is(err, ErrNotImplemented) {
				// do this
			}

			if errors.Is(err, ErrTruckNotFound) {
				// do that
			}
			// etc.

			// ALTERNATIVELY :

			switch err {
			case ErrNotImplemented:
				// do this
			case ErrTruckNotFound:
				// do that
			// etc.
			default:
			}

		// OR
		log.Fatalf("Error Processing truck: %s", err)
		}

	}

}
