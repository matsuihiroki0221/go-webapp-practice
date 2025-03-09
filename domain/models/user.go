package models

type User struct {
	ID      int
	Auth0ID string
	Name    string
}

func NewUser(auth0Id, name string) *User {
	return &User{Auth0ID: auth0Id, Name: name}
}
