package managers

import (
	"fooddelivery/models"
	"strings"
	"sync"
)

type RestaurantManager struct {
	RestaurantList []*models.Restaurant
	nextID         int
}

var (
	restManagerInstance *RestaurantManager
	restOnce            sync.Once
)

func GetRestaurantManager() *RestaurantManager {
	restOnce.Do(func() {
		restManagerInstance = &RestaurantManager{
			RestaurantList: []*models.Restaurant{},
			nextID:         1,
		}
	})
	return restManagerInstance
}

func (rm *RestaurantManager) Add(res *models.Restaurant) {
	res.Id = rm.nextID
	rm.nextID++
	rm.RestaurantList = append(rm.RestaurantList, res)
}
func (rm *RestaurantManager) Get(id int) *models.Restaurant {
	for idx := range rm.RestaurantList {
		if rm.RestaurantList[idx].Id == id {
			return rm.RestaurantList[idx]
		}
	}
	return nil
}
func (rm *RestaurantManager) Update(id int, name, address string) {
	if r := rm.Get(id); r != nil {
		r.Name = name
		r.Address = address
	}
}

func (rm *RestaurantManager) Remove(id int) {
	for i, r := range rm.RestaurantList {
		if r.Id == id {
			rm.RestaurantList = append(rm.RestaurantList[:i], rm.RestaurantList[i+1:]...)
			break
		}
	}
}

func (rm *RestaurantManager) SearchByLocation(location string) []*models.Restaurant {
	var result []*models.Restaurant
	for _, r := range rm.RestaurantList {
		if strings.EqualFold(r.Address, location) {
			result = append(result, r)
		}
	}
	return result
}
