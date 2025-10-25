// little programme to show how maps are not concurrent safe in Go
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {

	m := make(map[string]int)

	// Pattern for checking if a key exists before accessing its value

	if val, exists := m["a"]; !exists {
		log.Println("key \"a\" does not exists")
	} else {
		fmt.Printf("value is %d\n", val)
	}

	// How to run into a race condition
	// this will prove maps are not concurrent safe by themselves - see Mutexes chapter

	var wg sync.WaitGroup

	for i := range 100 {
		wg.Add(1)
		go func(i int) {
			time.Sleep(time.Second * 1)
			m[fmt.Sprintf("Key-%d", i)] = i
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println("Map m: ", m)

}
