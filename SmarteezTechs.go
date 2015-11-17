// SmarteezTechs project SmarteezTechs.go
package smarteeztechs

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./public_html/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./public_html/templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./public_html/templates/register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	r.ParseForm()
	fmt.Println(r)
	//var vars = mux.Vars(r)
	var user User
	var Message MessageOut
	//UserName := vars["UserName"]
	//Password := vars["Password"]
	UserName := r.PostFormValue("UserName")
	Password := r.PostFormValue("Password")

	Message, user = FetchControlsByUser(UserName, Password)

	if Message.Clear == false {
		MessageJSON, _ := json.Marshal(&Message)
		w.Write(MessageJSON)
	} else {
		ControlJSON, err := json.Marshal(&user.Controls)
		Message = CheckErr(err, "JSON Parsing of Controls @ Status Handler")
		if Message.Clear == false {
			MessageJSON, _ := json.Marshal(&Message)
			w.Write(MessageJSON)
		} else {
			w.Write([]byte(ControlJSON))
		}
	}
}

func TemporaryHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./public_html/templates/temporary.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
