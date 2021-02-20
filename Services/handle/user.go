package handle

import "Services/models"

type UserData struct {
	data []models.User
}

func NewUser() *UserData {
	return &UserData{}
}

func (c *UserData) AllData() []models.User {
	return c.data
}
