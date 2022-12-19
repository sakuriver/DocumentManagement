package order_math_model

// SaleSlipData 購入された時の売上伝票
type SaleSlipData struct {
	// 購入者の名前
	SalerName string
	// 購入年月日
	SaleDate string
	// 注文された時の商品一覧
	Products []ProductInfo
}

// ProductInfo 商品データ
type ProductInfo struct {
	// 商品やコンテンツs名
	ProductName string
	// 商品のカートに入れた数
	ProductLen int
	// 商品単価
	Amount int
	// 商品勝った時のメモ
	Note string
}
