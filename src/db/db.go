package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

type Person struct {
	FirstName string
	LastName string
}

func GetFirstNames() []Person {
	rows, err := DB.Query("Select fname, lname from youtube")
	if err != nil {
		fmt.Println("[db.go] Error line 19", err)
	}
	defer rows.Close()

	var names []Person

	for rows.Next() {
		var fname string
		var lname string
		err := rows.Scan(&fname, &lname)
		if err != nil {
			fmt.Println("[db.go] Error line 30", err)
		}

		names = append(names, Person{
			FirstName: fname,
			LastName:  lname,
		})
	}

	return names
}

func New() (*sql.DB, error) {
	conn := "user=postgres dbname=postgres password=mysecretpassword host=0.0.0.0 sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", conn)
	if err != nil{
		fmt.Println("[db.go] Error on line 48", err)
	}

	return DB, DB.Ping()
}


