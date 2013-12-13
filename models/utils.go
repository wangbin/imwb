package models

import (
	r "github.com/dancannon/gorethink"
	"time"
)

const (
	DbUri  = "localhost:28015"
	DbName = "imwb"
)

type RethinkMap map[string]interface{}

var (
	Database RethinkMap
	Conn     *r.Session
)

func init() {
	Database = RethinkMap{
		"address":  DbUri,
		"database": DbName,
		//        "authkey":  "14daak1cad13dj",
		"maxIdle":     10,
		"idleTimeout": time.Second * 30,
		"maxActive":   50,
	}
	Conn, _ = getSession()
}

func getSession() (*r.Session, error) {
	return r.Connect(Database)
}
