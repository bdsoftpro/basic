package users
import (
	"fmt"
	"regexp"
	"strconv"
	"net/http"
	"html/template"
	"app/views"
	"app/config"
	"app/models/users"
	"app/system"
)
/*
index()			--  view all	-- get
create()		--	create view	-- get
store()			--  save		-- post
show($id)		--	single view	-- get
edit($id)		--  edit view	-- get
update($id)		--  update		-- post
destroy($id)	--  delete		-- post, get

===============================================================================
Verb		URI								Action		Route Name
===============================================================================
GET			/users							index		users.index
-------------------------------------------------------------------------------
GET			/users/create					create		users.create
-------------------------------------------------------------------------------
POST		/users							store		users.store
-------------------------------------------------------------------------------
GET			/users/{user}					show		users.show
-------------------------------------------------------------------------------
GET			/users/{user}/edit				edit		users.edit
-------------------------------------------------------------------------------
PUT/PATCH	/users/{user}					update		users.update
-------------------------------------------------------------------------------
DELETE		/users/{user}					destroy		users.destroy
-------------------------------------------------------------------------------
-------------------------------------------------------------------------------
GET			/users/{user}/comments			index		users.comments.index
-------------------------------------------------------------------------------
GET			/users/{user}/comments/create	create		users.comments.create
-------------------------------------------------------------------------------
POST		/users/{user}/comments			store		users.comments.store
-------------------------------------------------------------------------------
GET			/comments/{comment}				show		comments.show
-------------------------------------------------------------------------------
GET			/comments/{comment}/edit		edit		comments.edit
-------------------------------------------------------------------------------
PUT/PATCH	/comments/{comment}				update		comments.update
-------------------------------------------------------------------------------
DELETE		/comments/{comment}				destroy		comments.destroy
===============================================================================
*/
func getField(r *http.Request, index int) string{
	field := r.Context().Value("fields").([]string)
	return field[index]
}

// Index Method Defination
func Index(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		udb := usrmodels.UserDB{
			Db: db,
		}
		users, err := udb.FindAll()
		if err != nil {
			pageTpl, _ := resources.Asset("templates/users/notfound.html")
			t := template.Must(template.New("user.html").Parse(string(pageTpl)))
			item := struct {
				Name string
			}{
				Name: "User",
			}
			if err := t.Execute(w, item); err != nil {
				panic(err)
			}
		} else {
			pageTpl, _ := resources.Asset("templates/users/index.html")
			t := template.Must(template.New("index.html").Parse(string(pageTpl)))
			if err := t.Execute(w, users); err != nil {
				panic(err)
			}
		}
	}
}

// Create Method Defination
func Create(w http.ResponseWriter, r *http.Request) {
	p := struct {
		Title string
		Body  string
	}{
		Title:	"TestPage",
		Body:	"We have created a fictional band website. Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	}
	pageTpl, err := resources.Asset("templates/pages/create.html")
	if err != nil {
		pageTpl, _ = resources.Asset("templates/pages/notsource.html")
	}
	t := template.Must(template.New("create.html").Parse(string(pageTpl)))
	if err := t.Execute(w, p); err != nil {
		panic(err)
	}
}

// Store Method Defination
func Store(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		udb := usrmodels.UserDB{
			Db: db,
		}
		usr := usrmodels.User{
			Fname:	r.FormValue("fname"),
			Lname:	r.FormValue("lname"),
            Name:	r.FormValue("uname"),
			Pass:	crypt.Encrypt(r.FormValue("pass")),
			Email:	r.FormValue("email"),
			Status:	true,
		}
		err := udb.Create(&usr)
		if err != nil {
			dpKey := regexp.MustCompile(`^Error [0-9]+: Duplicate entry '([^/]+)' for key '([^/]+)'.*$`).FindStringSubmatch(err.Error())
			pageTpl, _ := resources.Asset("templates/users/userexist.html")
			t := template.Must(template.New("user.html").Parse(string(pageTpl)))
			item := struct {
				Field string
				Value string
			}{
				Field:	dpKey[1:][1],
				Value:	dpKey[1:][0],
			}
			if err := t.Execute(w, item); err != nil {
				panic(err)
			}
		} else {
			http.Redirect(w, r, "/users", http.StatusMovedPermanently)
		}
	}
}

// Show Method Defination
func Show(w http.ResponseWriter, r *http.Request) {
	pid, _ := strconv.ParseInt(getField(r, 0), 10, 64)
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		udb := usrmodels.UserDB{
			Db: db,
		}
		user, err := udb.Find(pid)
		if err != nil {
			pageTpl, _ := resources.Asset("templates/users/notfound.html")
			t := template.Must(template.New("user.html").Parse(string(pageTpl)))
			item := struct {
				Name int64
			}{
				Name: pid,
			}
			if err := t.Execute(w, item); err != nil {
				panic(err)
			}
		} else {
			pageTpl, _ := resources.Asset("templates/users/user.html")
			t := template.Must(template.New("user.html").Parse(string(pageTpl)))
			if err := t.Execute(w, user); err != nil {
				panic(err)
			}
		}
	}
}

// Edit Method Defination
func Edit(w http.ResponseWriter, r *http.Request) {
	pid, _ := strconv.ParseInt(getField(r, 0), 10, 64)
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		udb := usrmodels.UserDB{
			Db: db,
		}
		user, err := udb.Find(pid)
		if err != nil {
			pageTpl, _ := resources.Asset("templates/users/notfound.html")
			t := template.Must(template.New("user.html").Parse(string(pageTpl)))
			item := struct {
				Name int64
			}{
				Name: pid,
			}
			if err := t.Execute(w, item); err != nil {
				panic(err)
			}
		} else {
			pageTpl, _ := resources.Asset("templates/users/user.html")
			t := template.Must(template.New("user.html").Parse(string(pageTpl)))
			if err := t.Execute(w, user); err != nil {
				panic(err)
			}
		}
	}
}

// Update Method Defination
func Update(w http.ResponseWriter, r *http.Request) {
	pid, _ := strconv.ParseInt(getField(r, 0), 10, 64)
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		udb := usrmodels.UserDB{
			Db: db,
		}
		user, err := udb.Find(pid)
		if err != nil {
			pageTpl, _ := resources.Asset("templates/users/notfound.html")
			t := template.Must(template.New("user.html").Parse(string(pageTpl)))
			item := struct {
				Name int64
			}{
				Name: pid,
			}
			if err := t.Execute(w, item); err != nil {
				panic(err)
			}
		} else {
			pageTpl, _ := resources.Asset("templates/users/user.html")
			t := template.Must(template.New("user.html").Parse(string(pageTpl)))
			if err := t.Execute(w, user); err != nil {
				panic(err)
			}
		}
	}
}

// Destroy Method Defination
func Destroy(w http.ResponseWriter, r *http.Request) {
	pid, _ := strconv.ParseInt(getField(r, 0), 10, 64)
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		udb := usrmodels.UserDB{
			Db: db,
		}
		user, err := udb.Find(pid)
		if err != nil {
			pageTpl, _ := resources.Asset("templates/users/notfound.html")
			t := template.Must(template.New("user.html").Parse(string(pageTpl)))
			item := struct {
				Name int64
			}{
				Name: pid,
			}
			if err := t.Execute(w, item); err != nil {
				panic(err)
			}
		} else {
			pageTpl, _ := resources.Asset("templates/users/user.html")
			t := template.Must(template.New("user.html").Parse(string(pageTpl)))
			if err := t.Execute(w, user); err != nil {
				panic(err)
			}
		}
	}
}