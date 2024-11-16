package tasks

import "fmt"

// Create a structure to represent the a task;
type Task struct {
	Title    string
	Priority int
}

// Create a structure to represent the queue:
type TodoList struct {
	Tasks []*Task
}

// Create a structure to repsent completed tasks:
type Done struct {
	Tasks []*Task
}

// Instantiote a new Queue:
func NewToDoList() *TodoList {
	queue := new(TodoList)
	return queue
}

// Create a function to force priority using the heap property:
func (list *TodoList) Heapify() {
	index := len(list.Tasks) - 1
	for index != 0 && list.Tasks[index].Priority > list.Tasks[(index-1)/2].Priority {
		list.Tasks[(index-1)/2], list.Tasks[index] = list.Tasks[index], list.Tasks[(index-1)/2]
		index = (index - 1) / 2
	}
}

// Instantiate a new task:
func (list *TodoList) EnqueuetTask(title string, priority int) {
	task := new(Task)
	task.Title = title
	task.Priority = priority
	list.Tasks = append(list.Tasks, task)
	list.Heapify()
}

// Extract the peak:
func (list *TodoList) Peak() *Task {
	if len(list.Tasks) == 0 {
		fmt.Println("Queue is empty!!!")
		return nil
	}
	return list.Tasks[0]
}

// Display my tasks:
func (list *TodoList) Display() {
	if len(list.Tasks) == 0 {
		fmt.Println("Queue is empty!!!")
		return
	}
	for i:=0; i<len(list.Tasks); i++ {
		fmt.Printf("The %dth task is: %s and the priority is: %d\n", (i+1), list.Tasks[i].Title,  list.Tasks[i].Priority)
	}
}

// Heapify down when deleting the first element in the queue to keep the heap property:
func (list *TodoList) HeapifyDown(i int) {
Max := i
Left := (i*2) + 1
Right := (i*2) + 2
if Left < len(list.Tasks) && list.Tasks[Left].Priority > list.Tasks[Max].Priority {
	Max = Left
}
if Right < len(list.Tasks) && list.Tasks[Right].Priority > list.Tasks[Max].Priority {
	Max = Right
}
if i != Max {
	list.Tasks[i], list.Tasks[Max] = list.Tasks[Max], list.Tasks[i]
	list.HeapifyDown(Max)
}
}

// Extract the peak from the queue:
func (list *TodoList) Dequeue() *Task {
	max := list.Tasks[0]
	list.Tasks[0] = list.Tasks[len(list.Tasks)-1]
	list.Tasks = list.Tasks[:len(list.Tasks)-1]
	list.HeapifyDown(0)
	return max
}