
package main  
import (  
    "fmt"
    "sort"
)  
type task struct {  
    TaskName string
    Priority int
}
type byPriority []task  
func (a byPriority) Len() int {  
    return len(a)  
}  
func (a byPriority) Swap(i, j int) {  
    a[i], a[j] = a[j], a[i]  
}  
func (a byPriority) Less(i, j int) bool {  
    return a[i].Priority > a[j].Priority  
}

func insertTask(tasks *[]task, newTask *task){
    //add the new task to the end of the slice
    *tasks = append(*tasks, *newTask)
    
    //Sort the slice in descending order based on the Priority
    sort.Sort(byPriority(*tasks))
}

func main() {
    tasks := []task{
        {TaskName: "task1", Priority: 5},
        {TaskName: "task2", Priority: 3},
        {TaskName: "task3", Priority: 2},
    }
    
    fmt.Println("Original tasks:")
    fmt.Println(tasks)
    
    //New task with priority 4 to be inserted
    newTask := task{TaskName: "task4", Priority: 4}
    insertTask(&tasks, &newTask)
    
    fmt.Println("\nTasks after inserting task4:")
    fmt.Println(tasks)
    
     //New task with priority 1 to be inserted
     newTask = task{TaskName: "task5", Priority: 1}
     insertTask(&tasks, &newTask)
     
     fmt.Println("\nTasks after inserting task5:")
     fmt.Println(tasks)
} 
