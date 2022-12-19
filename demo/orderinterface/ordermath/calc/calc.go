package order_math_calc

import (
	order_math_model "../model"
)

//CalcTotalPrice 商品の合計金額を計算
func CalcTotalPrice(products []order_math_model.ProductInfo) float64 {
	totalPrice := 0.0
	for _, product := range products {
		totalPrice += float64(product.Amount * product.ProductLen)
	}
	return totalPrice
}
