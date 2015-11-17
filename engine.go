package smarteeztechs

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var DialInfo mgo.DialInfo = mgo.DialInfo{
	//Addrs:    []string{fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_MONGODB_DB_HOST"), os.Getenv("OPENSHIFT_MONGODB_DB_PORT"))},
	Addrs:    []string{"127.4.15.2:27017"},
	Database: "smarteeztechs",
	Password: " hq8YHHYEWgu-",
	Username: "admin",
}

type MessageOut struct {
	Message string
	Clear   bool
}

func CheckErr(err error, Module string) MessageOut {
	var Message MessageOut
	if err != nil {
		fmt.Println(Module)
		fmt.Println(err.Error())
		Message.Message = err.Error()
		Message.Clear = false
	} else {
		Message.Message = ""
		Message.Clear = true
	}
	return Message
}
