package models

import (
	"log"

	"github.com/guilhermesbotelho/BLOG/db"
)

type Commt struct {
	IdCommt  int
	Commt    string
	IdPost   int
	UserName string
}

func GetAllCommts(id string) []Commt {
	db := db.ConnectDataBase()

	selectAllCommts, err := db.Query("select * from commts where id_from_post=$1 order by id_commt desc", id)
	if err != nil {
		panic(err.Error())
	}

	c := Commt{}
	commts := []Commt{}

	for selectAllCommts.Next() {
		var idCommt, idPost int
		var commt, userName string

		err = selectAllCommts.Scan(&idCommt, &commt, &idPost, &userName)
		if err != nil {
			panic(err.Error())
		}

		c.IdCommt = idCommt
		c.Commt = commt
		c.IdPost = idPost
		c.UserName = userName

		commts = append(commts, c)
	}
	defer db.Close()
	return commts
}

func CreateNewCommt(username, commt string, id int) string {
	db := db.ConnectDataBase()

	newPostDB, err := db.Prepare("insert into commts(username, commt, id_from_post) values($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}

	_, err = newPostDB.Exec(username, commt, id)
	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}

	defer db.Close()
	return ""
}

func DeletComment(id string) {
	db := db.ConnectDataBase()

	deletCommtDB, err := db.Prepare("delete from commts where id_commt=$1")
	if err != nil {
		panic(err.Error())
	}

	deletCommtDB.Exec(id)
	defer db.Close()
}
