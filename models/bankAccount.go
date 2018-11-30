package models 

import (
	"gopkg.in/mgo.v2/bson"
)

type BankAccount struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	BankId string `bson:"bid" json:"bid"`
	UserId string `bson:"uid" json:"uid"`
	AccountNumber int `bson:"account_number" json:"account_number"`
	Balance int `bson:"balance" json:"balance"`
}