package main

import (
	"fmt"
	"net/http"
	"html/template"
)
var UserName string
func indexHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		http.ServeFile(w,r,"index.html")
	case "POST":
		UserName = r.FormValue("user_name")
		t, err := template.ParseFiles("home.html")
		if err != nil {
			fmt.Println (err)
		}else{
			if err := t.Execute(w,UserName); err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("POST")
	}
}

func main(){
	http.HandleFunc("/",indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	if err := http.ListenAndServe(":8086",nil); err != nil {
		fmt.Println(err)
	}

}