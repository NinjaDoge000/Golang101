package main

import (
	"fmt"
	"sync"
	"time"
)

type Task interface {
	Process()
}

type EmailTask struct {
	email       string
	subject     string
	messageBody string
}

func (t *EmailTask) Process() {
	fmt.Println("Sending Email", t.email)
	time.Sleep(time.Second)
}

type ImageProcessingTask struct {
	imagePath string
}

func (t *ImageProcessingTask) Process() {
	fmt.Println("Processing Image", t.imagePath)
	time.Sleep(2 * time.Second)
}

type WorkerPool struct {
	tasks       []Task
	concurrency int
	taskChan    chan Task
	wg          sync.WaitGroup
}

func (wp *WorkerPool) worker() {

	// worker who takes in task
	for task := range wp.taskChan {
		defer wp.wg.Done()
		task.Process()
	}

}

func (wp *WorkerPool) run() {

	wp.taskChan = make(chan Task, len(wp.tasks))

	// start workers
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	wp.wg.Add(len(wp.tasks))

	// send tasks to tasks channel
	for _, task := range wp.tasks {
		wp.taskChan <- task
	}

	close(wp.taskChan)

	// wait for all task to complete
	wp.wg.Wait()
}

func main() {

	// create tasks
	myTasks := []Task{
		&EmailTask{email: "abc@email.com", subject: "test", messageBody: "test"},
		&ImageProcessingTask{imagePath: "/image1.jpg"},
		&EmailTask{email: "efg@email.com", subject: "test", messageBody: "test"},
		&ImageProcessingTask{imagePath: "/image2.jpg"},
		&ImageProcessingTask{imagePath: "/image3.jpg"},
		&EmailTask{email: "hij@email.com", subject: "test", messageBody: "test"},
		&EmailTask{email: "lmn@email.com", subject: "test", messageBody: "test"},
	}

	wp := WorkerPool{
		tasks:       myTasks,
		concurrency: 3,
	}

	wp.run()
	fmt.Println("all process completed")

}
