package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "mytodoapp",
    Short: "My To-Do CLI Application",
    Long:  `A simple CLI application to manage your to-dos.`,
    Run: func(cmd *cobra.Command, args []string) {
        mainMenu()
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func mainMenu() {
    for {
        fmt.Println("\nMy To-Do CLI Application")
        fmt.Println("1. Add a new to-do")
        fmt.Println("2. List all to-dos")
        fmt.Println("3. Delete a to-do")
        fmt.Println("4. Exit")

        var choice int
        fmt.Print("Enter your choice: ")
        fmt.Scan(&choice)

        switch choice {
        case 1:
            addTask()
        case 2:
            listTasks()
        case 3:
            deleteTask()
        case 4:
            fmt.Println("Exiting...")
            os.Exit(0)
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}
