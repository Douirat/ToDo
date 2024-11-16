package main

import (
	"fmt"
	"time"

	"github.com/Douirat/ToDo/tasks"
)

// The main function:
func main() {
	// Record the start time
	start := time.Now()

	todoList := tasks.NewToDoList()
	todoList.EnqueuetTask("Reading atomic habits for 30 minuts!", 5)
	todoList.EnqueuetTask("Learning graph data structure components!", 2)
	todoList.EnqueuetTask("Continue the work on the groupie tracker!!!", 4)
	todoList.EnqueuetTask("continue the work on push swap!", 1)
	todoList.EnqueuetTask("Reading atomic habits for 30 minuts!", 7)
	todoList.EnqueuetTask("Data base engineering!", 3)
	todoList.EnqueuetTask("Preparing for the picine of js!", 6)
	todoList.Display()
	peak := todoList.Dequeue()
	fmt.Printf("The most important task is: %s and the priority is: %d\n", peak.Title, peak.Priority)
	fmt.Println()
	todoList.Display()
	// Record the end time
	duration := time.Since(start)
	// Display time in seconds
	fmt.Printf("Execution time: %.6f seconds\n", duration.Seconds())
}
