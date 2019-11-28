package viewmodel
import (
	"04-looping/model"
)
type ShopDetail struct {
	Title    string
	Active   string
	Products []Product
}

func NewShopDetail(products []model.Product) ShopDetail {
	var result = ShopDetail{
		Title: "Lemonde Stand Supply",
		Active: "shop",
		Products: []Product{},
	}
	
	for _,p := range products {
		result.Products = append(result.Products, productToVM(p))
	}
	return result
}
