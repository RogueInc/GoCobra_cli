package ui

import (
    "fmt"
    "github.com/rivo/tview"
    "github.com/RogueInc/golcli/cmd"
)

func StartUI() {
    app := tview.NewApplication()
    list := tview.NewList()

    list.AddItem("Add a new to-do", "", 'a', func() {
        addTask(app, list)
    }).AddItem("List all to-dos", "", 'l', func() {
        listTasks(app)
    }).AddItem("Delete a to-do", "", 'd', func() {
        deleteTask(app)
    }).AddItem("Exit", "", 'e', func() {
        app.Stop()
    })

    if err := app.SetRoot(list, true).Run(); err != nil {
        panic(err)
    }
}

func addTask(app *tview.Application, list *tview.List) {
    form := tview.NewForm()
    var taskField *tview.InputField

    form.AddInputField("Task", "", 20, nil, func(text string) {
        taskField = form.GetFormItemByLabel("Task").(*tview.InputField)
    })

    form.AddButton("Add", func() {
        task := taskField.GetText()
        todos, err := cmd.LoadToDos()
        if err != nil {
            fmt.Println("Error loading to-dos:", err)
            return
        }

        todos = append(todos, cmd.ToDo{Task: task})
        if err := cmd.SaveToDos(todos); err != nil {
            fmt.Println("Error saving to-dos:", err)
            return
        }

        app.SetRoot(list, true)
    })

    form.AddButton("Cancel", func() {
        app.SetRoot(list, true)
    })

    app.SetRoot(form, true)
}

func listTasks(app *tview.Application) {
    todos, err := cmd.LoadToDos()
    if err != nil {
        fmt.Println("Error loading to-dos:", err)
        return
    }

    list := tview.NewList()
    for i, todo := range todos {
        list.AddItem(todo.Task, "", rune('a'+i), nil)
    }

    list.AddItem("Back to menu", "", 'b', func() {
        StartUI() // Redirects back to the main menu
    })

    app.SetRoot(list, true)
}

func deleteTask(app *tview.Application) {
    todos, err := cmd.LoadToDos()
    if err != nil {
        fmt.Println("Error loading to-dos:", err)
        return
    }

    list := tview.NewList()
    for i, todo := range todos {
        index := i
        list.AddItem(todo.Task, "", rune('a'+i), func() {
            todos = append(todos[:index], todos[index+1:]...)
            if err := cmd.SaveToDos(todos); err != nil {
                fmt.Println("Error saving to-dos:", err)
                return
            }
            StartUI() // Redirects back to the main menu after deletion
        })
    }

    list.AddItem("Back to menu", "", 'b', func() {
        StartUI() // Redirects back to the main menu
    })

    app.SetRoot(list, true)
}
