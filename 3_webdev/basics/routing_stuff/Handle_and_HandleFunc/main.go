package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

// USING Handler
type index int

func (i index) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "homepage")
}

type dog int

func (d dog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog page")
}

type me int

func (m me) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(res, "index.gohtml", "Sergey")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {
	var idx index
	var d dog
	var m me

	http.Handle("/", idx)
	http.Handle("/dog", d)
	http.Handle("/me", m)

	http.ListenAndServe(":8080", nil)
}

////// USING HandlerFunc
// func main() {
// 	http.HandleFunc("/", foo)
// 	http.HandleFunc("/dog/", bar)
// 	http.HandleFunc("/me/", myName)
// 	http.ListenAndServe(":8080", nil)
// }

////// OR
// func main() {
// 	http.Handle("/", http.HandlerFunc(foo))
// 	http.Handle("/dog/", http.HandlerFunc(bar))
// 	http.Handle("/tm/", http.HandlerFunc(mcleod))
// 	http.ListenAndServe(":8080", nil)
// }

// func foo(w http.ResponseWriter, req *http.Request) {
// 	io.WriteString(w, "foo ran")
// }

// func bar(w http.ResponseWriter, req *http.Request) {
// 	io.WriteString(w, "bar ran")
// }

// func myName(w http.ResponseWriter, req *http.Request) {
// 	io.WriteString(w, "hello lala")
// }
