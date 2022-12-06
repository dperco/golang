package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Users struct {
	Name   string
	Skills string
	Age    int
	img    string
}

func Index(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./templates/index.html")

	user := Users{"Ariel", "developer", 34, "https://t4.ftcdn.net/jpg/02/87/68/17/240_F_287681725_BduR6WZCsS4ivuBInHdqnCWOvaTpxxUz.jpg"}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, user)
	}
}

func main() {
	http.HandleFunc("/", Index)

	fmt.Println("Servidor corriendo en port 3000")
	http.ListenAndServe("localhost:3000", nil)
}
