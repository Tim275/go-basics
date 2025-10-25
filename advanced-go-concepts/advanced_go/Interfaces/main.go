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

type Truck interface{
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	Id string
	cargo int
}
func (t *NormalTruck) LoadCargo() error {
	t.cargo += 2
	return nil
}
func (t *NormalTruck) UnloadCargo() error {
	t.cargo =0
	return nil
}



type ElectricTruck struct {
	Id string
	cargo int
	battery float64
}
func (e *ElectricTruck) LoadCargo() error {
	e.cargo +=2
	e.battery -= 1
	return nil
}
func (e *ElectricTruck) UnloadCargo() error {
	e.cargo =0
	e.battery -= 1
	return nil
}

func processTruck(truck Truck) error {
	fmt.Printf("Processing truck : %+v\n", truck)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}
	return nil

}

func main() {


	nt := &NormalTruck{Id: "1"}
	et := &ElectricTruck{Id: "2"}

	err := processTruck(nt)
	if err != nil {
		log.Fatalf("Error processing truck: %s\n", err)
	}
		
	err = processTruck(et)
	if err != nil {
		log.Fatalf("Error processing truck: %s\n", err)
	}

	log.Println(nt.cargo)
	log.Println(et.battery)

}