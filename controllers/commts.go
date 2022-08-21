package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/guilhermesbotelho/BLOG/models"
)

func NewCommt(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		username := r.FormValue("username")
		commt := r.FormValue("commt")

		idConvertToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID para int:", err)
		}

		comment := models.CreateNewCommt(username, commt, idConvertToInt)
		if comment != "" {
			temp.ExecuteTemplate(w, "error", comment)
			return
		}

	}

	id := r.FormValue("id")
	http.Redirect(w, r, "/seepost?id="+id, 301)
}

func InsertCommt(w http.ResponseWriter, r *http.Request) {
	idPost := r.URL.Query().Get("id")
	post := models.EditPost(idPost)
	temp.ExecuteTemplate(w, "InsertCommt", post)
}

func DeletCommt(w http.ResponseWriter, r *http.Request) {
	idPostDelet := r.URL.Query().Get("id")
	models.DeletComment(idPostDelet)

	http.Redirect(w, r, "/", 301)
}
