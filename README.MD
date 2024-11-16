# TodoList Priority Queue in Go

## Project Overview

This project implements a **priority queue** to manage tasks in a `TodoList`. Each task has a title and a priority level. The priority queue ensures that tasks are processed in order of priority (higher priority tasks are processed first). The tasks are stored in a heap-based data structure, which is efficient for adding, removing, and accessing tasks based on priority.

## Priority Queue Explanation

A **priority queue** is a data structure where each element has a "priority" associated with it. Elements are dequeued in the order of their priority, not necessarily the order in which they were added. In this implementation:

- Higher numerical priority values indicate higher priority.
- The queue is implemented as a **max-heap**, meaning that the task with the highest priority is always at the front of the queue.
- Tasks are added to the queue with a specific priority, and tasks with the highest priority are extracted first.

The heap property is maintained by the `Heapify` and `HeapifyDown` methods to ensure the correct task order after additions and removals.

## Key Components

- **Task**: Represents a task with a `Title` and `Priority`.
- **TodoList**: A list of tasks implemented as a max-heap, where tasks with higher priorities are dequeued first.
- **Done**: A list of tasks that have been completed.
