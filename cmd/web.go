package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/paked/down"
	"github.com/paked/models"
	"gopkg.in/mgo.v2/bson"
)

var (
	homePage []byte
)

func init() {
	var err error

	rand.Seed(time.Now().UnixNano())

	homePage, err = ioutil.ReadFile("templates/home.html")
	if err != nil {
		panic(err)
	}
}

func main() {
	models.Init("localhost", "downline")

	r := mux.NewRouter()

	r.HandleFunc("/", getEditorHandler).Methods("GET")
	r.HandleFunc("/new_content", postRegisterContentHandler).Methods("POST")
	r.HandleFunc("/view/{key}", viewContentHandler)

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

func getEditorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<html>
	<body> 
		<form action="/new_content" method="POST">
			<textarea name="content" id="content" cols="30" rows="10"></textarea>
			<br>
			<input type="submit" />
		</form>
	</body>
</html>`)
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
