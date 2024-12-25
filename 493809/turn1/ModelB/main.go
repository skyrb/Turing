package main  
import (  
    "fmt"
    "log"
    "net/http"
	"os"
    "time"
)

type LoggingMiddleware struct {  
    Logger *log.Logger
}

func (m *LoggingMiddleware) Handler(next http.Handler) http.Handler {  
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {  
        start := time.Now()
        next.ServeHTTP(w, r)
        duration := time.Since(start)
        logMessage := fmt.Sprintf("%s %s %s %s", r.Method, r.RequestURI, w.Header().Get("Content-Type"), duration)
        m.Logger.Println(logMessage)
    })
}

func main() {  
    logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
    middleware := LoggingMiddleware{Logger: logger}
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {  
        fmt.Fprintf(w, "Hello, World!")
    })
    http.ListenAndServe(":8080", middleware.Handler(http.DefaultServeMux))
}