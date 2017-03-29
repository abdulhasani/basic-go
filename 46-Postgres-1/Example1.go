package main

import (
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "prototype"
	password = "your_password"
	dbname   = "db_belajar_golang"
)

func main() {
	//test constant variable
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//opening connection our database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		//trace error
		panic(err)
	}
	defer db.Close()
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Successfully connected")
	sqlStatement := `INSERT INTO users (age,email,first_name,last_name) VALUES ($1,$2,$3,$4) RETURNING id`
	id := 0
	//insert data
	err = db.QueryRow(sqlStatement, 24, "oke.@gmail.com", "Kucing Belang", "Meong").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New Record ID is :", id)

}
