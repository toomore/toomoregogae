package hello

import "fmt"
import "net/http"
import "appengine"
import "appengine/user"

func init() {
    http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    u := user.Current(c)

    if u == nil {
        url, err := user.LoginURL(c, r.URL.String())
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Location", url)
        w.WriteHeader(http.StatusFound)
        return
    }

    fmt.Fprintf(w, "Hello, %v", u)
}
