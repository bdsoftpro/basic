package usrmodels

import (
	"database/sql"
	"errors"
	"app/system"
)

// User Format
type User struct {
	Id		int64	`json:"id"`
	Fname	string	`json:fname`
	Lname	string	`json:lname`
	Name	string	`json:"name"`
	Pass	string	`json:"pass"`
	Email	string	`json:"email"`
	Status	bool	`json:"status"`
}

// UserDB Format
type UserDB struct {
	Db *sql.DB
}

// Create Method Defination
func (udb UserDB) Create(user *User) error {
	result, err := udb.Db.Exec("insert into user(fname, lname, name, pass, email, status) values(?, ?, ?, ?, ?, ?)", user.Fname, user.Lname, user.Name, user.Pass, user.Email, user.Status)
	if err != nil {
		return err
	}
	user.Id, _ = result.LastInsertId()
	return nil
}

// Login Method Defination
func (udb UserDB) Login(usr User) (User, error) {
	var user User
	rows, err := udb.Db.Query("select id, fname, lname, name, email, pass, status from user where name = ?", usr.Name)
	if err != nil {
		return user, err
	}
	for rows.Next() {
		err2 := rows.Scan(&user.Id, &user.Fname, &user.Lname, &user.Name, &user.Email, &user.Pass, &user.Status)
		if err2 != nil {
			return user, err2
		}
	}
	if crypt.Decrypt(user.Pass) != usr.Pass {
		return User{}, errors.New("the password not match")
	}
	return user, nil
}

// Update Method Defination
func (udb UserDB) Update(user User) (int64, error) {
	result, err := udb.Db.Exec("update user set name = ?, email = ?, status = ? where id = ?", user.Name, user.Email, user.Status, user.Id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// Find Method Defination
func (udb UserDB) Find(id int64) (User, error) {
	rows, err := udb.Db.Query("select id, fname, lname, name, email, status from user where id = ?", id)
	if err != nil {
		return User{}, err
	}
	var user User
	for rows.Next() {
		err2 := rows.Scan(&user.Id, &user.Fname, &user.Lname, &user.Name, &user.Email, &user.Status)
		if err2 != nil {
			return User{}, err2
		}
	}
	return user, nil
}

// FindAll Method Defination
func (udb UserDB) FindAll() ([]User, error) {
	rows, err := udb.Db.Query("select id, fname, lname, name, email, status from user")
	if err != nil {
		return nil, err
	}
	users := []User{}
	for rows.Next() {
		var user User
		err2 := rows.Scan(&user.Id, &user.Fname, &user.Lname, &user.Name, &user.Email, &user.Status)
		if err2 != nil {
			return nil, err2
		}
		/*user := User{id, name, email, status}
		user := User{
			Id:       id,
			Name:     name,
			Email:    email,
			Status:   status,
		}*/
		users = append(users, user)
	}
	return users, nil
}
