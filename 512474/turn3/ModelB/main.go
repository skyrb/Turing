package main  
import (  
    "fmt"
    "sort"
)  
type Student struct {  
    Name  string
    Class string
    Score int
}  
func main() {  
    students := []Student{  
        {"Alice", "Class A", 85},  
        {"Charlie", "Class B", 90},  
        {"Bob", "Class A", 75},  
        {"David", "Class B", 80},  
        {"Eve", "Class A", 95},  
    }  
    // Sort the students by class, then by scores within each class
    sort.Slice(students, func(i, j int) bool {
        if students[i].Class != students[j].Class {
            return students[i].Class < students[j].Class
        }
        return students[i].Score > students[j].Score
    })  
    var output string
    for i := 0; i < len(students); {
        class := students[i].Class
        j := i + 1
        for ; j < len(students) && students[j].Class == class; j++ {
        }
        studentsInClass := students[i:j]
        sort.Slice(studentsInClass, func(i, j int) bool {
            return studentsInClass[i].Score > studentsInClass[j].Score
        })
        classOutput := fmt.Sprintf("%s: [", class)
        for _, student := range studentsInClass {
            classOutput += fmt.Sprintf("%s - %d, ", student.Name, student.Score)
        }
        classOutput = classOutput[:len(classOutput)-2] + "]"
        output += classOutput + "\n"
        i = j
    }
    fmt.Println(output)
}