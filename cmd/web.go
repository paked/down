package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/paked/down"
	"github.com/paked/models"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	var err error

	rand.Seed(time.Now().UnixNano())

	url := os.Getenv("DOWN_MONGODB_URL")
	if url == "" {
		url = "localhost"
	}
	fmt.Println("DB url is: ", url)

	err = models.Init(url, "downbase")
	if err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/new_content", postRegisterContentHandler).Methods("POST")
	r.HandleFunc("/view/{key}", viewContentHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static/")))

	http.Handle("/", r)

	fmt.Println("Listening on :8080...")
	fmt.Println(http.ListenAndServe(":8080", nil))
}

type Content struct {
	ID    bson.ObjectId `bson:"_id"`
	Down  string        `bson:"down"`
	Title string        `bson:"title"`
}

func (c Content) BID() bson.ObjectId {
	return c.ID
}

func (c Content) C() string {
	return "contents"
}

type Page struct {
	Content template.HTML
	Title   string
}

func postRegisterContentHandler(w http.ResponseWriter, r *http.Request) {
	content := Content{bson.NewObjectId(), r.FormValue("content"), r.FormValue("title")}

	models.Persist(content)

	http.Redirect(w, r, fmt.Sprintf("/view/%x", string(content.ID)), http.StatusFound)
}

func viewContentHandler(w http.ResponseWriter, r *http.Request) {
	var c Content
	id := mux.Vars(r)["key"]

	models.RestoreByID(&c, bson.ObjectIdHex(id))

	fmt.Println(c.Down)
	t, err := template.ParseFiles("templates/view.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	t.Execute(w, Page{template.HTML(down.Parse(c.Down)), c.Title})
}
