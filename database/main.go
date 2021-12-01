package main
import (
 "database/sql"
 "fmt"
 _ "github.com/go-sql-driver/mysql"
 "github.com/gorilla/mux"
 "log"
 "net/http"
 "html/template"
 
)

const (
	DBHost = "127.0.0.1"
	DBPort = ":3306"
	DBUser = "root"
	DBPass = "Start7"
	DBDbase = "test"
	PORT = ":8080"
   )

   var database *sql.DB
type Page struct {
 Title string
 RawContent string
 Content template.HTML
 Date string
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
 thisPage := Page{}
 fmt.Println(pageGUID)
 err := database.QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE page_guid=?",
 pageGUID).Scan(&thisPage.Title, &thisPage.RawContent,
 &thisPage.Date)
  thisPage.Content = template.HTML(thisPage.RawContent)
 if err != nil {
 log.Println("Couldn't get page: +pageID")
 http.Error(w, http.StatusText(404), http.StatusNotFound)
 log.Println(err.Error)
 return
 }
 t, _ := template.ParseFiles("templates/blog.html")
 t.Execute(w, thisPage)
 
}

func main() {
	dbConn := fmt.Sprintf("%s:%s@/%s", DBUser, DBPass, DBDbase)
	fmt.Println(dbConn)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
	log.Println("Couldn't connect to"+DBDbase)
	log.Println(err.Error)
	}
	database = db
	routes := mux.NewRouter()
	routes.HandleFunc("/page/{guid:[0-9a-zA\\-]+}", ServePage)
	http.Handle("/", routes)
	http.ListenAndServe(PORT, nil)
   }
