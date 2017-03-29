package main

import "database/sql"
import "fmt"
import _ "github.com/go-sql-driver/mysql"
import "os"

/**
siapkan struct sesuai dengan table database dalam kasus ini sesuaikan
dengan table student
*/
type student struct {
	id    string
	name  string
	age   int
	grade int
}

/**
crate function  for connection databse

*/
func connect() *sql.DB {
	//format connection user:password(host:port)/db_name
	var db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	return db
}

/**
prepare function sqlquery
*/
func sqlQUery() {
	//call function connection for connection with database
	var db = connect()
	defer db.Close() //whatever occur close db
	var age = 24
	var rows, err = db.Query("SELECT id,name,grade FROM tb_student WHERE age = ? ", age)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer rows.Close() //must close rows succes or failed

	//create variabel type array student
	var result []student
	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	for _, each := range result {
		fmt.Println(each.name)
	}
}

func main() {
	sqlQUery()
}
