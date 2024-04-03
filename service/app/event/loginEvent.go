package event

import "github.com/sonhineboy/gsadmin/service/app/models"

type LoginEvent struct {
	Name string
	User models.AdminUser
}

func NewLoginEvent(name string, user models.AdminUser) *LoginEvent {
	return &LoginEvent{
		Name: name,
		User: user,
	}
}

func (t LoginEvent) GetEventName() string {
	return "loginEvent"
}
