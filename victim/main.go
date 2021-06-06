package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/google/uuid"
)

type Post struct {
	Name    string
	Content string
}

var users = make(map[string]string) // key: SID, value: name
var posts []Post = []Post{{"HIRATA", "Hello, World."}}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}
	_, err := r.Cookie("sid")
	if err != nil {
		t, err := template.ParseFiles("login.html")
		if err != nil {
			panic(err.Error())
		}
		if err := t.Execute(w, nil); err != nil {
			panic(err.Error())
		}
		return
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, posts); err != nil {
		panic(err.Error())
	}
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}
	uuidObj, _ := uuid.NewUUID()
	SID := uuidObj.String()
	users[SID] = r.PostForm.Get("name")
	cookie := &http.Cookie{
		Name:  "sid",
		Value: SID,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "http://localhost:8080/", http.StatusMovedPermanently)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}
	cookie, _ := r.Cookie("sid")
	posts = append(posts, Post{users[cookie.Value], r.PostForm.Get("content")})
	fmt.Println(posts)
	http.Redirect(w, r, "http://localhost:8080/", http.StatusMovedPermanently)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	posts = []Post{}
	http.Redirect(w, r, "http://localhost:8080/", http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/signup", signinHandler)
	http.HandleFunc("/send", postHandler)
	http.HandleFunc("/reset", resetHandler)
	http.ListenAndServe(":8080", nil)
}
