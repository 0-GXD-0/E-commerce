package service

import (
	"E-commerce/dao"
	"E-commerce/serializer"
	"context"
	"strconv"
)

type ListProductImg struct{}

func (service *ListProductImg) List(ctx context.Context, pId string) serializer.Response {
	productImgDao := dao.NewProductImgDao(ctx)
	productId, _ := strconv.Atoi(pId)
	productImgs, _ := productImgDao.ListProductImg(uint(productId))
	return serializer.BuildListResponse(serializer.BuildProductImgs(productImgs), uint(len(productImgs)))
}
