package main

import (
	"fmt"
	"time"
)

type job struct {
	name string
	job_type int
	msg string
}

func worker(id int, jobs <-chan *job)  {
	for i := range jobs {
		fmt.Printf("I'm worker[%d], I doing work %s, it's job_type %d, it's msg %s\n", id, i.name, i.job_type, i.msg)
		time.Sleep(time.Second)
		fmt.Printf("worker[%d] finished work.\n", id)
	}
}

func main()  {
	jobs := make(chan *job, 10)

	for i := 0; i < 5; i++ {
		go worker(i, jobs)
	}
	
	for i := 0; i < 255; i++ {
		work_name := fmt.Sprintf("work%d", i)
		a_job := job {
			work_name,
			i,
			"test",
		}

		jobs <- &a_job
	}
}