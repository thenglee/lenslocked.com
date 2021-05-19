package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>\n")
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:supprt@lenslocked.com\">support@lenslocked.com</a>.")
}

func main() {
	router := httprouter.New()
	router.GET("/", home)
	router.GET("/contact", contact)
	http.ListenAndServe(":3000", router)
}
