package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"net/http"
	/*Random numbers*/

	/*Database*/

	_ "github.com/mattn/go-sqlite3"
)

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	inputName := r.FormValue("inputName")
	textarea := r.FormValue("textarea")

	// database, err := sql.Open("sqlite3", "./posts.db")
	// isError(err)
	//
	// stmt, err := database.Prepare("INSERT INTO post(uid, username, message, postDate) values(?,?,?,?)")
	// stmt.Exec(RandomID(), inputName, textarea, time.Now())
	// isError(err)
	fmt.Println("Username:"+inputName, "\nMessage:"+textarea)

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
func RandomID() string {
	b := make([]byte, 16)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}

// data := map[string]string{
// 		"id": RandomId(),
// 		"username": inputName,
// 		"message": textarea,
// }
// json, e := json.MarshalIndent(data, "", "")
// if e != nil {
// 	log.Fatal(e)
// }
