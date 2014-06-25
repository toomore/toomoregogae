package hello

import "fmt"
import "net/http"

func init() {
    http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World!")
}
