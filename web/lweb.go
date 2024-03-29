package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm();
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println(r.Form["url_long"], r.Form.Get("url_long"))
	fmt.Println(r.Form["none_arg"], r.Form.Get("none_arg"))
	fmt.Println(r.Form["test"], r.Form.Get("test"))

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Fprintf(w, strings.Join(append(r.Form["username"], r.Form["password"]...), "-"));
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}