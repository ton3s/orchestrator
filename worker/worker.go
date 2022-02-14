package worker

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/ton3s/orchestrator/task"
)

/*
- Run tasks as docker containers
- Accept tasks to run from manager
- Provide relevant statistics
- Keep track of tasks and their state
*/

type Worker struct {
	Queue     queue.Queue
	Db        map[uuid.UUID]task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("I will collect stats")
}

func (w *Worker) RunTask() {
	fmt.Println("I will start or stop a task")
}

func (w *Worker) StartTask() {
	fmt.Println("I will start a task")
}

func (w *Worker) StopTask() {
	fmt.Println("I will stop a task")
}
