package models

import (
	r "github.com/dancannon/gorethink"
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
	}
	Conn, _ = GetSession()
}

func GetSession() (*r.Session, error) {
	return r.Connect(Database)
}
