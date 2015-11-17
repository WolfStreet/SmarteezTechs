package smarteeztechs

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Address struct {
	HouseNumber string `bson: "HouseNumber" json:"HouseNumber"`
	Street      string `bson:"Street" json:"Street"`
	Landmark    string `bson:"Landmark" json:"Landmark"`
	City        string `bson:"City" json:"City"`
	State       string `bson:"State" json:"State"`
	Country     string `bson:"Country" json:"Country"`
	Pin         string `bson:"Pin" json:"Pin"`
}

type User struct {
	ID          bson.ObjectId `bson:"_id" josn:"id"`
	UserName    string        `bson:"UserName" json:"UserName"`
	Name        string        `bson:"Name" json:"Name"`
	EMail       string        `bson:"EMail" json:"EMail"`
	Password    string        `bson:"Password" json:"Password"`
	Address     Address       `bson:"Address" json:"Address"`
	PhoneNumber string        `bson:"PhoneNumber" json:"PhoneNumber"`
	Controls    Controls      `bson:"Controls" json:"Controls"`
}

func (user *User) InitUser() MessageOut {
	user.ID = bson.NewObjectIdWithTime(time.Now())
	user.UserName = ""
	user.Name = ""
	user.Password = ""
	user.EMail = ""
	user.PhoneNumber = ""
	user.Address.HouseNumber = ""
	user.Address.Street = ""
	user.Address.Landmark = ""
	user.Address.City = ""
	user.Address.State = ""
	user.Address.Country = ""
	user.Address.Pin = ""

	var fan Control
	fan.Name = "Fan 1"
	fan.Speed = 80
	fan.Status = true
	user.Controls.Fan = append(user.Controls.Fan, fan)

	var Message MessageOut
	session, err := mgo.DialWithInfo(&DialInfo)
	defer session.Close()

	Message = CheckErr(err, "InitUser Session error")
	err = session.DB("smarteeztechs").C("User").Insert(&user)
	Message = CheckErr(err, "InitUser Insert error")
	return Message
}

func (user *User) FetchUserByID() MessageOut {
	var Message MessageOut

	session, err := mgo.DialWithInfo(&DialInfo)
	defer session.Close()
	Message = CheckErr(err, "Session Dial error @ Fetch User By Id")

	return Message
}

func (user *User) FetchUserByName() MessageOut {
	var Message MessageOut

	session, err := mgo.DialWithInfo(&DialInfo)
	defer session.Close()
	Message = CheckErr(err, "Session Dial error @ Fetch User By Name")

	return Message
}

func (user *User) SaveUser() MessageOut {
	var Message MessageOut
	session, err := mgo.DialWithInfo(&DialInfo)
	defer session.Close()

	Message = CheckErr(err, "SaveUser Session error")
	err = session.DB("smarteeztechs").C("User").Update(bson.M{"_id": user.ID}, &user)
	Message = CheckErr(err, "SaveUser Update error")
	return Message
}
