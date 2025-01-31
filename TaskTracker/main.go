// Implement a task tracking system using maps and slices
// Tasks should have priority, status, and deadlines

package main

import (
	"fmt"
	"time"
)

//structure for the task
type Task struct {
	Content   string
	Priority  string //Very High, High, Medium, Low
	Status    string
	Deadlines time.Time
	Id        int
}

//structure for task manager
type TaskManager struct {
	Tasks  map[int]Task
	NextID int
}

//function for adding the task
func (taskManager *TaskManager) AddingTask(content string, priority string, status string, date string) {
	deadline, _ := time.Parse("2006-01-02", date)

	//created new task
	newTask := Task{
		Content:   content,
		Priority:  priority,
		Status:    status,
		Deadlines: deadline,
		Id:        taskManager.NextID,
	}

	//create a new task  
	taskManager.Tasks[taskManager.NextID] = newTask
	taskManager.NextID++
	fmt.Println("Task added successfully with id ", taskManager.NextID-1)
}

// deleting the task
func (taskManager *TaskManager) DeletingTask(id int) {

	delete(taskManager.Tasks, id)
	fmt.Println("Task added successfully with id ", id)
}

// showing all the tasks
func (taskManager *TaskManager) ShowAllTask() {

	for k, v := range taskManager.Tasks {
		fmt.Printf("%d - %v\n", k, v) //v for struct type
	}
}

//show particular task
func (taskManager *TaskManager) ShowTask(id int) {

	_, present := taskManager.Tasks[id]

	if !present {
		fmt.Println("No task found with this id")
	} else {
		fmt.Printf("Task with id %d - %v\n", id, taskManager.Tasks[id])
	}
}

//function for updating task with the id
func (taskManager *TaskManager) UpdateTask(content string, priority string, status string, date string, id int) {
	_, present := taskManager.Tasks[id]

	//if task not found then return
	if !present {
		fmt.Println("No task found with this id")
		return
	}

	//creating a New Task
	deadline, _ := time.Parse("2006-01-02", date)
	newTask := Task{
		Content:   content,
		Priority:  priority,
		Status:    status,
		Deadlines: deadline,
		Id:        id,
	}

	taskManager.Tasks[id] = newTask
}

func main() {
	//creating a new task
	taskManager := TaskManager{Tasks: make(map[int]Task)}

	// content string,priority int,status string,date string
	taskManager.AddingTask("Task-1", "Very High", "Completed", "2025-01-20")
	taskManager.AddingTask("Task-2", "Medium", "Pending", "2025-01-21")
	taskManager.AddingTask("Task-3", "Low", "Doing", "2025-01-22")
	taskManager.AddingTask("Task-4", "High", "Doing", "2025-01-19")
	fmt.Println()

	
	taskManager.ShowTask(0)
	taskManager.UpdateTask("Updated-Task-1", "Very High", "Completed", "2025-01-20", 0)
	taskManager.ShowTask(0)

	// taskManager.ShowAllTask()

}
