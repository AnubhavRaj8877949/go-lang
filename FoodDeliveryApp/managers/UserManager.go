package managers

import (
	"fooddelivery/models"
	"sync"
)

type UserManager struct {
	UserList []*models.User
	nextID   int
}

var (
	userManagerInstance *UserManager
	userOnce            sync.Once
)

func GetUserManager() *UserManager {
	userOnce.Do(func() {
		userManagerInstance = &UserManager{
			UserList: make([]*models.User, 0),
			nextID:   1,
		}
	})
	return userManagerInstance
}

func (um *UserManager) Add(user *models.User) {
	user.Id = um.nextID
	um.nextID++
	um.UserList = append(um.UserList, user)
}
