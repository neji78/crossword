package user

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type UserRepository interface {
	Create(user *User) error
	FindByID(id int) (*User, error)
	FindByUsername(username string) (*User, error)
	Update(user *User) error
	Delete(id int) error
}