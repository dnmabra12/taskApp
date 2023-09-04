package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Task struct {
	Description string
}

var tasks []Task

func loadTasks(fileName string) error {
	log.Println("Loading tasks from", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to load tasks from %s: %w", fileName, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&tasks)
}

func saveTasks(fileName string) error {
	log.Println("Saving tasks to", fileName)
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to load tasks from %s: %w", fileName, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}

func createFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte("[]"))
	return err
}

func addTask(description string, fileName string) error {
	if err := loadTasks(fileName); err != nil {
		return err
	}
	tasks = append(tasks, Task{Description: description})
	return saveTasks(fileName)
}

func listTasks(fileName string) error {
	if err := loadTasks(fileName); err != nil {
		return err
	}
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task.Description)
	}
	return nil
}

// Removal of tasks
func removeTask(identifier string, fileName string, byName bool) error {
	if err := loadTasks(fileName); err != nil {
		return err
	}

	if byName {
		found := false
		for i, task := range tasks {
			if task.Description == identifier {
				tasks = append(tasks[:i], tasks[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			return errors.New("Task not found")
		}
	} else {
		index, err := strconv.Atoi(identifier)
		if err != nil || index < 1 || index > len(tasks) {
			return errors.New("Invalid task index")
		}
		tasks = append(tasks[:index-1], tasks[index:]...)
	}

	return saveTasks(fileName)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Missing command or filename")
		return
	}

	command := os.Args[1]
	fileName := os.Args[2]

	switch command {
	case "createFile":
		if err := createFile(fileName); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Successful creation of the file : ", fileName)
		}
	case "addTask":
		if len(os.Args) < 4 {
			fmt.Println("Missing task description")
			return
		}
		description := os.Args[3]
		if err := addTask(description, fileName); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Successfully added: ", description, " to ", fileName)
		}
	case "listTasks":
		if err := listTasks(fileName); err != nil {
			fmt.Println("Error:", err)
		}
	case "removeTask":
		if len(os.Args) < 5 {
			fmt.Println("Missing task identifier and mode")
			return
		}
		identifier := os.Args[3]
		mode := os.Args[4]
		byName := false
		if mode == "name" {
			byName = true
		} else if mode != "index" {
			fmt.Println("Invalid mode. Use 'name' or 'index'.")
			return
		}

		if err := removeTask(identifier, fileName, byName); err != nil {
			fmt.Println("Error:", err)
		}
	default:
		fmt.Println("Invalid command")
	}
}
