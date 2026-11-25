package models

type Restaurant struct {
	Id       int
	Name     string
	Address  string
	MenuList []MenuItem
}

func NewRestaurant(name, address string) *Restaurant {
	return &Restaurant{
		Name:     name,
		Address:  address,
		MenuList: []MenuItem{},
	}
}

func (r *Restaurant) AddMenuItem(item MenuItem) {
	r.MenuList = append(r.MenuList, item)
}

func (r *Restaurant) RemoveMenuItem(code int) {
	for i, item := range r.MenuList {
		if item.Code == code {
			r.MenuList = append(r.MenuList[:i], r.MenuList[i+1:]...)
			break
		}
	}
}

func (r *Restaurant) GetMenuItem(code int) *MenuItem {
	for _, item := range r.MenuList {
		if item.Code == code {
			return &item
		}
	}
	return nil
}

func (r *Restaurant) UpdateMenuItem(code int, name string, price float64) {
	for i := range r.MenuList {
		if r.MenuList[i].Code == code {
			r.MenuList[i].Name = name
			r.MenuList[i].Price = price
			break
		}
	}
}
