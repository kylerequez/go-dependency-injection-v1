package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/kylerequez/go-dependency-injection/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllUsers() ([]models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(dsn string) *UserRepository {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("ERROR IN CREATING USER REPOSITORY:::-:::", err)
		return nil
	}
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetAllUsers() (users []models.User, err error) {
	var allUsers []models.User
	result := ur.db.Find(&allUsers)
	if result.Error != nil {
		fmt.Println("ERROR:::-:::", result.Error)
		return nil, errors.New("error in get users")
	}
	return allUsers, nil
}
