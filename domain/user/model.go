package user

type User struct {
	Id       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required,email"`
	PassHash string `json:"password" db:"password" binding:"required"`
}
