package main

import (
	"os"
	"testing"
)

func TestAddTask(t *testing.T) {
	fileName := "test_tasks.json"
	defer os.Remove(fileName)

	if err := createFile(fileName); err != nil {
		t.Fatalf("Failed to set up test: %v", err)
	}

	if err := addTask("Test task", fileName); err != nil {
		t.Fatalf("Failed to add task: %v", err)
	}

	if err := loadTasks(fileName); err != nil {
		t.Fatalf("Failed to load tasks: %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Description != "Test task" {
		t.Fatalf("Expected 'Test task', got '%s'", tasks[0].Description)
	}

}
