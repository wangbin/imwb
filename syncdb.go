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
	fmt.Println(response)
}

func createUsers(session *r.Session) {
	admin := models.User{
		UserName:    "admin",
		Password:    "1234",
		Email:       "admin@example.com",
		IsSuperUser: true,
		IsActive:    true,
		DateJoined:  time.Now(),
		LastLogin:   time.Now(),
		Groups:      []string{"admin"},
	}
	wangbin := models.User{
		UserName:    "wangbin",
		Password:    "1234",
		Email:       "admin@example.com",
		IsSuperUser: true,
		IsActive:    true,
		DateJoined:  time.Now(),
		LastLogin:   time.Now(),
		Groups:      []string{"admin", "author"},
	}

	var response r.WriteResponse
	r.Table(UserTableName).Insert([]models.User{admin, wangbin}).Run(session).One(&response)
	fmt.Println(response)
}

func main() {
	session, err = r.ConnectWithAuth(Address, DbName, AuthKey)
	if err != nil {
		panic(err)
	}
	reCreateTable(session)
	createUsers(session)
}
