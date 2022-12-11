package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

const taxRate = 0.1

// 商品売上系の試し中
func main() {

	file, err := os.Open("./sale.csv")
	if err != nil {
		println("file open error")
		return
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		println(fmt.Sprintf("csv data read error %s", err.Error()))
		return
	}

	saleSlipData := SaleSlipData{}

	for record := range records {
		// 販売者のヘッダー情報を取得
		cols := records[record]
		saleSlipData.SalerName = cols[0]
		saleSlipData.SaleDate = cols[1]
		// 販売者の商品個別売り上げを読み込み
		productFile, err := os.Open("./product.csv")
		if err != nil {
			println("product open error")
			return
		}
		productR := csv.NewReader(productFile)
		productRecords, err := productR.ReadAll()
		if err != nil {
			println(fmt.Sprintf("csv data read error %s", err.Error()))
			return
		}

		// 商品情報の売上結果から取得する
		for productRecord := range productRecords {
			productLenI, _ := strconv.Atoi(productRecords[productRecord][1])
			productAmountI, _ := strconv.Atoi(productRecords[productRecord][2])
			saleSlipData.Products = append(saleSlipData.Products, ProductInfo{
				ProductName: productRecords[productRecord][0],
				ProductLen:  productLenI,
				Amount:      productAmountI,
				Note:        productRecords[productRecord][3],
			})

		}

	}

	// 売上伝票のテスト用データ

	resultFile, err := os.Create("./file_result.csv")
	defer resultFile.Close()

	csvWriter := csv.NewWriter(resultFile)
	displaySaleSlipHeader(csvWriter, saleSlipData.SalerName, saleSlipData.SaleDate)

	totalPrice := productDetailDisplay(csvWriter, saleSlipData.Products)
	displaySaleSlipFooter(csvWriter, totalPrice)

	for {

	}

}

//displaySaleSlipHeader 売上伝票の上部を表示する
func displaySaleSlipHeader(csvWriter *csv.Writer, salerName string, saleDate string) {
	headers := []string{"売上伝票情報"}

	csvWriter.Write(headers)
	csvWriter.Write([]string{fmt.Sprintf("購入者名 %s 購入日 %s", salerName, saleDate)})
	csvWriter.Write([]string{""})
	csvWriter.Write([]string{""})
	csvWriter.Flush()

}

//displaySaleSlipFooter 売上伝票のフッターを表示する
func displaySaleSlipFooter(csvWriter *csv.Writer, totalPrice float64) {
	//
	csvWriter.Write([]string{""})
	csvWriter.Write([]string{""})
	csvWriter.Write([]string{""})
	csvWriter.Write([]string{fmt.Sprintf("商品合計金額 %d円", (int64)(totalPrice))})
	csvWriter.Write([]string{fmt.Sprintf("消費税 %d円", (int64)(totalPrice*taxRate))})
	csvWriter.Write([]string{fmt.Sprintf("商品合計金額 %d円", (int64)(totalPrice))})
	csvWriter.Write([]string{fmt.Sprintf("合計金額 %d円", (int64)(totalPrice+(totalPrice*taxRate)))})
	csvWriter.Flush()

}

//productDetailDisplay 商品の詳細ぶぶんの表示と書き込みを表示する
func productDetailDisplay(csvWriter *csv.Writer, products []ProductInfo) float64 {
	totalPrice := 0.0
	csvWriter.Write([]string{"商品名", "個数", "金額", "備考"})
	for _, product := range products {
		csvWriter.Write([]string{product.ProductName, fmt.Sprintf("%d", product.ProductLen), fmt.Sprintf("%d", product.Amount*product.ProductLen), product.Note})
		println(fmt.Sprintf("%s,%d,%d,%s", product.ProductName, product.ProductLen, product.Amount*product.ProductLen, product.Note))
		totalPrice += float64(product.Amount * product.ProductLen)
	}
	csvWriter.Flush()
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
