package user

import (
	"database/sql"
	"strconv"

	"github.com/BorisGujvin/gin-api/config"
	_ "github.com/go-sql-driver/mysql"
)

func GetUserList() []User {
	var template = "SELECT id, name, email, password FROM users"
	var db = getDb()
	defer db.Close()
	results, err := db.Query(template)
	if err != nil {
		panic(err.Error())
	}
	var mySlice []User

	for results.Next() {
		var user User
		err = results.Scan(&user.Id, &user.Name, &user.Email, &user.PassHash)
		if err != nil {
			panic(err.Error())
		}
		mySlice = append(mySlice, user)
	}
	return mySlice
}

func StoreUser(user User) (User, error) {
	var template = `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`
	var db = getDb()
	defer db.Close()
	result, err := db.Exec(template, user.Name, user.Email, user.PassHash)
	if err != nil {
		return user, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}
	user.Id = strconv.FormatInt(id, 10)
	return user, nil
}

func getDb() *sql.DB {
	var connectString = config.GetMySQLConnectionString()
	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err.Error())
	}
	return db
}
