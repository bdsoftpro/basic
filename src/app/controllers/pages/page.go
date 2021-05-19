package pages

import (
	"fmt"
	"time"
	"net"
	"encoding/json"
	"net/http"
	"strconv"
	"html/template"
	"app/views"
	"app/system"
	"app/config"
	"app/models/pages"
	"app/models/users"
)
func getField(r *http.Request, index int) string{
	field := r.Context().Value("fields").([]string)
	return field[index]
}

// Home Methods Defination
func Home(w http.ResponseWriter, r *http.Request) {
	p := struct {
		Title string
		Body  string
	}{
		Title:	"Software Shop Management",
		Body:	"Home Page",
	}

	pageTpl, _ := resources.Asset("templates/pages/index.html")

	if _, err := net.DialTimeout("tcp","127.0.0.1:3306", 1000); err != nil {
		p.Body = "Refresh"
		pageTpl, _ = resources.Asset("templates/mysql.html")
	}

	t := template.Must(template.New("index.html").Parse(string(pageTpl)))
	if err := t.Execute(w, p); err != nil {
		panic(err)
	}
}

// Page Methods Defination
func Page(w http.ResponseWriter, r *http.Request) {
	pid, _ := strconv.ParseInt(getField(r, 0), 10, 64)
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		pdb := pgmodels.PageDB{
			Db: db,
		}
		page, err := pdb.Find(pid)
		if err != nil {
			pageTpl, _ := resources.Asset("templates/pages/notfound.html")
			t := template.Must(template.New("page.html").Parse(string(pageTpl)))
			item := struct {
				Name int64
			}{
				Name: pid,
			}
			if err := t.Execute(w, item); err != nil {
				panic(err)
			}
		} else {
			pageTpl, _ := resources.Asset("templates/pages/page.html")
			t := template.Must(template.New("page.html").Parse(string(pageTpl)))
			if err := t.Execute(w, page); err != nil {
				panic(err)
			}
		}
	}
}

// UserCheck Method Defination
func UserCheck(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		udb := usrmodels.UserDB{
			Db: db,
		}
		usr := usrmodels.User{
            Name:	r.FormValue("uname"),
            Pass:	r.FormValue("pass"),
		}
		user, err := udb.Login(usr)
		if err != nil {
			pageTpl, _ := resources.Asset("templates/users/notfound.html")
			t := template.Must(template.New("user.html").Parse(string(pageTpl)))
			item := struct {
				Name string
			}{
				Name: "Password",
			}
			if err := t.Execute(w, item); err != nil {
				panic(err)
			}
		} else {
			jsonBytes, _ :=json.Marshal(user)
			cookie := http.Cookie{
				Name: "session",
				Value: crypt.Encrypt(string(jsonBytes)),
				Path: "/",
				Expires: time.Now().AddDate(0, 0, 1),
				HttpOnly: true,
				MaxAge: 1*24*60*60,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/users", http.StatusMovedPermanently)
		}
	}
}

// Register Methods Defination
func Register(w http.ResponseWriter, r *http.Request) {
	p := struct {
		Title string
		Body  string
	}{
		Title:	"TestPage",
		Body:	"We have created a fictional band website. Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	}
	pageTpl, err := resources.Asset("templates/pages/register.html")
	if err != nil {
		pageTpl, _ = resources.Asset("templates/pages/notsource.html")
	}
	t := template.Must(template.New("register.html").Parse(string(pageTpl)))
	if err := t.Execute(w, p); err != nil {
		panic(err)
	}
}

// RegisterSave Methods Defination
func RegisterSave(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "register\n")
}

// Login Methods Defination
func Login(w http.ResponseWriter, r *http.Request) {
	p := struct {
		Title string
		Body  string
	}{
		Title:	"Login Page",
		Body:	"You can login by click <strong>Login</strong> button!",
	}
	pageTpl, _ := resources.Asset("templates/pages/login.html")
	t := template.Must(template.New("login.html").Parse(string(pageTpl)))
	if err := t.Execute(w, p); err != nil {
		panic(err)
	}
}

// SessionData Methods Defination
func SessionData(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")
	var user usrmodels.User
	json.Unmarshal([]byte(crypt.Decrypt(cookie.Value)), &user)
	fmt.Fprintf(w, "First Name: %s, Last Name: %s", user.Fname, user.Lname)
}

// Logout Methods Defination
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")
	fmt.Fprint(w, crypt.Decrypt(cookie.Value))
}

// Contact Methods Defination
func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "contact\n")
}
// PageCategory Methods Defination
func PageCategory(w http.ResponseWriter, r *http.Request) {
	page := getField(r, 0)
	cat := getField(r, 1)
	fmt.Fprintf(w, "Page is %s with %s Category\n", page, cat)
}