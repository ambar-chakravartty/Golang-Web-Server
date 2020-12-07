package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Server running at port 9000")
	http.HandleFunc("/",handle)
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static/"))))
	http.ListenAndServe(":9000",nil)
}

func handle(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello world")
}