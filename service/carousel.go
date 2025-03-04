package service

import (
	"E-commerce/dao"
	"E-commerce/pkg/e"
	"E-commerce/pkg/util"
	"E-commerce/serializer"
	"context"
)

type CarouselService struct{}

func (service *CarouselService) List(ctx context.Context) serializer.Response {
	carouselDao := dao.NewCarouselDao(ctx)
	code := e.Success
	carousels, err := carouselDao.ListCarousel()
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}
