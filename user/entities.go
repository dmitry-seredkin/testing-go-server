package user

import "gopkg.in/guregu/null.v4"

type User struct {
	Id    string      `json:"id"`
	Login string      `json:"login"`
	Name  string      `json:"name"`
	Email null.String `json:"email"`
	Phone null.String `json:"phone"`
}

type UserItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateUser struct {
	Login    string      `json:"login"`
	Name     string      `json:"name"`
	Password string      `json:"password"`
	Email    null.String `json:"email,omitempty"`
	Phone    null.String `json:"phone,omitempty"`
}

type UpdateUser struct {
	Name  null.String `json:"name"`
	Email null.String `json:"email,omitempty"`
	Phone null.String `json:"phone,omitempty"`
}

type LoginUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
