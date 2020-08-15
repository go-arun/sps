package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/sessions"
)

var (
    // key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
    key = []byte("super-secret-key")
    store = sessions.NewCookieStore(key)
    sessionLive bool
)

func home(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	fmt.Println("HOMEEEEEE")

    // Check if user is authenticated
        auth, ok := session.Values["authenticated"].(bool)
    if !ok || !auth || !sessionLive {
		// http.Error(w, "Forbidden", http.StatusForbidden)
		http.Redirect(w, r, "/login", 301)
        return
    }
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	w.Header().Set("Pragma", "no-cache") // HTTP 1.0.
	w.Header().Set("Expires", "0") // Proxies.
	// Print secret message
	http.ServeFile(w,r,"home.html")
    
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//http.ServeFile(w,r,"home.html")
		sessionLive = true
		http.Redirect(w, r, "/home", 301)
	}else{
    	sessionLive = true
    	session, _ := store.Get(r, "cookie-name")

    	// Authentication goes here
    	// Set user as authenticated
    	session.Values["authenticated"] = true
    	if err := session.Save(r, w); err != nil {

        	fmt.Println("Erro",err)
		}
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
		w.Header().Set("Pragma", "no-cache") // HTTP 1.0.
		w.Header().Set("Expires", "0") // Proxies.
		http.ServeFile(w,r,"login.html")
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
    sessionLive = false
    session, _ := store.Get(r, "cookie-name")

    // Revoke users authentication
    session.Values["authenticated"] = false
        //session.Save(r, w)
        session.Options.MaxAge = -1
	session.Save(r, w)
	http.ServeFile(w,r,"logout.html")
}

func main() {
    http.HandleFunc("/home", home)
    http.HandleFunc("/", login)
    http.HandleFunc("/logout", logout)

    http.ListenAndServe(":8089", nil)
}








// package main

// import (
// 	"fmt"
// 	"database/sql"
// 	_"github.com/go-sql-driver/mysql"
// 	"http/net"
// )
// func main(){
// 	http.HandleFunc("/signin",signin)
// 	http.HandleFunc("/signup",signup)
// 	initDB()

// 	log.Fatal(http.ListenAndServe(":8086",nil))




// }

// func initDB(){
// 	db,err := sql.Open("mysql","root:uit@123@tcp(127.0.0.1:3306)/test")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("Connected!!")
// 	defer db.Close()
// }