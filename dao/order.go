package dao

import (
	"E-commerce/model"
	"context"

	"gorm.io/gorm"
)

// 数据访问对象
type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

// 创建新的订单
func (dao *OrderDao) CreateOrder(order *model.Order) error {
	return dao.DB.Model(&model.Order{}).Create(&order).Error
}

func (dao *OrderDao) GetOrderByCid(id, cId uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id = ? AND user_id = ?", id, cId).First(&order).Error
	return
}

func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (order []*model.Order, total int64, err error) {
	err = dao.DB.Model(&model.Order{}).
		Where(condition).Count(&total).
		Error
	if err != nil {
		return
	}
	err = dao.DB.Model(&model.Order{}).
		Where(condition).
		Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).
		Find(&order).Error
	return
}

func (dao *OrderDao) DeleteOrderByOrderId(aId, uId uint) error {
	return dao.DB.Model(&model.Order{}).Where("id = ? AND user_id = ?", aId, uId).Delete(&model.Order{}).Error
}
