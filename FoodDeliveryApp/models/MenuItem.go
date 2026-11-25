package models

type MenuItem struct {
	Code  int
	Name  string
	Price float64
}

func NewMenuItem(code int, name string, price float64) *MenuItem {
	return &MenuItem{
		Code:  code,
		Name:  name,
		Price: price,
	}
}
