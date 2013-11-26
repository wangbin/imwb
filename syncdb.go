package main

import (
	"fmt"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/wangbin/imwb/models"
	"time"
)

const (
	DbName        = "imwb"
	UserTableName = "auth_user"
	AuthKey       = "wangbin7972"
	Address       = "localhost:28015"
)

var (
	session       *r.Session
	err           error
	userTableSpec = r.TableSpec{Name: UserTableName}
)

func reCreateTable(session *r.Session) {
	var tables []string
	r.TableList().Run(session).All(&tables)
	for _, table := range tables {
		r.TableDrop(table).Run(session).Exec()
	}
	r.TableCreateWithSpec(userTableSpec).Run(session).Exec()
	var response map[string]int
	r.Table(UserTableName).IndexCreate("username", nil).Run(session).All(&response)
}

func createUsers(session *r.Session) {
	var err error
	admin := &models.User{
		UserName:    "admin",
		Password:    "1234",
		Email:       "admin@example.com",
		IsSuperUser: true,
		IsActive:    true,
		DateJoined:  time.Now(),
		LastLogin:   time.Now(),
		Groups:      []string{"admin"},
	}
	err = admin.Save(session)
	if err != nil {
		panic(err)
	}
	fmt.Println(admin.Id)
	wangbin := &models.User{
		UserName:    "wangbin",
		Password:    "1234",
		Email:       "admin@example.com",
		IsSuperUser: true,
		IsActive:    true,
		DateJoined:  time.Now(),
		LastLogin:   time.Now(),
		Groups:      []string{"admin", "author"},
	}
	err = wangbin.Save(session)
	if err != nil {
		panic(err)
	}
	fmt.Println(wangbin.Id)
}

func main() {
	session, err = r.ConnectWithAuth(Address, DbName, AuthKey)
	if err != nil {
		panic(err)
	}
	reCreateTable(session)
	createUsers(session)
}
