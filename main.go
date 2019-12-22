package main

import (
	"fmt"
	"os/exec"
)

func callPython(job int) {
	fmt.Println("Running job", job)
	cmd := exec.Command("python", "script.py")

	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(output))
}

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)

		callPython(j) // time.Sleep(time.Second)

		fmt.Println("worker", id, "finishh job", j)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}
}
