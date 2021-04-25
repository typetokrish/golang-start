package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}
type Category struct {
	Id   int
	Name string
}

//Go application entrypoint
func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/api/", apiResponse).Methods(http.MethodGet)
	myRouter.HandleFunc("/api/categories", getCategories).Methods(http.MethodGet)
	myRouter.HandleFunc("/api/categories", postCategory).Methods(http.MethodPost)
	//This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))

	}).Methods(http.MethodGet)

	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	//Print any errors from starting the webserver using fmt
	fmt.Println(http.ListenAndServe(":8080", myRouter))
}

func apiResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))

}

func getCategories(w http.ResponseWriter, r *http.Request) {
	var categories []Category
	db, err := sql.Open("mysql", "admin:Start#123@tcp(127.0.0.1:3306)/go_test")
	defer db.Close()
	if err != nil {
		fmt.Println("Error Connecting to Database", err)
		log.Fatalln("The database connection failed for the request", err)
	}

	res, errQ := db.Query("SELECT * FROM category")
	if errQ != nil {
		fmt.Println("Cannot fetch data at this time")
		log.Fatalln("Select query failed", errQ)

	}

	for res.Next() {
		var cat Category
		errR := res.Scan(&cat.Id, &cat.Name)

		if errR != nil {
			fmt.Println("Fetch Failed")

		}
		categories = append(categories, cat)
		fmt.Printf("%v/n", cat)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	//json_data2, err := json.Marshal(categories)
	json.NewEncoder(w).Encode(categories)

}

func postCategory(w http.ResponseWriter, r *http.Request) {

	var cat Category
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", "admin:Start#123@tcp(127.0.0.1:3306)/go_test")
	defer db.Close()
	if err != nil {
		fmt.Println("Error Connecting to Database", err)
		log.Fatalln("The database connection failed for the request", err)
	}

	sql := "INSERT INTO category (name) VALUES (?)"
	res, errQ := db.Exec(sql, cat.Name)
	if errQ != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	id, errId := res.LastInsertId()
	if errId != nil {
		http.Error(w, err.Error(), http.StatusAccepted)
	}
	cat.Id = int(id)

	json.NewEncoder(w).Encode(cat)

}
