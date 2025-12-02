package main

import (
	"bufio"
	"encoding/json"
	print "fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int `json:"ID"`
	Title       string `json:"Title"`
	AssignedTo  string `json:"AssignedTo"`
	Status      bool `json:"Status"`
	DueDate     time.Time `json:"DueDate"`
	TimeCreated time.Time `json:"TimeCreated"`
}

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Reset  = "\033[0m"
)

var TasksSlice = []Task{}

// Utility Functions
func clearScreen() {
	print.Print("\033[H\033[2J")
}
func hr() {
	print.Println(Green + "------------------" + Reset)
}
func menuOptions() {

	if len(TasksSlice) > 0 {
		print.Println(Green + "Add a New Task -> Press 'A'")
		print.Println("View all Tasks -> Press 'V'")
		print.Println("Delete a Task -> Press 'D'\n")
	print.Println("Quit the application -> Press 'Q'\n" + Reset)
	}else {
		print.Println(Green + "Add a New Task -> Press 'A'")
		print.Println("View all Tasks -> Press 'V'")
		print.Println("Quit the application -> Press 'Q'\n" + Reset)
	}

}
func saveTask(){
	data, err := json.MarshalIndent(TasksSlice, "", "  ")
	if err != nil{
		print.Println(Red + "Error saving tasks!" + Reset)
		return
	}
err = os.WriteFile("Tasks.json", data, 0644)
}
func loadTasks(){
	data, err := os.ReadFile("Tasks.json")
	if err != nil {
		return
	}
	json.Unmarshal(data, &TasksSlice)
}
func showMenu() {
	print.Println(Yellow + "---Task Manager---\n" + Reset)

	if len(TasksSlice) > 0 {
		start := 0
		if len(TasksSlice) > 5 {
			start = len(TasksSlice) - 5
		}
		for i := start; i < len(TasksSlice); i++ {
			print.Printf(Yellow+"%d. %s\n", TasksSlice[i].ID, TasksSlice[i].Title)
		}
	} else {
		print.Println(Red + "No Tasks Yet!\n" + Reset)
	}
	hr()
	menuOptions()
}
func addTask() {
	print.Println(Yellow + "\n---Task Manager | Add a Task---\n")
	for {
		reader := bufio.NewReader(os.Stdin)
		// Ask Title
		print.Print(Blue + "Add your Task Title: ")
		taskTitle, _ := reader.ReadString('\n')
		taskTitle = strings.ToUpper(strings.TrimSpace(taskTitle))
		// Ask Assigned To
		print.Print(Blue + "Add your Task Assigned To: ")
		taskAssigned, _ := reader.ReadString('\n')
		taskAssigned = strings.ToUpper(strings.TrimSpace(taskAssigned))
		// Ask Due Date
		print.Print(Blue + "When is the Due Date (Format: Jan 02, 2006): ")
		taskDueDate, _ := reader.ReadString('\n')
		taskDueDate = strings.TrimSpace(taskDueDate)
		// Check for Error in Due Date Parsing
		dueDate, err := time.Parse("Jan 02, 2006", taskDueDate)
		if err != nil {
			print.Println(Red + "Wrong format! Try again.")
			time.Sleep(700 * time.Millisecond)
			clearScreen()
			print.Println(Yellow + "\n---Task Manager / Add a Task---\n")
			continue
		}
		// Generate new ID, and Append Task to TasksSlice
		newID := len(TasksSlice) + 1
		newTask := Task{
			ID:          newID,
			Title:       taskTitle,
			AssignedTo:  taskAssigned,
			Status:      false,
			DueDate:     dueDate,
			TimeCreated: time.Now(),
		}
		TasksSlice = append(TasksSlice, newTask)
		saveTask()
		print.Println(Reset + "\nTask added successfully!\n")
		print.Printf(Yellow+"ID: %v \nTitle: %v\nAssigned To: %v\nDue Date: %v\nTime Created %v\n", newTask.ID, newTask.Title, newTask.AssignedTo, newTask.DueDate.Format("Jan 02, 2006"), newTask.TimeCreated.Format("Jan 02, 2006"))
		// Ask if they wanna add another task
		for {
			print.Print(Blue + "\nDo you want to add another task? (Y/N): " + Reset)
			taskAgain, _ := reader.ReadString('\n')
			taskAgain = strings.ToUpper(strings.TrimSpace(taskAgain))

			if taskAgain == "Y" {
				clearScreen()
				break
			} else if taskAgain == "N" {
				return
			} else {
				print.Println(Red + "Respond with Y or N only!" + Reset)
			}
		}
	}
}
func pressAnyButton(){
	print.Println(Green + "Press any button to return to the Main Screen..\n" + Reset)
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
}
func showAllTask() {
	print.Println(Yellow + "---Task Manager | All Tasks---\n" + Reset)
	if len(TasksSlice) > 0 {
		for _, task := range TasksSlice {
			print.Printf(Yellow+"ID: %v, Title: %v, Assigned To: %v, Due Date: %v\n\n", task.ID, task.Title, task.AssignedTo, task.DueDate.Format("Jan 02, 2006"))
		}
		pressAnyButton()
		clearScreen()
	} else {
		print.Println(Red + "No Tasks Yet!\n" + Reset)
		pressAnyButton()
		clearScreen()
	}
}
func deleteTask() {
	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()
		print.Println(Yellow + "\n---Task Manager | Delete a Task---\n" + Reset)

		// Show tasks fresh each loop
		for _, task := range TasksSlice {
			print.Printf(Yellow+"ID: %v, Title: %v, Assigned To: %v, Due Date: %v\n\n",
				task.ID, task.Title, task.AssignedTo, task.DueDate.Format("Jan 02, 2006"))
		}

		print.Print(Blue + "Enter the ID to delete (or 'Q' to go back): " + Reset)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Let them quit
		if strings.ToUpper(input) == "Q" {
			return
		}

		// Check if it's a number
		id, err := strconv.Atoi(input)
		if err != nil {
			print.Println(Red + "Invalid! Enter a number." + Reset)
			time.Sleep(1 * time.Second)
			continue // loop again
		}

		// Find the task
		indexToDelete := -1
		for i, task := range TasksSlice {
			if task.ID == id {
				indexToDelete = i
				break
			}
		}

		if indexToDelete == -1 {
			print.Println(Red + "Task not found!" + Reset)
			time.Sleep(1 * time.Second)
			continue // loop again
		}

		// Delete it
		TasksSlice = append(TasksSlice[:indexToDelete], TasksSlice[indexToDelete+1:]...)
		saveTask()
		print.Println(Green + "\nTask deleted!" + Reset)
		time.Sleep(1 * time.Second)

		// If no more tasks, exit to menu
		if len(TasksSlice) == 0 {
			return
		}
		// else loop continues, shows remaining tasks
	}
}
// Application Entry Point
func main() {
	loadTasks()
	for {
		showMenu()
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.ToUpper(strings.TrimSpace(choice))
		switch choice {
		case "A":
			clearScreen()
			addTask()
			clearScreen()
		case "V":
			clearScreen()
			showAllTask()
		case "Q":
			clearScreen()
			print.Println(Blue+"Good bye!")
			time.Sleep(700 * time.Millisecond)
			clearScreen()
			print.Println(Yellow+"Good bye!!")
			time.Sleep(700 * time.Millisecond)
			clearScreen()
			print.Println(Red+"Good bye!!!")
			time.Sleep(700 * time.Millisecond)
			clearScreen()
			return
		case "D":
			if len(TasksSlice) > 0 {
				deleteTask()
			}else{
				clearScreen()
			print.Println("Wrong Input.")
			time.Sleep(400 * time.Millisecond)
			clearScreen()
			print.Println("Wrong Input..")
			time.Sleep(400 * time.Millisecond)
			clearScreen()
			print.Println("Wrong Input...")
			time.Sleep(400 * time.Millisecond)
			clearScreen()
			}
		default:
			clearScreen()
			print.Println("Wrong Input.")
			time.Sleep(400 * time.Millisecond)
			clearScreen()
			print.Println("Wrong Input..")
			time.Sleep(400 * time.Millisecond)
			clearScreen()
			print.Println("Wrong Input...")
			time.Sleep(400 * time.Millisecond)
			clearScreen()
		}

	}
}
