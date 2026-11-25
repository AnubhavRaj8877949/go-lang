package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
	City  string
}

type UserBuilder struct {
	user User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (b *UserBuilder) SetName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) SetEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) SetAge(age int) *UserBuilder {
	b.user.Age = age
	return b
}

func (b *UserBuilder) SetCity(city string) *UserBuilder {
	b.user.City = city
	return b
}

func (b *UserBuilder) Build() User {
	return b.user
}

func main() {
	user := NewUserBuilder().
		SetName("Rahul").
		SetEmail("rahul@gmail.com").
		SetAge(25).
		Build()

	fmt.Println(user)

}
