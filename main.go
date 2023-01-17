package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "adfolks/restapi/restapi"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type students struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age "`
	Email string `json:"email"`
}

var db *gorm.DB


func main() {
	// restapi.Connection()
	var err error

	dsn := "host=localhost user=postgres password=root dbname=forgolang port=5433 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}else{
		fmt.Println("connected to database")
	}
	dbinstance, _ := db.DB()

	defer dbinstance.Close()

	db.AutoMigrate(&students{})





	router := mux.NewRouter()
	// router.HandleFunc("/health-check", restapi.HealthCheck).Methods("GET")
	// router.HandleFunc("/persons", restapi.GetPersons).Methods("GET")
	// router.HandleFunc("/persons/{id}", restapi.GetPerson).Methods("GET")
	// router.HandleFunc("/newpersons", restapi.CreatePerson).Methods("POST")
	// router.HandleFunc("/uppersons/{id}", restapi.UpdatePerson).Methods("PUT")
	// router.HandleFunc("/delpersons/{id}", restapi.DeletePerson).Methods("DELETE")

	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/persons", GetPersons).Methods("GET")
	router.HandleFunc("/persons/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/newpersons", CreatePerson).Methods("POST")
	router.HandleFunc("/uppersons/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/delpersons/{id}", DeletePerson).Methods("DELETE")


	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
	fmt.Println("server is running on 8080 port")

}




func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Super Secret Area")
}

func GetPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var persons []students
	db.Find(&persons)
	json.NewEncoder(w).Encode(persons)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var person students
	db.First(&person, params["id"])
	json.NewEncoder(w).Encode(person)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person students
	_ = json.NewDecoder(r.Body).Decode(&person)

	db.Create(&person)
	json.NewEncoder(w).Encode(&person)

}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person students
	_ = json.NewDecoder(r.Body).Decode(&person)
	db.Save(&person)

}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var person students
	db.First(&person, params["id"])
	db.Delete(&person)
}
