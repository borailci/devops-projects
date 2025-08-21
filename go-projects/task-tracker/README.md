# Task Tracker CLI

A simple command-line interface (CLI) application for tracking and managing your tasks. Built with Go, this tool helps you manage what you need to do, what you're currently working on, and what you have completed.

## Features

- **Add tasks**: Create new tasks with descriptions
- **Update tasks**: Modify existing task descriptions
- **Delete tasks**: Remove tasks you no longer need
- **Status management**: Mark tasks as todo, in-progress, or done
- **List tasks**: View all tasks or filter by status
- **Persistent storage**: Tasks are stored in a JSON file
- **Error handling**: Graceful handling of errors and edge cases

## Installation

1. Make sure you have Go installed on your system (Go 1.19 or later)
2. Clone or download this project
3. Navigate to the project directory
4. Build the application:

```bash
go build -o task-cli main.go
```

This will create an executable named `task-cli` in your current directory.

## Usage

### Adding a New Task

```bash
./task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

./task-cli add "Complete project documentation"
# Output: Task added successfully (ID: 2)
```

### Updating a Task

```bash
./task-cli update 1 "Buy groceries and cook dinner"
# Output: Task updated successfully (ID: 1)
```

### Deleting a Task

```bash
./task-cli delete 1
# Output: Task deleted successfully (ID: 1)
```

### Marking Task Status

Mark a task as in progress:

```bash
./task-cli mark-in-progress 1
# Output: Task marked as in-progress (ID: 1)
```

Mark a task as done:

```bash
./task-cli mark-done 1
# Output: Task marked as done (ID: 1)
```

### Listing Tasks

List all tasks:

```bash
./task-cli list
```

List tasks by status:

```bash
./task-cli list todo        # List pending tasks
./task-cli list in-progress # List tasks in progress
./task-cli list done        # List completed tasks
```

### Getting Help

```bash
./task-cli help
# or
./task-cli --help
# or
./task-cli -h
```

## Task Properties

Each task contains the following properties:

- **id**: A unique identifier for the task
- **description**: A short description of the task
- **status**: The current status (todo, in-progress, done)
- **createdAt**: The date and time when the task was created
- **updatedAt**: The date and time when the task was last updated

## Data Storage

Tasks are stored in a `tasks.json` file in the same directory as the executable. The file is automatically created if it doesn't exist. The JSON structure looks like this:

```json
[
  {
    "id": 1,
    "description": "Buy groceries",
    "status": "todo",
    "createdAt": "2025-08-21T14:30:00Z",
    "updatedAt": "2025-08-21T14:30:00Z"
  },
  {
    "id": 2,
    "description": "Complete project documentation",
    "status": "in-progress",
    "createdAt": "2025-08-21T14:31:00Z",
    "updatedAt": "2025-08-21T14:35:00Z"
  }
]
```

## Example Workflow

Here's a typical workflow using the Task Tracker CLI:

```bash
# Add some tasks
./task-cli add "Set up development environment"
./task-cli add "Write project documentation"
./task-cli add "Implement user authentication"
./task-cli add "Write unit tests"

# View all tasks
./task-cli list

# Start working on a task
./task-cli mark-in-progress 1

# Complete a task
./task-cli mark-done 1

# Update a task description
./task-cli update 2 "Write comprehensive project documentation with examples"

# View tasks by status
./task-cli list done
./task-cli list in-progress
./task-cli list todo

# Delete a task
./task-cli delete 4
```

## Error Handling

The application handles various error scenarios gracefully:

- **Invalid commands**: Shows usage information for unknown commands
- **Missing arguments**: Provides specific error messages for missing required arguments
- **Invalid task IDs**: Shows error for non-numeric or non-existent task IDs
- **Empty descriptions**: Prevents adding or updating tasks with empty descriptions
- **Invalid status**: Shows error for invalid status values
- **File system errors**: Handles JSON file read/write errors

## Command Summary

| Command                       | Description              | Example                               |
| ----------------------------- | ------------------------ | ------------------------------------- |
| `add "<description>"`         | Add a new task           | `./task-cli add "Buy milk"`           |
| `update <id> "<description>"` | Update task description  | `./task-cli update 1 "Buy groceries"` |
| `delete <id>`                 | Delete a task            | `./task-cli delete 1`                 |
| `mark-in-progress <id>`       | Mark task as in progress | `./task-cli mark-in-progress 1`       |
| `mark-done <id>`              | Mark task as completed   | `./task-cli mark-done 1`              |
| `list`                        | List all tasks           | `./task-cli list`                     |
| `list <status>`               | List tasks by status     | `./task-cli list done`                |
| `help`                        | Show usage information   | `./task-cli help`                     |

## Development

This project is built with Go and uses only standard library packages:

- `encoding/json` for JSON serialization
- `fmt` for formatted I/O
- `io/ioutil` for file operations
- `os` for command-line arguments and file system operations
- `strconv` for string conversions
- `strings` for string manipulation
- `time` for timestamps

## Contributing

Feel free to contribute to this project by:

1. Reporting bugs
2. Suggesting new features
3. Submitting pull requests
4. Improving documentation

## License

This project is open source and available under the MIT License.
