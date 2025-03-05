package service

import (
	"E-commerce/dao"
	"E-commerce/pkg/e"
	"E-commerce/pkg/util"
	"E-commerce/serializer"
	"context"
)

type CategoryService struct{}

func (service *CategoryService) List(ctx context.Context) serializer.Response {
	categoryDao := dao.NewCategoryDao(ctx)
	code := e.Success
	category, err := categoryDao.ListCategory()
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategories(category), uint(len(category)))
}
