package mongo

import (
	"github.com/globalsign/mgo"
	"time"
)

const (
	hosts      = "localhost:27017"
	database   = "db"
	username   = ""
	password   = ""
	collection = "jobs"
)

func openMongoConnection() *mgo.Session {
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}
	return session
}