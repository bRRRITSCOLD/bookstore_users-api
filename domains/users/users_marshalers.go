package users_domain

import (
	"encoding/json"
	"time"
)

type PublicUser struct {
	UserID int64 `json:"userId" mysql:"id"`
	// FirstName   string    `json:"firstName" mysql:"firstName"`
	// LastName    string    `json:"lastName" mysql:"lastName"`
	// Email       string    `json:"email" mysql:"email"`
	DateCreated time.Time `json:"dateCreated" mysql:"dateCreated"`
	Status      string    `json:"status"`
	// Password    string    `json:"password"`
}

type PrivateUser struct {
	UserID      int64     `json:"userId" mysql:"id"`
	FirstName   string    `json:"firstName" mysql:"firstName"`
	LastName    string    `json:"lastName" mysql:"lastName"`
	Email       string    `json:"email" mysql:"email"`
	DateCreated time.Time `json:"dateCreated" mysql:"dateCreated"`
	Status      string    `json:"status"`
	// Password    string    `json:"password"`
}

func (users Users) Marshal(isPublic bool) interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshal(isPublic)
	}
	return result
}

func (user *User) Marshal(isPublic bool) interface{} {
	userJson, _ := json.Marshal(user)

	if isPublic {
		var publicUser PublicUser
		json.Unmarshal(userJson, &publicUser)
		return publicUser
	}

	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
