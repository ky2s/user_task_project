package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserViews struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func SchemaPublic(tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table("public" + "." + tableName)
	}
}

type UserModels interface {
	CreateUser(user Users) (Users, error)
	GetUserRows(user Users) ([]Users, error)
	GetUserRow(user Users) (Users, error)
	UpdateUser(id int, fields Users) (int, error)
	DeleteUser(id int) (bool, error)
}

type connection struct {
	db *gorm.DB
}

func NewUserModels(dbg *gorm.DB) UserModels {
	return &connection{
		db: dbg,
	}
}

func (con *connection) CreateUser(data Users) (Users, error) {
	err := con.db.Scopes(SchemaPublic("users")).Create(&data).Error
	if err != nil {
		fmt.Println(err)
		return Users{}, err
	}
	return data, nil
}

func (con *connection) GetUserRows(fields Users) ([]Users, error) {

	var data []Users
	err := con.db.Find(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (con *connection) GetUserRow(fields Users) (Users, error) {

	var data Users
	err := con.db.Where(fields).First(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Users{}, nil
	}

	if err != nil {
		return Users{}, err
	}

	return data, err
}

func (con *connection) UpdateUser(userID int, fields Users) (int, error) {

	err := con.db.Scopes(SchemaPublic("users")).Where("id = ?", userID).Updates(fields).Error
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return userID, nil
}

func (con *connection) DeleteUser(id int) (bool, error) {
	var data Users
	err := con.db.Scopes(SchemaPublic("users")).Where("id = ?", id).Delete(&data).Error
	if err != nil {
		return false, err
	}

	return true, err
}
