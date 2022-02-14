package manager

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/ton3s/orchestrator/task"
)

/*
- Accept requests from users to start and stop tasks
- Schedule tasks onto worker machines
- Keep track of tasks, states and the machines they run on
*/

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]task.Task
	EventDb       map[string][]task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

// Determine appropriate worker
func (m *Manager) SelectWorker() {
	fmt.Println("I will select an appropriate worker based on the task requirements")
}

// Keep track of tasks, states and the machines they run on
func (m *Manager) UpdateTasks() {
	fmt.Println("I will update tasks")
}

// Send tasks to worker
func (m *Manager) SendWork() {
	fmt.Println("I will send work to workers")
}
