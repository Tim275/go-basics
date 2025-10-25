package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	//"time"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck Not Found")
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	Id    string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 2
	return nil
}
func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

type ElectricTruck struct {
	Id      string
	cargo   int
	battery float64
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 2
	e.battery -= 1
	return nil
}
func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery -= 1
	return nil
}

func processTruck(truck Truck) error {

	return ErrTruckNotFound
	/*
		 	fmt.Printf("Started processing truck : %+v\n", truck)

			time.Sleep(time.Second)

			if err := truck.LoadCargo(); err != nil {
				return fmt.Errorf("error loading cargo: %w", err)
			}

			if err := truck.UnloadCargo(); err != nil {
				return fmt.Errorf("error loading cargo: %w", err)
			}

			fmt.Printf("Finished processing truck %+v\n", truck)
			return nil
	*/
}

func processFleet(fleet []Truck) error {

	var wg sync.WaitGroup
	errorsChannel := make(chan error, len(fleet))

	for _, t := range fleet {

		wg.Add(1)

		go func(t Truck) {
			if err := processTruck(t); err != nil {
				errorsChannel <- err
			}
			wg.Done()
		}(t)

	}
	wg.Wait()

	close(errorsChannel)
	var errs []error
	for err := range errorsChannel {
		errs = append(errs, err)
		log.Printf("error processing truck : %v\n", err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("Fleet processing had %d errors", len(errs))
	}
	return nil

}

func main() {

	Fleet := []Truck{
		&NormalTruck{Id: "NT1", cargo: 0},
		&ElectricTruck{Id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{Id: "NT2", cargo: 0},
		&ElectricTruck{Id: "ET2", cargo: 0, battery: 100},
	}

	if err := processFleet(Fleet); err != nil {
		log.Fatalf("Error processing fleet %v\n", err)
	}

	fmt.Println("All trucks processed succesfully")

}
