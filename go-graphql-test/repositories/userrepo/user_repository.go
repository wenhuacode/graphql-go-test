package userrepo

import (
	"go-graphql-test/domain/user"

	"gorm.io/gorm"
)

// Repo interface
type Repo interface {
	GetByID(id uint) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
	Create(user *user.User) error
	Update(user *user.User) error
	Delete(id uint) error
}

type userRepo struct {
	db *gorm.DB
}

func (u *userRepo) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}

// NewUserRepo will instantiate User Repository
func NewUserRepo(db *gorm.DB) Repo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) GetByID(id uint) (*user.User, error) {
	var usr user.User
	if err := u.db.First(&usr, id).Error; err != nil {
		return nil, err
	}
	return &usr, nil
}

func (u *userRepo) GetByEmail(email string) (*user.User, error) {
	var usr user.User
	if err := u.db.Where("email = ?", email).First(&usr).Error; err != nil {
		return nil, err
	}
	return &usr, nil
}

func (u *userRepo) Create(user *user.User) error {
	return u.db.Create(user).Error
}

func (u *userRepo) Update(user *user.User) error {
	return u.db.Save(user).Error
}
