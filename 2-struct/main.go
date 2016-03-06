package main

import (
	"fmt"
	"time"
)

type user struct {
	Name string `json:"name"`
	DOB  string `json:"date_of_birth"`
	City string `json:"city"`
}

func (u *user) sayHello() {
	fmt.Println("Hello", u.Name)
}

func (u *user) sayHistory() {

	layout := "January 2, 2006"
	t, _ := time.Parse(layout, u.DOB)
	bettyYear := t.Year()
	now := time.Now()
	thisYear := now.Year()
	bettyAge := thisYear - bettyYear

	fmt.Printf("%s who was born in %s would be... The current date: %d Holberton's birthdate: %d.\n", u.Name, u.City, thisYear, bettyAge)
}

func main() {
	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
	u.sayHello()
	u.sayHistory()
}
