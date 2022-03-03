package accountcontroller

import (
	"fmt"
	"html/template"
	"net/http"

	sessions "github.com/ermiaswalelgne/receipt_uploader/src/sessions"
)

func Index(res http.ResponseWriter, req *http.Request) {
	tmp, _ := template.ParseFiles("views/accountcontroller/index.html")
	tmp.Execute(res, nil)
}

// function to athenticate users
func Login(res http.ResponseWriter, req *http.Request) {
	//Check the request is method is POST and the url is the right path
	if req.URL.Path != "/account/login" {
		http.Error(res, "404 not found.", http.StatusNotFound)
		return
	}
	if req.Method != "POST" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Read values from form
	req.ParseForm()
	username := req.Form.Get("username")
	password := req.Form.Get("password")
	fmt.Println("username: ", username, "\npassword: ", password)
	//Check if ussername and password are not empty and if the password and user name maches
	if len(username) == 0 || len(password) == 0 {
		data := map[string]interface{}{
			"err": "User name or password can not be empty.",
		}
		tmp, _ := template.ParseFiles("views/accountcontroller/index.html")
		tmp.Execute(res, data)
	} else if username != "testuser" && password != "secret" {
		data := map[string]interface{}{
			"err": "Invalid username or password. Please try again",
		}
		tmp, _ := template.ParseFiles("views/accountcontroller/index.html")
		tmp.Execute(res, data)
	} else {
		// if login is successful create user's session and redirect to the receipt uploader page
		session, err := sessions.Get(req)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		session.Values["username"] = username
		fmt.Println("User session", session.Values["username"])
		err = session.Save(req, res)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		http.Redirect(res, req, "/receipt/uploadimage", http.StatusSeeOther)
	}
}
