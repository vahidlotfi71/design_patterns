package main

import (
	"fmt"
	"time"
)

// تابع worker که کارها را پردازش می‌کند
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // شبیه‌سازی یک کار زمان‌بر
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // ارسال نتیجه به کانال results
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	// ایجاد کانال‌ها برای کارها و نتایج
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// ایجاد کارگران
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// ارسال کارها به کانال jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // بستن کانال jobs پس از ارسال همه کارها

	// دریافت نتایج از کانال results
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}