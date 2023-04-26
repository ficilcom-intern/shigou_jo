package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")

	defer Db.Close()

	//////create new tale

	cmd := `create table if not exists persons(
					name string,
					age int
				)`
	_, err := Db.Exec(cmd)

	//////insert data

	cmd = `insert into persons(name, age) values (?, ?)`
	_, err = Db.Exec(cmd, "tarou", 20)

	//////update data

	cmd = `update persons set age = ? where name = ?`
	_, err = Db.Exec(cmd, 25, "tarou")

	//////select one data

	cmd = `insert into persons(name, age) values(?, ?)`
	_, err = Db.Exec(cmd, "hanako", 19)

	cmd = `select * from persons where age = ?`
	//get one record
	row := Db.QueryRow(cmd, 25)
	var p Person
	err = row.Scan(&p.Name, &p.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}

	//////select multiple datas

	cmd = `select * from persons`
	rows, _ := Db.Query(cmd)
	defer rows.Close()

	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}

	for _, p := range pp {
		fmt.Println(p)
	}

	//////delete data
	cmd = `delete from persons where name = ?`
	_, err = Db.Exec(cmd, "hanako")
	if err != nil {
		log.Fatalln(err)
	}

}
