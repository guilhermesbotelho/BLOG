package controllers

import (
	"net/http"

	"github.com/guilhermesbotelho/BLOG/models"
)

func Likes(w http.ResponseWriter, r *http.Request) {
	idPost := r.URL.Query().Get("id")
	amountLikes := models.CountLikes(idPost)
	temp.ExecuteTemplate(w, "countLikes", amountLikes)

	post := models.GetAllLikes(idPost)
	temp.ExecuteTemplate(w, "Likes", post)
}
