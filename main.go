package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       string `json:"Id"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
type newusers struct {
	sync.Mutex
	store map[string]users
}
type newPosts struct{
	sync.Mutex
	store map[string]posts
}
type Post struct {
	Id               string `json:"Id"`
	Caption          string `json:"Id"`
	Image_URL        string `json:"Id"`
	Posted_Timestamp string `json:"Id"`
}
func (h *users) user(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}
func (g *posts) post(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		g.get(w, r)
		return
	case "POST":
		g.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

var Posts []Post
var Users []User
func (h *users) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	var user User
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
func (g *posts) post(w http.ResponseWriter, r *http.Request) 
{
		bodyBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	
		ct := r.Header.Get("content-type")
		if ct != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
			return
		}
	
		var post Post
		err = json.Unmarshal(bodyBytes, &post)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}


func newUsers() *users {
	return &users{
		store: map[string]User{},
	}

}
func newPosts() *posts {
	return &posts{
		store: map[string]Post{},
	}

}

func handleRequests() {
	users := newUsers()
	posts = newPosts()
	
	http.HandleFunc("/users", users.user)
	http.HandleFunc("/users/<id here>", guserid)
	http.HandleFunc("/posts", posts.post)
	http.HandleFunc("/posts/<id here>", postid)
	http.HandleFunc("/posts/users/<id here>", gpostid)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	Users = []User{}
	handleRequests()
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }
	collection := client.Database("db").Collection("users")
   users := []users{}

    insertManyResult, err := collection.InsertMany(context.TODO(), users)
    if err != nil {
      log.Fatal(err)
  }
    fmt.Println( insertManyResult.InsertedIDs)
}
