package main

import (
	"fmt"
)

type AuthModule struct {
	users  map[int]*User
	lastID int
}

type User struct {
	Login string
	Pass  string
}

func NewAuthModule() *AuthModule {
	return &AuthModule{lastID: -1, users: make(map[int]*User)}
}

func (am *AuthModule) AddUser(login string, pass string) {
	am.users[am.lastID+1] = &User{login, pass}
	am.lastID++
}

func (am *AuthModule) GetUser(id int) (*User, error) {
	user, exists := am.users[id]

	if !exists {
		return nil, fmt.Errorf("user %d does not exist", id)
	}

	return user, nil
}

func (am *AuthModule) UpdateUser(id int, login string, pass string) error {
	user, err := am.GetUser(id)

	if err != nil {
		return err
	}

	user.Login = login
	user.Pass = pass

	return nil
}

func (am *AuthModule) DeleteUser(id int) error {
	_, err := am.GetUser(id)
	if err != nil {
		return err
	}
	delete(am.users, id)

	return nil
}

func (am *AuthModule) PrintUsers() {
	for id, user := range am.users {
		fmt.Printf("User #%d is %s\n", id, user.Login)
	}
}

func main() {
	authModule := NewAuthModule()

	authModule.AddUser("Bilbo", "#Baggins95$")
	authModule.AddUser("Frodo", "~myPrecious~")

	fmt.Println("Users list after authModule init: ")
	authModule.PrintUsers()

	err := authModule.UpdateUser(1, "Gorlum", "~myPreciousssssss~")

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println("Users list after user ID#1 update: ")
	authModule.PrintUsers()

	err = authModule.DeleteUser(0)

	if err != nil {
		println(err.Error())
	}

	fmt.Println("Users list after user ID#0 delete: ")
	authModule.PrintUsers()
}
