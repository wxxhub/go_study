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

func worker(id int, jobs1 <-chan *job, jobs2 <-chan *job)  {
	for {
		select {
		case i := <- jobs1:
			fmt.Printf("I'm worker[%d], I'm doing jobs1, I doing work %s, it's job_type %d, it's msg %s\n", id, i.name, i.job_type, i.msg)
			time.Sleep(time.Second)
			fmt.Printf("worker[%d] finished work.\n", id)
		case i := <- jobs2:
			fmt.Printf("I'm worker[%d], I'm doing jobs2, I doing work %s, it's job_type %d, it's msg %s\n", id, i.name, i.job_type, i.msg)
			time.Sleep(time.Second)
			fmt.Printf("worker[%d] finished work.\n", id)
		default:
			fmt.Println("错误操作")
		}
	}
}

func main()  {
	jobs1 := make(chan *job, 20)
	jobs2 := make(chan *job, 20)

	for i := 0; i < 5; i++ {
		go worker(i, jobs1, jobs2)
	}
	
	for i := 0; i < 100; i++ {
		work_name := fmt.Sprintf("work%d", i)
		a_job := job {
			work_name,
			i,
			"test1",
		}

		jobs1 <- &a_job
	}

	for i := 100; i < 200; i++ {
		work_name := fmt.Sprintf("work%d", i)
		a_job := job {
			work_name,
			i,
			"test2",
		}

		jobs2 <- &a_job
	}
}