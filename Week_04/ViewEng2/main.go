package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type news struct{
	RightHeadLine,RightDetails,LeftHeadLine,LeftDetails string
}
var newLeftnews news
var newRightnews news
func indexHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "index.html")
	case "POST":
		if whichSide := r.FormValue("which_side"); "right" == whichSide {
			newNews.RightHeadLine = r.FormValue("news_head")
			newNews.RightDetails = r.FormValue("news_head1_details")
		}else {
			newNews.LeftHeadLine = r.FormValue("news_head")
			newNews.LeftDetails = r.FormValue("news_head1_details")			

		}
		t,error := template.ParseFiles("./template/index.html")
		if error != nil {
			fmt.Println(error)
			fmt.Println("ERROR")
		}else {
			if err := t.Execute(w,newNews); err != nil {
				fmt.Println(err)
			}
		}
	}
}
func newsUpdater(w http.ResponseWriter, r *http.Request){
	fmt.Println("coming ",r.Method)
	switch r.Method {
	case "POST":
		http.ServeFile(w, r, "newsupdater.html")
	}
}
var newNews news
func main(){

	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/newsupdater",newsUpdater)

	fmt.Println("Starting Server")
	if err := http.ListenAndServe(":8086",nil); err != nil {
		fmt.Println("Some errr. handle later!!")
	}

}