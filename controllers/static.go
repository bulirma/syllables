package controllers

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func ServeStaticFile(path string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile(fmt.Sprintf("./static/%s", path))
		if err != nil {
			http.Redirect(w, r, "/", http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, string(content))
	}
}
