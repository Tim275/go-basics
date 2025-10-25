package main

import (
	"errors"
	"fmt"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (*Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (tM *truckManager) AddTruck(id string, c int) error {

	// Add a Truck{ID : id, cargo:c} as a value to key "Truck-%d",id in the trucks
	// map of struct truckManager

	tM.trucks[fmt.Sprintf("Truck-%v", id)] = &Truck{ID: id, Cargo: c}
	return nil
}

func (tM *truckManager) GetTruck(id string) (*Truck, error) {

	// Fetch a Truck{ID : id, cargo:c} from the trucks
	// map of struct truckManager if the id exists, otherwise return ErrTruckNotFound

	if t, exists := tM.trucks[fmt.Sprintf("Truck-%v", id)]; !exists {
		return nil, ErrTruckNotFound
	} else {
		return t, nil
	}
}

func (tM *truckManager) RemoveTruck(id string) error {

	// Remove a Truck{ID : id, cargo:c} from the trucks
	// map of struct truckManager if the id exists, otherwise return ErrTruckNotFound

	if _, exists := tM.trucks[fmt.Sprintf("Truck-%v", id)]; !exists {
		return ErrTruckNotFound
	} else {
		delete(tM.trucks, fmt.Sprintf("Truck-%v", id))
		return nil
	}
}

func (tM *truckManager) UpdateTruckCargo(id string, c int) error {

	// update the cargo of a Truck{ID : id, cargo:c} from the trucks
	// map of struct truckManager if the id exists, 
	// otherwise return ErrTruckNotFound


	if _, exists := tM.trucks[fmt.Sprintf("Truck-%v", id)]; !exists {
		return ErrTruckNotFound
	} else {
		tM.trucks[fmt.Sprintf("Truck-%v", id)].Cargo = c
		return nil
	}
}
