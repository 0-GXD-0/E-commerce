package dao

import (
	"E-commerce/model"
	"context"

	"gorm.io/gorm"
)

type userDao struct {
	*gorm.DB
}

// 创建一个新的 userDao 实例，并将带有上下文的 GORM 数据库实例赋值给 userDao 的 DB 字段
// 这样在 userDao 中的所有数据库操作都会带有上下文信息
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

// 根据用户id获取用户信息
func (dao *userDao) GetUserById(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id = ?", id).First(&user).Error
	return
}

func (dao *userDao) UpdateUserById(uid uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id = ?", uid).Updates(&user).Error
}
