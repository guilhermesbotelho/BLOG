package routes

import (
	"net/http"

	"github.com/guilhermesbotelho/BLOG/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/insertpost", controllers.InsertPost)
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/insert", controllers.NewPost)
	http.HandleFunc("/delet", controllers.Delet)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/commt", controllers.InsertCommt)
	http.HandleFunc("/insertcommt", controllers.NewCommt)
	http.HandleFunc("/seepost", controllers.SeePost)
	http.HandleFunc("/deletcommt", controllers.DeletCommt)
	http.HandleFunc("/likes", controllers.Likes)
}
