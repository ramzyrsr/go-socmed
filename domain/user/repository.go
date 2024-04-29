package user

type UserRepository interface {
	Create(user *User) error
	GetByID(id int) (*User, error)
	Update(user *User) error
	Delete(id int) error
}
