package main

import (
	"fmt"
	"net/http"
	"os"

	"go.uber.org/zap"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HTTPサーバ")
}

func main() {
	logger, _ := zap.NewProduction()
	fmt.Println(os.Args)

	//  http.HandleFunc("/", handler)
	//  http.ListenAndServe(":8080", nil)
}
