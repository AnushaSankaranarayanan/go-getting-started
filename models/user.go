package models

import (
	"errors"
	"fmt"
)

// User user
type User struct {
	ID        int
	Firstname string
	Lastname  string
}

var (
	users  []*User
	nextID = 1
)

//GetUsers GetUsers
func GetUsers() []*User {
	return users
}

//AddUser AddUser
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user cannot have id")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil

}

//GetUserByID  GetUserById
func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}

	}
	return User{}, fmt.Errorf("User with id %v not found", id)
}

//UpdateUser  UpdateUser
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}

	}
	return User{}, fmt.Errorf("User with id %v not found", u.ID)
}

//DeleteUserByID  DeleteUserByID
func DeleteUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with id %v not found", id)
}
