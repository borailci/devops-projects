package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Task represents a single task with all required properties
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// TaskManager handles all task operations
type TaskManager struct {
	tasks    []Task
	filename string
}

// NewTaskManager creates a new TaskManager instance
func NewTaskManager() *TaskManager {
	tm := &TaskManager{
		tasks:    []Task{},
		filename: "tasks.json",
	}
	tm.loadTasks()
	return tm
}

// loadTasks loads tasks from the JSON file
func (tm *TaskManager) loadTasks() error {
	if _, err := os.Stat(tm.filename); os.IsNotExist(err) {
		// File doesn't exist, start with empty tasks
		return nil
	}

	data, err := os.ReadFile(tm.filename)
	if err != nil {
		return fmt.Errorf("error reading tasks file: %v", err)
	}

	if len(data) == 0 {
		// Empty file, start with empty tasks
		return nil
	}

	err = json.Unmarshal(data, &tm.tasks)
	if err != nil {
		return fmt.Errorf("error parsing tasks file: %v", err)
	}

	return nil
}

// saveTasks saves tasks to the JSON file
func (tm *TaskManager) saveTasks() error {
	data, err := json.MarshalIndent(tm.tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling tasks: %v", err)
	}

	err = os.WriteFile(tm.filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing tasks file: %v", err)
	}

	return nil
}

// getNextID returns the next available ID for a new task
func (tm *TaskManager) getNextID() int {
	maxID := 0
	for _, task := range tm.tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

// addTask adds a new task
func (tm *TaskManager) addTask(description string) error {
	if strings.TrimSpace(description) == "" {
		return fmt.Errorf("task description cannot be empty")
	}

	task := Task{
		ID:          tm.getNextID(),
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tm.tasks = append(tm.tasks, task)
	err := tm.saveTasks()
	if err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
	return nil
}

// updateTask updates an existing task's description
func (tm *TaskManager) updateTask(id int, description string) error {
	if strings.TrimSpace(description) == "" {
		return fmt.Errorf("task description cannot be empty")
	}

	for i, task := range tm.tasks {
		if task.ID == id {
			tm.tasks[i].Description = description
			tm.tasks[i].UpdatedAt = time.Now()
			err := tm.saveTasks()
			if err != nil {
				return err
			}
			fmt.Printf("Task updated successfully (ID: %d)\n", id)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

// deleteTask deletes a task by ID
func (tm *TaskManager) deleteTask(id int) error {
	for i, task := range tm.tasks {
		if task.ID == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			err := tm.saveTasks()
			if err != nil {
				return err
			}
			fmt.Printf("Task deleted successfully (ID: %d)\n", id)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

// markTask marks a task with a specific status
func (tm *TaskManager) markTask(id int, status string) error {
	validStatuses := []string{"todo", "in-progress", "done"}
	statusValid := false
	for _, validStatus := range validStatuses {
		if status == validStatus {
			statusValid = true
			break
		}
	}

	if !statusValid {
		return fmt.Errorf("invalid status: %s. Valid statuses are: %s", status, strings.Join(validStatuses, ", "))
	}

	for i, task := range tm.tasks {
		if task.ID == id {
			tm.tasks[i].Status = status
			tm.tasks[i].UpdatedAt = time.Now()
			err := tm.saveTasks()
			if err != nil {
				return err
			}
			fmt.Printf("Task marked as %s (ID: %d)\n", status, id)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

// listTasks lists tasks based on status filter
func (tm *TaskManager) listTasks(statusFilter string) {
	if len(tm.tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	filteredTasks := []Task{}
	for _, task := range tm.tasks {
		if statusFilter == "" || task.Status == statusFilter {
			filteredTasks = append(filteredTasks, task)
		}
	}

	if len(filteredTasks) == 0 {
		if statusFilter != "" {
			fmt.Printf("No tasks found with status: %s\n", statusFilter)
		} else {
			fmt.Println("No tasks found.")
		}
		return
	}

	fmt.Println("ID  | Status      | Description                    | Created At          | Updated At")
	fmt.Println("----|-------------|--------------------------------|---------------------|---------------------")

	for _, task := range filteredTasks {
		createdAt := task.CreatedAt.Format("2006-01-02 15:04:05")
		updatedAt := task.UpdatedAt.Format("2006-01-02 15:04:05")
		description := task.Description
		if len(description) > 30 {
			description = description[:27] + "..."
		}
		fmt.Printf("%-3d | %-11s | %-30s | %s | %s\n",
			task.ID, task.Status, description, createdAt, updatedAt)
	}
}

// printUsage prints the usage information
func printUsage() {
	fmt.Println("Task CLI - A simple command line task tracker")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  task-cli add \"<description>\"           - Add a new task")
	fmt.Println("  task-cli update <id> \"<description>\"   - Update a task")
	fmt.Println("  task-cli delete <id>                   - Delete a task")
	fmt.Println("  task-cli mark-in-progress <id>         - Mark task as in progress")
	fmt.Println("  task-cli mark-done <id>                - Mark task as done")
	fmt.Println("  task-cli list                          - List all tasks")
	fmt.Println("  task-cli list done                     - List completed tasks")
	fmt.Println("  task-cli list todo                     - List pending tasks")
	fmt.Println("  task-cli list in-progress              - List tasks in progress")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  task-cli add \"Buy groceries\"")
	fmt.Println("  task-cli update 1 \"Buy groceries and cook dinner\"")
	fmt.Println("  task-cli mark-done 1")
	fmt.Println("  task-cli list")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	tm := NewTaskManager()
	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task description")
			fmt.Println("Usage: task-cli add \"<description>\"")
			os.Exit(1)
		}
		err := tm.addTask(os.Args[2])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: Please provide task ID and new description")
			fmt.Println("Usage: task-cli update <id> \"<description>\"")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error: Invalid task ID '%s'. Please provide a valid number.\n", os.Args[2])
			os.Exit(1)
		}
		err = tm.updateTask(id, os.Args[3])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID")
			fmt.Println("Usage: task-cli delete <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error: Invalid task ID '%s'. Please provide a valid number.\n", os.Args[2])
			os.Exit(1)
		}
		err = tm.deleteTask(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID")
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error: Invalid task ID '%s'. Please provide a valid number.\n", os.Args[2])
			os.Exit(1)
		}
		err = tm.markTask(id, "in-progress")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID")
			fmt.Println("Usage: task-cli mark-done <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error: Invalid task ID '%s'. Please provide a valid number.\n", os.Args[2])
			os.Exit(1)
		}
		err = tm.markTask(id, "done")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	case "list":
		statusFilter := ""
		if len(os.Args) >= 3 {
			statusFilter = os.Args[2]
			// Validate status filter
			validStatuses := []string{"todo", "in-progress", "done"}
			statusValid := false
			for _, validStatus := range validStatuses {
				if statusFilter == validStatus {
					statusValid = true
					break
				}
			}
			if !statusValid {
				fmt.Printf("Error: Invalid status '%s'. Valid statuses are: %s\n", statusFilter, strings.Join(validStatuses, ", "))
				os.Exit(1)
			}
		}
		tm.listTasks(statusFilter)

	case "help", "--help", "-h":
		printUsage()

	default:
		fmt.Printf("Error: Unknown command '%s'\n", command)
		printUsage()
		os.Exit(1)
	}
}
