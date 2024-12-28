package main  
import "fmt"  
type Result interface {  
    GetValue() int  
    GetError() error  
}  
type SuccessfulResult struct {  
    value int  
}  
func (r *SuccessfulResult) GetValue() int {  
    return r.value  
}  
func (r *SuccessfulResult) GetError() error {  
    return nil  
}  
type FailedResult struct {  
    error  
}  
func (r *FailedResult) GetValue() int {  
    return 0  
}  
func (r *FailedResult) GetError() error {  
    return r.error  
}  
func calculate(x, y int) Result {  
    if y == 0 {  
        return &FailedResult{fmt.Errorf("division by zero")}  
    }  
    result := x / y  
    return &SuccessfulResult{result}  
}  
func main() {  
    res := calculate(10, 2)  
    if err := res.GetError(); err != nil {  
        fmt.Println("Error:", err)  
    } else {  
        fmt.Println("Result:", res.GetValue())  
    }  
    res = calculate(10, 0)  
    if err := res.GetError(); err != nil {  
        fmt.Println("Error:", err)  
    } else {  
        fmt.Println("Result:", res.GetValue())  
    }  
}