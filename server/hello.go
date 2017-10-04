package hello

import (
    "fmt"
    "net/http"
    "html/template"

    "appengine"
    "appengine/user"
)

var (
    indexTemplate = template.Must(template.ParseFiles("templates/index.html"))
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "text/html; charset=utf-8")
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

    url, err := user.LogoutURL(c, "/")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, `Hello! Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
    indexTemplate.Execute(w, nil)    
}

const guestbookForm = `
<html>
  <body>
  	Hello!
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`