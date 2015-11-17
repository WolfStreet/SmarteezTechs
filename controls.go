package smarteeztechs

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ON = true
const OFF = false

type Control struct {
	Name   string `bson:"Name" json:"Name"`
	Status bool   `bson:"Status" json:"Status"`
	Speed  int    `bson:"Speed" json:"Speed"`
}

type Controls struct {
	Fan   []Control `bson: "Fan" json:"Fan"`
	Light []Control `bson: "Light" json:"Light"`
}

func FetchControlsByUser(UserName, Password string) (MessageOut, User) {
	var user User
	session, err := mgo.DialWithInfo(&DialInfo)
	defer session.Close()
	Message := CheckErr(err, "Session login @ Fetch Controls By User")
	err = session.DB("smarteeztechs").C("User").Find(bson.M{"UserName": UserName, "Password": Password}).One(&user)
	Message = CheckErr(err, "Search Controls @ Fetch Controls By User")
	return Message, user
}
