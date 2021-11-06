package user

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}
