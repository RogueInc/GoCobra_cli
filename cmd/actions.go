package cmd

import (
    "fmt"
    "strconv"
)

func addTask() {
    var task string
    fmt.Print("Enter the new task: ")
    fmt.Scanln(&task)

    todos, err := LoadToDos()
    if err != nil {
        fmt.Println("Error loading to-dos:", err)
        return
    }

    todos = append(todos, ToDo{Task: task})
    if err := SaveToDos(todos); err != nil {
        fmt.Println("Error saving to-dos:", err)
        return
    }

    fmt.Printf("Added new to-do: %s\n", task)
}

func listTasks() {
    todos, err := LoadToDos()
    if err != nil {
        fmt.Println("Error loading to-dos:", err)
        return
    }

    fmt.Println("Listing all to-dos:")
    for i, todo := range todos {
        fmt.Printf("%d. %s\n", i+1, todo.Task)
    }
}

func deleteTask() {
    todos, err := LoadToDos()
    if err != nil {
        fmt.Println("Error loading to-dos:", err)
        return
    }

    listTasks()
    var taskNumber string
    fmt.Print("Enter the number of the task to delete: ")
    fmt.Scan(&taskNumber)
    taskIndex, err := strconv.Atoi(taskNumber)
    if err != nil || taskIndex <= 0 || taskIndex > len(todos) {
        fmt.Println("Invalid task number.")
        return
    }

    todos = append(todos[:taskIndex-1], todos[taskIndex:]...)
    if err := SaveToDos(todos); err != nil {
        fmt.Println("Error saving to-dos:", err)
        return
    }

    fmt.Printf("Deleted to-do number: %d\n", taskIndex)
}
