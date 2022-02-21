package main

import (
	"fmt"
	"time"

	"github.com/ton3s/orchestrator/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"

	"github.com/ton3s/orchestrator/worker"
)

func main() {

	// Initialize task db and setup worker
	db := make(map[uuid.UUID]*task.Task)
	w := worker.Worker{
		Queue: *queue.New(),
		Db:    db,
	}

	// Setup task to schdule to run on the worker
	t := task.Task{
		ID:    uuid.New(),
		Name:  "test-container-1",
		State: task.Scheduled,
		Image: "strm/helloworld-http",
	}

	// first time the worker will see the task
	fmt.Println("starting task")
	w.AddTask(t)          // Add the task to the queue
	result := w.RunTask() // Run the task from the queue
	if result.Error != nil {
		panic(result.Error)
	}

	t.ContainerID = result.ContainerId

	fmt.Printf("task %s is running in container %s\n", t.ID, t.ContainerID)
	fmt.Println("Sleepy time")
	time.Sleep(time.Second * 30)

	fmt.Printf("stopping task %s\n", t.ID)
	t.State = task.Completed

	w.AddTask(t)         // Add new task to stop the container from running
	result = w.RunTask() // Run the task from the queue
	if result.Error != nil {
		panic(result.Error)
	}
}
