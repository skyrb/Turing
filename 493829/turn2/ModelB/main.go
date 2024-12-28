package main  
import "fmt"  
type Result interface {  
    GetValue() int  
    GetError() error  
    GetMetadata() map[string]interface{}  
}  
type SuccessfulResult struct {  
    value    int  
    metadata map[string]interface{}  
}  
func (r *SuccessfulResult) GetValue() int {  
    return r.value  
}  
func (r *SuccessfulResult) GetError() error {  
    return nil  
}  
func (r *SuccessfulResult) GetMetadata() map[string]interface{} {  
    return r.metadata  
}  
type FailedResult struct {  
    error  
    metadata map[string]interface{}  
}  
func (r *FailedResult) GetValue() int {  
    return 0  
}  
func (r *FailedResult) GetError() error {  
    return r.error  
}  
func (r *FailedResult) GetMetadata() map[string]interface{} {  
    return r.metadata  
}  
type CustomError struct {  
    Code    int    `json:"code"`  
    Message string `json:"message"`  
}  
func (e *CustomError) Error() string {  
    return fmt.Sprintf("Custom Error: code=%d, message=%s", e.Code, e.Message)  
}  
func calculate(x, y int) Result {  
    metadata := map[string]interface{}{}  
    if y == 0 {  
        err := &CustomError{Code: 1001, Message: "division by zero"}  
        metadata["debug_info"] = "Additional debugging details"  
        return &FailedResult{err, metadata}  
    }  
    result := x / y  
    metadata["execution_time"] = "10ms"  
    return &SuccessfulResult{result, metadata}  
}  
func main() {  
    res := calculate(10, 2)  
    if err := res.GetError(); err != nil {  
        fmt.Println("Error:", err)  
        if customErr, ok := err.(*CustomError); ok {  
            fmt.Println("Custom Error Code:", customErr.Code)  
        }  
    } else {  
        fmt.Println("Result:", res.GetValue())  
        metadata := res.GetMetadata()  
        fmt.Println("Execution Time:", metadata["execution_time"])  
    }  
    res = calculate(10, 0)  
    if err := res.GetError(); err != nil {  
        fmt.Println("Error:", err)  
        fmt.Println("Debug Info:", res.GetMetadata()["debug_info"])  
    } else {  
        fmt.Println("Result:", res.GetValue())  
    }  
}  