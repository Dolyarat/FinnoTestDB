package main

import (
	"backend/manageDB"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func personsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			personsJSON, err := json.Marshal(manageDB.SelectAll(db))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-type", "application/json")
			w.Write(personsJSON)
		case http.MethodPost:
			var newPerson manageDB.Person
			bodyByte, err := ioutil.ReadAll(r.Body)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			err = json.Unmarshal(bodyByte, &newPerson)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			manageDB.Insert(db, newPerson)
		}
	}
}

func personsHandlerByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/persons/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		switch r.Method {
		case http.MethodGet:
			personJSON, err := json.Marshal(manageDB.Select(db, id))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-type", "application/json")
			w.Write(personJSON)
		case http.MethodPut:
			var changePerson manageDB.Person
			bodyByte, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			err = json.Unmarshal(bodyByte, &changePerson)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			manageDB.Update(db, changePerson, id)
		case http.MethodDelete:
			manageDB.Delete(db, id)
		}
	}
}

func enableCorsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-type, Content-Length, Authorization, X-Custom-Header, Upgrade-Insecure-Requests")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/FinnoTestDB")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("success connect db")
	}
	defer db.Close()

	http.Handle("/persons", enableCorsMiddleware(personsHandler(db)))
	http.Handle("/persons/", enableCorsMiddleware(personsHandlerByID(db)))
	http.ListenAndServe(":5000", nil)
}
