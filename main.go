package main

import (
	"fmt"
	"os"
)

// const projectPath = "/Users/user/projects/directory/" // Change this line to match the path to your projects.

func main() {
	args := os.Args[1:]

	// Prints help menu if no project name is given.
	if len(args) <= 0 {
		help()
		return
	}

	// Passes args through to create function so that a new project can be created.
	if len(args) >= 1 {
		create(args)
		return
	}
}

// help displays the creation scripts menu to the user.
func help() {
	fmt.Println("\nUsage:")
	fmt.Println("")
	fmt.Println("create <project-name> 		: Creates new project")
	fmt.Println("")
	fmt.Println("create <project-name> -t 		: Creates a new project using TypeScript")
	fmt.Println("")
}

// create takes the users input and creates a new react based project.
func create(args []string) {
	switch args[0] {
	case "-h":
		printMessage("Here are your options")
		help()
	case "-t":
		printMessage("Error: No project name provided.")
		help()
	default:
		projectName := args[0]
		printMessage("Attempting to create project: " + projectName)
	}
}

// printMessage is a reusable function that prints a formatted message to the console.
func printMessage(msg string) {
	fmt.Println("****************************************************")
	fmt.Println("*")
	fmt.Println("* " + msg)
	fmt.Println("*")
	fmt.Println("****************************************************")
}
