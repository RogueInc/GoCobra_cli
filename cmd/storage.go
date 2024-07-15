package cmd

import (
    "encoding/json"
    "os"
)

const todoFile = "todos.json"

type ToDo struct {
    Task string `json:"task"`
}

func LoadToDos() ([]ToDo, error) {
    var todos []ToDo
    if _, err := os.Stat(todoFile); os.IsNotExist(err) {
        return todos, nil
    }

    file, err := os.ReadFile(todoFile)
    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(file, &todos)
    if err != nil {
        return nil, err
    }

    return todos, nil
}

func SaveToDos(todos []ToDo) error {
    file, err := json.MarshalIndent(todos, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(todoFile, file, 0644)
}
