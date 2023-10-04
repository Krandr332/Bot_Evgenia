// route.go
package web

import (
	"html/template"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)

}
