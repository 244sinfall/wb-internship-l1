package main

import (
	"fmt"
	"math"
	"time"
)

type Human struct {
	firstName string
	lastName  string
	age       int
	email     string
}

func (h *Human) getFullName() string {
	return h.firstName + " " + h.lastName
}
func (h *Human) getLivedHours() string {
	birthday := time.Now().AddDate(-h.age, 0, 0)
	d := -birthday.Sub(time.Now())
	return fmt.Sprintf("%s lived %d hours", h.getFullName(), int(math.Round(d.Hours())))
}

type Action struct {
	*Human
}

func (a *Action) walk() {
	fmt.Println(a.getFullName() + " is walking...")
}

func task1() {
	dima := Human{
		firstName: "Dmitry",
		lastName:  "Filin",
		age:       24,
		email:     "dimafilin6@gmail.com",
	}
	fmt.Println(dima.getFullName())
	fmt.Println(dima.getLivedHours())
	dimaAction := Action{&dima}
	dimaAction.walk()
}
