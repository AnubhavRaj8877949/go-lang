package models

type User struct {
	Id       int
	Name     string
	Address  string
	Email    string
	MobileNo string
	Cart     *Cart
}

func NewUser(Name string, Address string, MobileNo, Email string) *User {
	return &User{
		Name:     Name,
		Address:  Address,
		Email:    Email,
		MobileNo: MobileNo,
		Cart:     &Cart{},
	}
}
