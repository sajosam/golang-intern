// package main

// import (
// 	"adfolks/restapi/jwt"
// 	"adfolks/restapi/models"
// 	"encoding/json"
// 	"math/rand"
// 	"net/http"
// 	"strconv"

// 	"fmt"

// 	"github.com/gorilla/mux"
// )

// type Response struct {
// 	Persons []Person `json:"persons"`
// }

// type Person struct {
// 	Id        string    `json:"id"`
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// }

// var persons []Person

// func main() {
// 	fmt.Println("Hello, World!")
// 	models.ModelFrom()
// 	// models.Insertion()

// 	router := mux.NewRouter()

// 	persons = append(persons, Person{Id: "1", FirstName: "Issac", LastName: "Newton"})
// 	persons = append(persons, Person{Id: "2", FirstName: "Albert", LastName: "Einstein"})
// 	persons = append(persons, Person{Id: "3", FirstName: "Marie", LastName: "Curie"})
// 	persons = append(persons, Person{Id: "4", FirstName: "Charles", LastName: "Darwin"})
// 	persons = append(persons, Person{Id: "5", FirstName: "Nikola", LastName: "Tesla"})
// 	//specify endpoints, handler functions and HTTP method
// 	// http.Handle("/api", jwt.ValidateJWT(HealthCheck))
// 	// http.HandleFunc("/getjwt", jwt.GetJwt)
// 	// http.Handle("/person", jwt.ValidateJWT(Persons))
// 	// http.Handle("/person/{personID}", jwt.ValidateJWT(getPerson))
// 	// http.Handle("/aperson", jwt.ValidateJWT(addPerson))
// 	// http.Handle("/uperson/{personID}", jwt.ValidateJWT(updatePerson))
// 	// http.Handle("/dperson/{personID}", jwt.ValidateJWT(deletePerson))

// 	http.Handle("/", router)
// 	router.HandleFunc("/api", HealthCheck).Methods("GET")
// 	router.HandleFunc("/getjwt", jwt.GetJwt).Methods("GET")
// 	router.HandleFunc("/person", Persons).Methods("GET")
// 	router.HandleFunc("/person/{personID}", getPerson).Methods("GET")
// 	router.HandleFunc("/aperson", addPerson).Methods("POST")
// 	router.HandleFunc("/uperson/{personID}", updatePerson).Methods("PUT")
// 	router.HandleFunc("/dperson/{personID}", deletePerson).Methods("DELETE")

// //start and listen to requests
// 	http.ListenAndServe(":8080", router)
// }

// func HealthCheck(w http.ResponseWriter, r *http.Request) {
// 	//specify status code
// 	w.WriteHeader(http.StatusOK)
//   //update response writer
// 	fmt.Fprintf(w, "API is up and running")
// }

// func Persons(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(persons)
// }

// // get a single person

// func getPerson(w http.ResponseWriter, r *http.Request){

// 	w.Header().Set("Content-Type", "application/json")
// 	parms:=mux.Vars(r) //get parameters
// 	// loop through persons and find one with the id from the params
// 	for _, item := range persons {
// 		if item.Id == parms["personID"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Person{})

// }

// // add a new person

// func addPerson(w http.ResponseWriter, r *http.Request){

// 	w.Header().Set("Content-Type", "application/json")
// 	var person Person
// 	_ = json.NewDecoder(r.Body).Decode(&person)
// 	person.Id = strconv.Itoa(rand.Intn(1000000)) // Mock ID - not safe
// 	persons = append(persons, person)
// 	json.NewEncoder(w).Encode(person)
// }

// // // update a person

// func updatePerson(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	parms:=mux.Vars(r) //get parameters
// 	for index, item := range persons {
// 		if item.Id == parms["personID"] {
// 			persons = append(persons[:index], persons[index+1:]...)
// 			var person Person
// 			_ = json.NewDecoder(r.Body).Decode(&person)
// 			person.Id = parms["personID"]
// 			persons = append(persons, person)
// 			json.NewEncoder(w).Encode(person)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(persons)
// }

// // delete a person

// func deletePerson(w http.ResponseWriter, r *http.Request){

// 	w.Header().Set("Content-Type", "application/json")
// 	parms:=mux.Vars(r) //get parameters
// 	for index, item := range persons {
// 		if item.Id == parms["personID"] {
// 			persons = append(persons[:index], persons[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(persons)
// }

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

type students struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Age       int `json:"age "`
	Email     string `json:"email"`

}


var db *gorm.DB

func main() {
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

	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/persons", getPersons).Methods("GET")
	router.HandleFunc("/persons/{id}", getPerson).Methods("GET")
	router.HandleFunc("/newpersons", createPerson).Methods("POST")
	router.HandleFunc("/uppersons/{id}", updatePerson).Methods("PUT")
	router.HandleFunc("/delpersons/{id}", deletePerson).Methods("DELETE")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
	fmt.Println("server is running on 8080 port")

}



func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Super Secret Area")
}


func getPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var persons []students
	db.Find(&persons)
	json.NewEncoder(w).Encode(persons)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)	
	var person students
	db.First(&person, params["id"])
	json.NewEncoder(w).Encode(person)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person students
	_ = json.NewDecoder(r.Body).Decode(&person)

	db.Create(&person)
	json.NewEncoder(w).Encode(&person)

}


func updatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person students
	_ = json.NewDecoder(r.Body).Decode(&person)
	db.Save(&person)
	
}


func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var person students
	db.First(&person, params["id"])
	db.Delete(&person)
}

