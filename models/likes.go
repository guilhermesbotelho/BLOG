package models

import (
	"github.com/guilhermesbotelho/BLOG/db"
)

type Like struct {
	UserName string
}

type CountLike struct {
	Count int
}

func GetAllLikes(id string) []Like {
	db := db.ConnectDataBase()

	selectAllLikes, err := db.Query("select * from likes where id_from_post=$1", id)
	if err != nil {
		panic(err.Error())
	}

	l := Like{}
	likes := []Like{}

	for selectAllLikes.Next() {
		var idlikes, idpost int
		var username string

		err = selectAllLikes.Scan(&idlikes, &username, &idpost)
		if err != nil {
			panic(err.Error())
		}

		l.UserName = username

		likes = append(likes, l)
	}

	defer db.Close()
	return likes
}

func CountLikes(id string) []CountLike {
	db := db.ConnectDataBase()

	countLikes, err := db.Query("select count(*) from likes where id_from_post=$1", id)
	if err != nil {
		panic(err.Error())
	}

	c := CountLike{}
	countlikes := []CountLike{}

	for countLikes.Next() {
		var count int

		err = countLikes.Scan(&count)
		if err != nil {
			panic(err.Error())
		}

		c.Count = count

		countlikes = append(countlikes, c)
	}

	defer db.Close()
	return countlikes
}
