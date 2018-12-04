package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

func main() {
	// connect to the database
	db, err := sql.Open("mysql", "gouser:gouser@/Dino")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//general query with arguments
	rows, err := db.Query("select * from Dino.animals where id=?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		animals = append(animals, a)
	}
	fmt.Println(animals)

	// query a single row
	row := db.QueryRow("select * from Dino.animals where id=?", 2)
	a := animal{}
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)

	// insert a row
	result, err := db.Exec("Insert into Dino.animals (animal_type, nickname, zone, age) values (Carnotaurus, Carno, 3, 22)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	// Update a row
	age := 24
	id := 2
	result2, err := db.Exec("Update Dino.animals set age=? where id=?", age, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result2.LastInsertId())
	fmt.Println(result2.RowsAffected())
}
