package main

import (
	"encoding/json"
	r "github.com/dancannon/gorethink"
	"github.com/wangbin/imwb/models"
	"io/ioutil"
)

var (
	session *r.Session
	err     error
)

func init() {
	session, err = models.GetSession()
	if err != nil {
		panic(err)
	}
}

func tableList() (tables []string) {
	rows, _ := r.Db(models.DbName).TableList().Run(session)
	rows.ScanAll(&tables)
	return tables
}

func dropAll() {
	tables := tableList()
	for _, table := range tables {
		r.Db(models.DbName).TableDrop(table).Exec(session)
	}
}

func createUserTable() {
	r.Db(models.DbName).TableCreate(models.UserTable).Exec(session)
	r.Table(models.UserTable).IndexCreate(models.UserTableSecondIndex).Exec(session)
}

func createPostTable() {
	r.Db(models.DbName).TableCreate(models.PostTable).Exec(session)
	r.Table(models.PostTable).IndexCreate(models.PostTableSecondIndex).Exec(session)
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
	err = admin.Save(session)
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
	err = wangbin.Save(session)
	if err != nil {
		panic(err)
	}
}

func loadPosts() {
	content, _ := ioutil.ReadFile("/Users/wangbin/Downloads/dump.json")
	var posts []*models.Post
	json.Unmarshal(content, &posts)
	r.Table(models.PostTable).Insert(posts).RunWrite(session)
	var users []*models.User
	rows, _ := r.Table(models.UserTable).GetAllByIndex("username", "wangbin").Run(
		session)
	rows.ScanAll(&users)
	author := users[0]
	r.Table(models.PostTable).Update(models.RethinkMap{"user_id": author.Id}).RunWrite(
		session)
}

func main() {
	dropAll()
	createUserTable()
	createPostTable()
	createUsers()
	loadPosts()
}
