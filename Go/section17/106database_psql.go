package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB
var err error

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err = sql.Open("postgres", "user=test_user dbname=test_db password=password sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	///////create

	cmd := "insert into persons(name, age) values($1, $2)"
	_, err := Db.Exec(cmd, "Nancy", 20)
	if err != nil {
		log.Fatalln(err)
	}

	//////read
	cmd = "select * from persons where age = $1"
	//get one record
	row := Db.QueryRow(cmd, 20)
	var p Person
	err = row.Scan(&p.Name, &p.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p)

	cmd = "select * from persons"
	//multiple rows
	rows, _ := Db.Query(cmd)
	defer rows.Close()

	var pp []Person
	for rows.Next() {
		var p Person
		err = rows.Scan(&p.Name, &p.Age)
		//handle errors one by one
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}

	//handle errors collectively
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	for _, p := range pp {
		fmt.Println(p)
	}

	//////update
	cmd = "update persons set age = $1 where name = $2"
	_, err = Db.Exec(cmd, 25, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}

	//////delete
	cmd = "delete from persons where name = $1"
	_, err = Db.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}

}
