package controller

import (
	"blog-service-v3/internal/controller/authentication"
	"blog-service-v3/pkg/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	Title string   `json:"title"`
	Text  string   `json:"text"`
	Cats  []string `json:"cats"`
}

type PageSize struct {
	Size int `json:"size"`
}

func (h *postApi) createNewPost(w http.ResponseWriter, r *http.Request) {
	if !authentication.CheckAuthentication(r) {
		fmt.Fprintf(w, "Not allowed!")
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	var post Post
	err := json.Unmarshal(reqBody, &post)
	if err != nil {
		fmt.Fprintf(w, "Bad request!")
		return
	}

	post.Cats = lib.DeDuplicate(post.Cats)
	if len(post.Cats) > 6 {
		fmt.Fprintf(w, "More than 6 cats is not allowed!")
		return
	}

	h.app.CreatePost(post.Title, post.Text, post.Cats)
}

func (h *postApi) allPosts(w http.ResponseWriter, r *http.Request) {
	allPosts, err := h.app.AllPosts()
	if err != nil {
		fmt.Fprintf(w, "Something went wrong!")
		return
	}

	json.NewEncoder(w).Encode(allPosts)
}

func (h *postApi) pagePosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["id"])
	if page < 1 {
		fmt.Fprintf(w, "paging starts at 1")
		return
	}

	var pageSize PageSize
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &pageSize)

	posts, err := h.app.PagePosts(page, pageSize.Size)
	if err != nil {
		fmt.Fprintf(w, "Something went wrong!")
		return
	}

	json.NewEncoder(w).Encode(posts)
}

func (h *postApi) updatePost(w http.ResponseWriter, r *http.Request) {
	if !authentication.CheckAuthentication(r) {
		fmt.Fprintf(w, "Not allowed!")
		return
	}

	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	reqBody, _ := ioutil.ReadAll(r.Body)

	var post Post
	err := json.Unmarshal(reqBody, &post)
	if err != nil {
		fmt.Fprintf(w, "Bad request!")
		return
	}

	post.Cats = lib.DeDuplicate(post.Cats)
	if len(post.Cats) > 6 {
		fmt.Fprintf(w, "More than 6 cats is not allowed!")
		return
	}

	h.app.UpdatePost(key, post.Title, post.Text, post.Cats)
}

func (h *postApi) deletePost(w http.ResponseWriter, r *http.Request) {
	if !authentication.CheckAuthentication(r) {
		fmt.Fprintf(w, "Not allowed!")
		return
	}

	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	h.app.DeletePost(key)
}
