package main  
import (  
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func displayTasks(tasks []string) {
  fmt.Println("Your to-do list:")
  for i, task := range tasks {
    fmt.Printf("%d. %s\n", i+1, task)
  }
}

func addTask(tasks *[]string) {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter task to add: ")
  task, _ := reader.ReadString('\n')
  *tasks = append(*tasks, task[:len(task)-1])
}
func deleteTask(tasks *[]string) {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter the task number to delete: ")
  num, _ := reader.ReadString('\n')
  num = num[:len(num)-1]
  taskNumber := parseInt(num)
  if taskNumber > 0 && taskNumber <= len(*tasks) {
    *tasks = append((*tasks)[:taskNumber-1], (*tasks)[taskNumber:]...)
  } else {
    fmt.Println("Invalid task number.")
  }
}

func parseInt(num string) int {
  n, err := strconv.Atoi(num)
  if err != nil {
    fmt.Println("Invalid input.")
    return -1
  }
  return n
}

func main() {  
  tasks := []string{}  
  for {
    fmt.Println("\nMenu:")
    fmt.Println("1. Display tasks")
    fmt.Println("2. Add task")
    fmt.Println("3. Delete task")
    fmt.Println("4. Exit")
    fmt.Print("Enter your choice: ")

    reader := bufio.NewReader(os.Stdin)
    choice, _ := reader.ReadString('\n')
    choice = choice[:len(choice)-1]
    ch := parseInt(choice)

    switch ch {
    case 1:
      displayTasks(tasks)
    case 2:
      addTask(&tasks)
    case 3:
      deleteTask(&tasks)    
    case 4:
      fmt.Println("Goodbye!")
      os.Exit(0)
    default:
      fmt.Println("Invalid choice. Please try again.")
    }
  }
}
  