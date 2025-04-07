package web_dev

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've hit %s\n", r.URL.Path)
}

func Http() {
    http.HandleFunc("/", handler) 

    fmt.Println("Starting server at http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
