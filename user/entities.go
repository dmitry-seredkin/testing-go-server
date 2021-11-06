package user

import "gopkg.in/guregu/null.v4"

type User struct {
	Id    string      `json:"id"`
	Name  string      `json:"name"`
	Email null.String `json:"email"`
	Phone null.String `json:"phone"`
}
