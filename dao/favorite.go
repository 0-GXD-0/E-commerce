package dao

import (
	"E-commerce/model"
	"context"

	"gorm.io/gorm"
)

// 数据访问对象
type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func NewFavoriteDaoByDB(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}

// 根据id获取favorite
func (dao *FavoriteDao) ListFavorite(uId uint) (favorite []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id = ?", uId).Find(&favorite).Error
	return
}

// 检查收藏夹是否已经存在
func (dao *FavoriteDao) FavoriteExistOrNot(pId, uId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("product_id = ? AND user_id = ?", pId, uId).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, err
	}
	return true, nil
}

// 创建新的收藏夹
func (dao *FavoriteDao) CreateFavorite(in *model.Favorite) error {
	return dao.DB.Model(&model.Favorite{}).Create(&in).Error
}

// 删除收藏夹记录
func (dao *FavoriteDao) DeleteFavorite(uId, fId uint) error {
	return dao.DB.Model(&model.Favorite{}).
		Where("id = ? AND user_id= ?", fId, uId).Delete(&model.Favorite{}).Error
}
