package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	fmt.Printf("Worker %d waiting for jobs\n", id)
	for j := range jobs {
		fmt.Println("worker", id, "started job")
		time.Sleep(time.Second)
		fmt.Println("worker", id, " finished jon")
		results <- j * 2
	}
}

//alternative implementation with Waiting groups - see in main()
func worker_wg(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d waiting for jobs\n", id)
	for j := range jobs {
		fmt.Println("worker", id, "started job")
		time.Sleep(time.Second)
		fmt.Println("worker", id, " finished jon")
		results <- j * 2
	}
}

func main() {

	// Creates buffered channels for 5 jobs and their results
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Creates a pool of 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	time.Sleep(time.Second * 2)
	// filling the job channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)
	// Ensures that workers process all the jobs and send their results
	// Otherwise would jump to the next section
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	close(results)

	// alternative implementation using Waiting group in order to 
	// wait for the work to complete their jobs

	fmt.Println("Alternatively ...")

	var wg sync.WaitGroup
	jobs = make(chan int, numJobs)
	results = make(chan int, numJobs)

	// Creates a pool of 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker_wg(w, jobs, results, &wg)
	}

	time.Sleep(time.Second * 2)
	// filling the job channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	
	wg.Wait()

	close(results)
}
