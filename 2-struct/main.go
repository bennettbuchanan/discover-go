package main

import (
	"fmt"
	"strconv"
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

	now := time.Now()
	nowFormatted := now.Format("20060102")
	birthday, _ := time.Parse("January 2, 2006", u.DOB)
	birthdayParsed := birthday.Format("20060102")
	x, err := strconv.Atoi(nowFormatted)
	y, err := strconv.Atoi(birthdayParsed)
	if err != nil {
		fmt.Println("error")
	}

	currentAge := (x - y) / 10000

	fmt.Printf("%s who was born in %s would be %d years old today.\n", u.Name, u.City, currentAge)
}

func main() {
	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
	u.sayHello()
	u.sayHistory()
}
