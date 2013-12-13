package main

import (
	"encoding/json"
	r "github.com/dancannon/gorethink"
	"github.com/wangbin/imwb/models"
	"io/ioutil"
)

var (
	err error
)

func tableList() (tables []string) {
	rows, _ := r.Db(models.DbName).TableList().Run(models.Conn)
	rows.ScanAll(&tables)
	return tables
}

func dropAll() {
	tables := tableList()
	for _, table := range tables {
		r.Db(models.DbName).TableDrop(table).Exec(models.Conn)
	}
}

func createUserTable() {
	r.Db(models.DbName).TableCreate(models.UserTable).Exec(models.Conn)
	r.Table(models.UserTable).IndexCreate(models.UserTableSecondIndex).Exec(models.Conn)
}

func createPostTable() {
	r.Db(models.DbName).TableCreate(models.PostTable).Exec(models.Conn)
	r.Table(models.PostTable).IndexCreate(models.PostTableSecondIndex).Exec(models.Conn)
}

func createUsers() {
	admin := &models.User{
		UserName:    "admin",
		Email:       "admin@example.com",
		IsSuperUser: true,
		IsActive:    true,
		Groups:      []string{"admin"},
	}
	admin.SetPassword("1234")
	err = admin.Save()
	if err != nil {
		panic(err)
	}
	wangbin := &models.User{
		UserName:    "wangbin",
		Email:       "admin@example.com",
		IsSuperUser: true,
		IsActive:    true,
		Groups:      []string{"admin", "author"},
	}
	wangbin.SetPassword("1234")
	err = wangbin.Save()
	if err != nil {
		panic(err)
	}
}

func loadPosts() {
	content, _ := ioutil.ReadFile("/Users/wangbin/Downloads/dump.json")
	var posts []*models.Post
	json.Unmarshal(content, &posts)
	r.Table(models.PostTable).Insert(posts).RunWrite(models.Conn)
	var users []*models.User
	rows, _ := r.Table(models.UserTable).GetAllByIndex("username", "wangbin").Run(
		models.Conn)
	rows.ScanAll(&users)
	author := users[0]
	r.Table(models.PostTable).Update(models.RethinkMap{"user_id": author.Id}).RunWrite(
		models.Conn)
}

func main() {
	dropAll()
	createUserTable()
	createPostTable()
	createUsers()
	loadPosts()
}
