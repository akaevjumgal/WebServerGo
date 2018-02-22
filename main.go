package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	/*Random numbers*/
	"crypto/rand"
	/*Database*/
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	createDB      = "create table if not exists post (uid integer, username text, message text, postDate CURRENT_TIMESTAMP)"
	database, err = sql.Open("sqlite3", "/home/akaev_jumgal/golang/WebServerGo/posts.db")
)

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	inputName := r.FormValue("inputName")
	textarea := r.FormValue("textarea")
	isError(err)

	database.Exec(createDB)
	tx, _ := database.Begin()
	stmt, err := tx.Prepare("INSERT INTO post(uid, username, message, postDate) values(?, ?, ?, ?)")
	isError(err)
	rnd := RandomID()
	stmt.Exec(rnd, inputName, textarea, time.Now())

	tx.Commit()

}

///home/akaev_jumgal/.cache/go-build
func main() {
	// writeFile()

	fmt.Println("Serving on port 8000...")
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/write", writePage)
	http.HandleFunc("/savePost", savePostHandler)
	http.ListenAndServe(":8000", nil)
}

/*Main handler for index page*/
func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

/*Secondary page - write, for sending some messages*/
func writePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html", "templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "write", nil)
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

//RandomID() is the function, that generate random user ID
func RandomID() string {

	b := make([]byte, 16)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}

// data := map[string]string{
// 	"id":
// 	"username": inputName,
// 	"message":  textarea,
// }
//
// json, e := json.MarshalIndent(stmt, "", "")
// if e != nil {
// 	log.Fatal(e)
// }
//
// fmt.Println(string(json))
