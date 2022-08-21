package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/guilhermesbotelho/BLOG/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allPosts := models.GetAllPosts()
	temp.ExecuteTemplate(w, "Index", allPosts)
}

func InsertPost(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "InsertPost", nil)
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		blogger := r.FormValue("blogger")
		post := r.FormValue("post")

		models.CreateNewPost(blogger, post)

	}

	http.Redirect(w, r, "/", 301)
}

func Delet(w http.ResponseWriter, r *http.Request) {
	idPostDelet := r.URL.Query().Get("id")
	models.DeletPost(idPostDelet)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idPost := r.URL.Query().Get("id")
	post := models.EditPost(idPost)
	temp.ExecuteTemplate(w, "Edit", post)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		blogger := r.FormValue("blogger")
		post := r.FormValue("post")

		idConvertToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID para int:", err)
		}

		models.UpdatePost(idConvertToInt, post, blogger)
	}

	id := r.FormValue("id")
	http.Redirect(w, r, "/seepost?id="+id, 301)
}

func SeePost(w http.ResponseWriter, r *http.Request) {
	idPost := r.URL.Query().Get("id")
	post := models.EditPost(idPost)
	temp.ExecuteTemplate(w, "SeePost", post)

	allCommts := models.GetAllCommts(idPost)
	temp.ExecuteTemplate(w, "_commt", allCommts)

}
