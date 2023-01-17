package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
  gorm.Model
  Id int `json:"id" gorm:"primary_key" autoIncrement:"true"`
  Name string `json:"name" gorm:"not null"`
  Age int `json:"age" gorm:"default:20" gorm:"not null" where:"age>18"`
  Email string `json:"email" gorm:"unique" gorm:"not null"`
}

var (
	  db *gorm.DB
)

func ModelFrom() {
	dsn := "host=localhost user=postgres password=root dbname=forgolang port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})


	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Student{})

	// add 5 students
	db.Create(&Student{Name: "Issac", Age: 20, Email: "demo@gmail.com"})
	db.Create(&Student{Name: "Albert", Age: 20, Email: "demo1@gmail.com"})
	db.Create(&Student{Name: "Marie", Age: 20, Email: "demo3@gmail.com"})
	db.Create(&Student{Name: "Charles", Age: 20, Email: "demo4@gmail.com"})
	db.Create(&Student{Name: "Nikola", Age: 20, Email: "demo5@gmail.com"})

	// Read
	var student Student
	db.First(&student, 1) // find student with integer primary key
	db.First(&student, "name = ?", "Issac") // find student with name Issac

	// fmt.Println(student.Name, student.Age, student.Email)

	// read all
	var students []Student
	db.Find(&students)
	fmt.Println(students)
	



}

// func Insertion(){
// 	db.Create(&Student{Name: "demo", Age: 20, Email: "demo10@gmai.com"})
// }


// func Deletion(){
// 	db.Delete(&Student{}, 1) // delete product
// }

// func Update(){
// 	db.Model(&Student{}).Where("name = ?", "demo").Update("name", "demo1")
// }



func getUsers() {
	var student Student
	db.First(&student, 1)
	// fmt.Println(student.Name, student.Age, student.Email)
	
}