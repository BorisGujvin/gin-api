package content

import (
	"database/sql"
	"fmt"

	"github.com/BorisGujvin/gin-api/config"

	_ "github.com/go-sql-driver/mysql"
)

func GetContentList() []content {
	var template = "SELECT id, name FROM contents"
	var connectString = config.GetMySQLConnectionString()
	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	results, err := db.Query(template)
	if err != nil {
		panic(err.Error())
	}
	var mySlice []content

	for results.Next() {
		var content content
		err = results.Scan(&content.Id, &content.Name)
		if err != nil {
			panic(err.Error())
		}
		row := content
		mySlice = append(mySlice, row)
	}
	return mySlice
}

func CreateContent(name string) bool {
	var template = "INSERT INTO contents (`name`) values(\"%s\")"
	var connectString = config.GetMySQLConnectionString()
	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var operator = fmt.Sprintf(template, name)
	results, err := db.Query(operator)
	if err != nil {
		panic(err.Error())
	}
	fmt.Print(results)
	return true
}
