package main

import (
	"errors"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func DisplayUser(u *user) {
	println("First Name: ", u.firstName)
	println("Last Name: ", u.lastName)
	println("Birth Date: ", u.birthDate)
}

// method with receiver of type user struct (value receiver)
func (u user) getUser() {
	println("First Name: ", u.firstName)
	println("Last Name: ", u.lastName)
	println("Birth Date: ", u.birthDate)
}

// mutation method with receiver of type user struct (pointer receiver)
func (u *user) clearUser() {
	// create a empty user struct and assign to the receiver
	*u = user{}

	// or assign empty values to the receiver
	// u.firstName = ""
	// u.lastName = ""
	// u.birthDate = ""
}

// constructor function
func NewUser(firstName, lastName, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("invalid user details")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
	}, nil
}

func main() {

	var appUser user = user{
		firstName: "John",
		lastName:  "Doe",
		birthDate: "01/01/1990",
		createdAt: time.Now(),
	}

	// Accessing struct fields
	println(appUser.firstName)
	println(appUser.lastName)

	DisplayUser(&appUser)
	appUser.getUser()
	appUser.clearUser()
	appUser.getUser()

	myAppUser, err := NewUser("", "", "01/01/1991")
	if err != nil {
		println(err.Error())
		return
	}
	DisplayUser(myAppUser)
}
