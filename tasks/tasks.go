package tasks

import ()

// Create a structure to represent the a task;
type Task struct {
	Title    string
	Starts   string
	Ends     string
	Priority int
}

// Create a structure to represent the queue:
type TodoList struct{
	Tasks []*Task
}

// Instantiote a new Queue:
func NewQueue() *TodoList {
	queue := new(TodoList)
	return queue
}

// Instantiate a new task:
func newTask()