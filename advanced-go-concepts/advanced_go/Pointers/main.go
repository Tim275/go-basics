package main

import "log"

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
func main() {

	t := NormalTruck{cargo: 0}
	log.Printf("Address of t:%p in main()\n", &t)
	log.Printf("value of t.cargo:%d in main before calling functions\n", t.cargo)
	FillTruckCargo(t)
	log.Printf("value of t.cargo after FillTruckCargo should be unchanged : %d ", t.cargo)
	t, _ = FillTruckCargo_B(t)
	log.Printf("value of t.cargo after FillTruckCargo_B should be changed (100): %d ", t.cargo)
	// reseting t.cargo
	t.cargo = 0
	log.Printf("value of t.cargo:%d in main reset to 0.\n", t.cargo)
	FillTruckCargo_C(&t)
	log.Printf("value of t.cargo after FillTruckCargo_C should be changed (100): %d ", t.cargo)

}

func FillTruckCargo(t NormalTruck) error {

	t.cargo = 100
	log.Printf("Address of t:%p in FillTruckCargo\n", &t)

	return nil
}

func FillTruckCargo_B(t NormalTruck) (NormalTruck, error) {

	t.cargo = 100
	log.Printf("Address of t:%p in FillTruckCargo_B\n", &t)

	return t, nil
}

func FillTruckCargo_C(t *NormalTruck) error {

	t.cargo = 100
	log.Printf("Address of t:%p in FillTruckCargo_C\n", t)

	return nil
}
