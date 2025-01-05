

package main  
import (  
    "errors"  
    "fmt"  
)  

// Function to calculate the area of a circle and return multiple values  
func calculateArea(radius float64) (float64, error) {  
    if radius < 0 {  
        return 0, errors.New("Radius cannot be negative")  
    }  
    area := 3.14 * radius * radius  
    return area, nil  
}  

func main() {  
    // Call the function with a positive radius  
    radius := 5.0  
    area, err := calculateArea(radius)  
    if err != nil {  
        fmt.Println("Error:", err)  
    } else {  
        fmt.Println("Area of the circle with radius", radius, "is:", area)  
    }  

    // Call the function with a negative radius  
    radius = -2.0  
    area, err = calculateArea(radius)  
    if err != nil {  
        fmt.Println("Error:", err)  
    } else {  
        fmt.Println("Area of the circle with radius", radius, "is:", area)  
    }  
}  

