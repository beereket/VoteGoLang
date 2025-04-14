package repositories

import (
	"VoteGolang/internals/data"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *data.User) error
	GetByID(id string) (*data.User, error)
	GetByUsername(username string) (*data.User, error)
	Update(user *data.User) error
	Delete(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *data.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id string) (*data.User, error) {
	var user data.User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(username string) (*data.User, error) {
	var user data.User
	err := r.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *data.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id string) error {
	return r.db.Delete(&data.User{}, "id = ?", id).Error
}
