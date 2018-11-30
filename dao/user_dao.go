package dao

import (
	"log"
	. "bank-service/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

func (u *UserDAO) Connect() {
	session, err := mgo.Dial(u.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(u.Database)
}

func (u *UserDAO) FindAll() ([]User, error) {
	var user []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&user)
	return user, err
}

func (u *UserDAO) Insert(user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}