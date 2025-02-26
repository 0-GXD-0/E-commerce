package dao

import (
	"E-commerce/model"
	"context"

	"gorm.io/gorm"
)

type userDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *userDao {
	return &userDao{DB: NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *userDao {
	return &userDao{DB: db}
}

// 根据username判断用户是否存在
func (dao *userDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).Find(&user).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	return user, true, nil
}

func (dao *userDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}
