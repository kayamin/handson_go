package main
import (
    "os"
    "fmt"
    "net/http"

    "go.uber.org/zap"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, HTTPサーバ")
}

func main(){
 logger, _ := zap.NewProduction()
 logger.Warn("warning test")
 fmt.Println(os.Args)


 http.HandleFunc("/", handler)
 http.ListenAndServe(":8080", nil)
}
