package main

import (
	"fmt"
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
	ID   bson.ObjectId `bson:"_id"`
	Down string        `bson: "down"`
	Key  string        `bson: "key"`
}

func (c Content) BID() bson.ObjectId {
	return c.ID
}

func (c Content) C() string {
	return "contents"
}

func postRegisterContentHandler(w http.ResponseWriter, r *http.Request) {
	content := Content{bson.NewObjectId(), r.FormValue("content"), randomString(64)}
	models.Persist(content)

	http.Redirect(w, r, fmt.Sprintf("/view/%v", content.Key), http.StatusFound)
}

func viewContentHandler(w http.ResponseWriter, r *http.Request) {
	var c Content
	key := mux.Vars(r)["key"]

	models.Restore(&c, bson.M{"key": key})

	fmt.Println(c.Down)
	fmt.Fprintf(w, "<html><body>%v</body></html>", down.Parse(c.Down))
}

// Random key generation
func randomString(length int) string {
	bytes := []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var end []byte
	for i := 0; i < length; i++ {
		r := rand.Intn(len(bytes))
		end = append(end, bytes[r])
	}

	return string(end)
}
