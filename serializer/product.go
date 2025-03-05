package serializer

import (
	"E-commerce/conf"
	"E-commerce/model"
)

type Product struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	Category      uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          uint64 `json:"view"`
	CreateAt      int64  `json:"create_at"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

func BuildProduct(item *model.Product) Product {
	return Product{
		Id:            item.ID,
		Name:          item.Name,
		Category:      item.Category,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		View:          uint64(item.View()),
		CreateAt:      item.CreatedAt.Unix(),
		Num:           item.Num,
		OnSale:        item.OnSale,
		BossId:        item.BossId,
		BossName:      item.BossName,
		BossAvatar:    conf.Host + conf.HttpPort + conf.AvatarPath + item.BossAvatar,
	}
}

func BuildProducts(items []*model.Product) (products []Product) {
	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, product)
	}
	return products
}
