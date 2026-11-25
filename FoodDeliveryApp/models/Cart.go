package models

type Cart struct {
	Res      *Restaurant
	MenuList []MenuItem
	Total    float64
}

func (c *Cart) AddToCart(res *Restaurant, item MenuItem) {
	if c.Res != nil && c.Res.Id != res.Id {
		c.Clear()
	}
	c.Res = res
	c.MenuList = append(c.MenuList, item)
	c.Total += item.Price
}

func (c *Cart) Clear() {
	c.Res = nil
	c.MenuList = []MenuItem{}
	c.Total = 0

}
func (c *Cart) IsEmpty() bool {
	return len(c.MenuList) == 0
}
