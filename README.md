# CLI Todo — Go

A simple command-line todo manager written in Go. Todos are stored locally in a `todos.json` file and displayed in a formatted table.

## Features

- Add, edit, delete, and list todos
- Toggle completion status with timestamps
- Persistent storage via JSON

## Requirements

- Go 1.21+

## Installation

```bash
git clone https://github.com/iamchitta07/CLI-TODO-GO.git
cd todo
go build -o todo
```

## Usage

```bash
# Add a new todo
./todo -add "Buy groceries"

# List all todos
./todo -list

# Toggle a todo's completion status (by id)
./todo -toggle 0

# Edit a todo (format: id:new_title)
./todo -edit "0:Buy groceries and cook dinner"

# Delete a todo (by id)
./todo -del 0
```

## Example Output

```
+----+--------------------+-----------+--------------------------+--------------------------+
| Id | Title              | Completed | Created At               | Completed At             |
+----+--------------------+-----------+--------------------------+--------------------------+
| 0  | Buy groceries      | ✔️        | Mon, 26 May 2026 10:00   | Mon, 26 May 2026 10:30   |
| 1  | Write unit tests   | ❌        | Mon, 26 May 2026 11:00   | Not completed yet!       |
+----+--------------------+-----------+--------------------------+--------------------------+
```

## Project Structure

```
.
├── main.go      # Entry point — loads storage, parses flags, saves
├── todo.go      # Todo struct and CRUD operations
├── command.go   # CLI flag definitions and dispatch
├── storage.go   # Generic JSON file storage
└── todos.json   # Local persistent store (auto-created)
```

## Dependencies

- [aquasecurity/table](https://github.com/aquasecurity/table) — terminal table rendering
