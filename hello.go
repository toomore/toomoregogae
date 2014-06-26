package hello

import "fmt"
import "net/http"
import "html/template"
import "appengine"
import "appengine/urlfetch"


func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/sign", sign)
    http.HandleFunc("/requests", requests)
}

func root(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, guestbookForm)
}

const guestbookForm = `
<html>
  <body>
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`

func sign(w http.ResponseWriter, r *http.Request) {
    err := signTemplate.Execute(w, r.FormValue("content"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

var signTemplate = template.Must(template.New("sign").Parse(signTemplateHTML))

const signTemplateHTML = `
<html>
  <body>
    <p>You wrote:</p>
    <pre>{{.}}</pre>
  </body>
</html>
`

func requests(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    client := urlfetch.Client(c)
    resp, err := client.Get("http://www.google.com/")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "HTTP GET returned status %v", resp.Status)
}
