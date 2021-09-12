package main

import (
	"net/http"
	"encoding/json"
)

type Post struct {
	Id int 			`json:id`
	Title string	`json:title`
	Text string 	`json:text`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{Id:1, Title:"Title 1", Text:"Text 1"}}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	res, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the posts list"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	w.WriteHeader(http.StatusOK)
	res, err := json.Marshal(posts)
	w.Write(res)
}