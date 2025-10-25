package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type ContextType string

var UserIdKey ContextType = "UserId"

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

func processTruck(ctx context.Context, truck Truck) error {
	fmt.Printf("Started processing truck : %+v\n", truck)

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	//simulate a long running process
	delay := time.Second * 3
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		break
	}

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	fmt.Printf("Finished processing truck %+v\n", truck)
	return nil

}

func processFleet(ctx context.Context, fleet []Truck) error {

	var wg sync.WaitGroup

	for _, t := range fleet {

		wg.Add(1)

		go func(t Truck) {
			if err := processTruck(ctx, t); err != nil {
				log.Println(err)
			}
			wg.Done()
		}(t)

	}
	wg.Wait()

	return nil

}

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, UserIdKey, 42)

	Fleet := []Truck{
		&NormalTruck{Id: "NT1", cargo: 0},
		&ElectricTruck{Id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{Id: "NT2", cargo: 0},
		&ElectricTruck{Id: "ET2", cargo: 0, battery: 100},
	}

	if err := processFleet(ctx, Fleet); err != nil {
		log.Fatalf("Error processing fleet %v", err)
	}

	fmt.Println("All trucks processed succesfully")

}
