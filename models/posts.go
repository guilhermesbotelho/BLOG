package models

import (
	"github.com/guilhermesbotelho/BLOG/db"
)

type Post struct {
	Id         int
	Post       string
	Blogger    string
	InsertTime string
}

func GetAllPosts() []Post {
	db := db.ConnectDataBase()

	selectAllPosts, err := db.Query("select * from posts order by id_post desc")
	if err != nil {
		panic(err.Error())
	}

	p := Post{}
	posts := []Post{}

	for selectAllPosts.Next() {
		var id int
		var post, blogger, inserttime string

		err = selectAllPosts.Scan(&id, &post, &blogger, &inserttime)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Post = post
		p.Blogger = blogger
		p.InsertTime = inserttime

		posts = append(posts, p)
	}

	defer db.Close()
	return posts
}

func CreateNewPost(blogger, post string) {
	db := db.ConnectDataBase()

	newPostDB, err := db.Prepare("insert into posts(blogger, post) values($1, $2)")
	if err != nil {
		panic(err.Error())
	}

	newPostDB.Exec(blogger, post)
	defer db.Close()
}

func DeletPost(id string) {
	db := db.ConnectDataBase()

	deletCommtDB, err := db.Prepare("delete from commts where id_from_post=$1")
	if err != nil {
		panic(err.Error())
	}

	deletCommtDB.Exec(id)

	deletLikeDB, err := db.Prepare("delete from likes where id_from_post=$1")
	if err != nil {
		panic(err.Error())
	}

	deletLikeDB.Exec(id)

	deletPostDB, err := db.Prepare("delete from posts where id_post=$1")
	if err != nil {
		panic(err.Error())
	}

	deletPostDB.Exec(id)
	defer db.Close()
}

func EditPost(id string) Post {
	db := db.ConnectDataBase()

	postDB, err := db.Query("select * from posts where id_post=$1", id)
	if err != nil {
		panic(err.Error())
	}

	postContent := Post{}

	for postDB.Next() {
		var id int
		var blogger, post, inserttime string

		err = postDB.Scan(&id, &post, &blogger, &inserttime)
		if err != nil {
			panic(err.Error())
		}
		postContent.Id = id
		postContent.Post = post
		postContent.Blogger = blogger
		postContent.InsertTime = inserttime
	}
	defer db.Close()
	return postContent
}

func UpdatePost(id int, post, blogger string) {
	db := db.ConnectDataBase()

	UpdatePost, err := db.Prepare("update posts set post=$1, blogger=$2 where id_post=$3")
	if err != nil {
		panic(err.Error())
	}

	UpdatePost.Exec(post, blogger, id)
	defer db.Close()
}
