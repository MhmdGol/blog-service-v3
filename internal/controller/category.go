package controller

import (
	"blog-service-v3/internal/controller/authentication"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Category struct {
	Name string `json:"name"`
}

func (h *categoryApi) createNewCategory(w http.ResponseWriter, r *http.Request) {
	if !authentication.CheckAuthentication(r) {
		fmt.Fprintf(w, "Not allowed!")
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	var category Category
	err := json.Unmarshal(reqBody, &category)
	if err != nil {
		fmt.Fprintf(w, "Not allowed!")
		return
	}

	h.app.CreateCategory(category.Name)
}

func (h *categoryApi) allCategories(w http.ResponseWriter, r *http.Request) {
	allCategories, err := h.app.AllCategories()
	if err != nil {
		fmt.Fprintf(w, "Something went wrong!")
		return
	}

	json.NewEncoder(w).Encode(allCategories)
}

func (h *categoryApi) updateCategory(w http.ResponseWriter, r *http.Request) {
	if !authentication.CheckAuthentication(r) {
		fmt.Fprintf(w, "Not allowed!")
		return
	}

	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	reqBody, _ := ioutil.ReadAll(r.Body)

	var category Category
	err := json.Unmarshal(reqBody, &category)
	if err != nil {
		fmt.Fprintf(w, "Bad request!")
		return
	}

	h.app.UpdateCategory(key, category.Name)
}

func (h *categoryApi) deleteCategory(w http.ResponseWriter, r *http.Request) {
	if !authentication.CheckAuthentication(r) {
		fmt.Fprintf(w, "Not allowed!")
		return
	}

	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	h.app.DeleteCategory(key)
}
