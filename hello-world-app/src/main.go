package main

import (
    "fmt"
    "net/http"
	"os"
	"github.com/gorilla/handlers"
)


func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}


func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Hello " + getEnv("NAME", "World"))
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/", hello)
    http.HandleFunc("/headers", headers)
    http.ListenAndServe(":" + getEnv("PORT","8080") , handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}