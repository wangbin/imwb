package main

import (
	"encoding/json"
	"fmt"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/wangbin/imwb/models/auth"
	"github.com/wangbin/imwb/models/blog"
	"github.com/wangbin/imwb/settings"
	"io/ioutil"
)

var (
	session       *r.Session
	err           error
	userTableSpec = r.TableSpec{Name: auth.UserTable}
	postTableSpec = r.TableSpec{Name: blog.PostTable}
)

func reCreateTable(session *r.Session) {
	var tables []string
	r.TableList().Run(session).All(&tables)
	for _, table := range tables {
		r.TableDrop(table).Run(session).Exec()
	}
	r.TableCreateWithSpec(userTableSpec).Run(session).Exec()
	r.TableCreateWithSpec(postTableSpec).Run(session).Exec()
	var response map[string]int
	r.Table(auth.UserTable).IndexCreate("username", nil).Run(session).All(&response)
	r.Table(blog.PostTable).IndexCreate("slug", nil).Run(session).All(&response)
}

func createUsers(session *r.Session) {
	admin := &auth.User{
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
	fmt.Println(admin.Id)
	wangbin := &auth.User{
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
	fmt.Println(wangbin.Id)
}

func loadPosts(session *r.Session) {
	content, _ := ioutil.ReadFile("/Users/wangbin/Downloads/dump.json")
	var posts []*blog.Post
	json.Unmarshal(content, &posts)
	var response r.WriteResponse
	r.Table(blog.PostTable).Insert(posts).Run(session).One(&response)
	fmt.Println(response)
}

func main() {
	//	session, err = r.ConnectWithAuth(Address, DbName, AuthKey)
	session, err = r.Connect(settings.DbUri, settings.DbName)
	if err != nil {
		panic(err)
	}
	reCreateTable(session)
	createUsers(session)
	loadPosts(session)
}
