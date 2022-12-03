package main

import "fmt"

const taxRate = 0.1

// 商品売上系の試し中
func main() {

	// 売上伝票のテスト用データ
	saleSlipData := SaleSlipData{
		SalerName: "佐藤　一郎",
		SaleDate:  "2022年10月8日",
		Products: []ProductInfo{
			ProductInfo{
				ProductName: "コンサートチケット 団体1式",
				ProductLen:  10,
				Amount:      2000,
				Note:        "団体購入で一式購入者の名前です",
			},
			ProductInfo{
				ProductName: "呪術廻戦 21",
				ProductLen:  2,
				Amount:      550,
				Note:        "旅行の暇つぶし用です",
			},
		},
	}

	displaySaleSlipHeader(saleSlipData.SalerName, saleSlipData.SaleDate)

	totalPrice := productDetailDisplay(saleSlipData.Products)

	displaySaleSlipFooter(totalPrice)
}

//displaySaleSlipHeader 売上伝票の上部を表示する
func displaySaleSlipHeader(salerName string, saleDate string) {
	println("売上伝票情報")
	println(fmt.Sprintf("購入者名 %s 購入日 %s", salerName, saleDate))
	println("")
	println("")
	println("")

	println("商品一覧")

}

func displaySaleSlipFooter(totalPrice float64) {
	//
	println("")
	println("")
	println("")
	println(fmt.Sprintf("商品合計金額 %d円", (int64)(totalPrice)))
	println(fmt.Sprintf("消費税 %d円", (int64)(totalPrice*taxRate)))
	println(fmt.Sprintf("合計金額 %d円", (int64)(totalPrice+(totalPrice*taxRate))))

}

// 商品の詳細画面を表示する
func productDetailDisplay(products []ProductInfo) float64 {
	totalPrice := 0.0
	for _, product := range products {
		println(fmt.Sprintf("商品名 %s 個数 %d 金額 %d 備考 %s", product.ProductName, product.ProductLen, product.Amount*product.ProductLen, product.Note))
		totalPrice += float64(product.Amount * product.ProductLen)
	}
	return totalPrice
}

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
