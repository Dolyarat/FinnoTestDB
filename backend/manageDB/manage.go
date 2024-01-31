package manageDB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	PersonID  int
	FirstName string
	LastName  string
	Tel       string
}

func SelectAll(db *sql.DB) []Person {
	var persons []Person
	results, err := db.Query("SELECT * FROM Persons")

	if err != nil {
		fmt.Println(err.Error())
		return persons
	}

	for results.Next() {
		var person Person
		err = results.Scan(&person.PersonID, &person.FirstName, &person.LastName, &person.Tel)
		if err != nil {
			fmt.Println(err.Error())
			return persons
		}
		persons = append(persons, person)
	}
	return persons
}

func Select(db *sql.DB, id int) Person {
	var person Person
	results, err := db.Query(`SELECT * FROM Persons WHERE PersonID = ?`, id)

	if err != nil {
		fmt.Println(err.Error())
		return person
	}
	err = results.Scan(&person.PersonID, &person.FirstName, &person.LastName, &person.Tel)

	for results.Next() {
		err = results.Scan(&person.PersonID, &person.FirstName, &person.LastName, &person.Tel)
		if err != nil {
			fmt.Println(err.Error())
			return person
		}
	}
	return person
}

func Insert(db *sql.DB, person Person) {
	insert, err := db.Query(`INSERT INTO Persons (FirstName, LastName, Tel) 
	VALUES (?, ?, ?)`, person.FirstName, person.LastName, person.Tel)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer insert.Close()
}

func Update(db *sql.DB, person Person, id int) {
	update, err := db.Query(`UPDATE Persons SET FirstName = ?, LastName = ?, Tel = ? 
	WHERE PersonID = ?`, person.FirstName, person.LastName, person.Tel, id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer update.Close()
}

func Delete(db *sql.DB, id int) {
	delete, err := db.Query(`DELETE FROM Persons WHERE PersonID = ?`, id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer delete.Close()
}
