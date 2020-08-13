package main

import (
	"fmt"
	"html/template"
	"net/http"
	"math/rand"
	"time"
)
type selectedClr struct{
	Clr1,Clr2,Clr3,Clr4,Clr5,Clr6 string
}
var randmPic selectedClr


func main (){

http.HandleFunc("/",indexHandler)
http.ListenAndServe(":8086",nil)

}
func indexHandler( w http.ResponseWriter, r *http.Request){
	randmPic.Clr1,randmPic.Clr2,randmPic.Clr3,randmPic.Clr4,randmPic.Clr5,randmPic.Clr6 = picSelecter()
	t,error := template.ParseFiles("./template/index.html")
	if error != nil {
		fmt.Println(error)
		fmt.Println("ERROR")
	}else {
		if err := t.Execute(w,randmPic); err != nil {
			fmt.Println(err)
		}

	}
}


func picSelecter()(link1,link2,link3,link4,link5,link6 string){
clrList := []string{
		"red",
		"blue",
		"yellow",
		"orange",
		"aqua",
		"chartreuse",
		"chocolate",
		"coronsilk",
		"lightcyan",
		"darkorange",
		"darksalmon",
		"deeppink",
		"gold",
	}
	rand.Seed(time.Now().UnixNano())
	
	link1 = clrList[rand.Intn(len(clrList))]
	link2 = clrList[rand.Intn(len(clrList))]
	link3 = clrList[rand.Intn(len(clrList))]
	link4 = clrList[rand.Intn(len(clrList))]
	link5 = clrList[rand.Intn(len(clrList))]
	link6 = clrList[rand.Intn(len(clrList))]
	return
}