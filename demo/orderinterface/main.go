package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	ordermath "./ordermath/calc"
	order_math_define "./ordermath/define"
	order_math_model "./ordermath/model"
)

// 商品売上系の試し中
func main() {

	println("sale slip calc start")
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

	saleSlipData := order_math_model.SaleSlipData{}

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
			saleSlipData.Products = append(saleSlipData.Products, order_math_model.ProductInfo{
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

	// 合計金額を計算する
	totalPrice := ordermath.CalcTotalPrice(saleSlipData.Products)
	displaySaleSlipFooter(csvWriter, totalPrice)

	for {

	}
	println("sale slip calc end")

}

//displaySaleSlipHeader 売上伝票の上部を表示する
func displaySaleSlipHeader(csvWriter *csv.Writer, salerName string, saleDate string) {
	println("sale slip header start")
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
	csvWriter.Write([]string{fmt.Sprintf("消費税 %d円", (int64)(totalPrice*order_math_define.TaxRate))})
	csvWriter.Write([]string{fmt.Sprintf("商品合計金額 %d円", (int64)(totalPrice))})
	csvWriter.Write([]string{fmt.Sprintf("合計金額 %d円", (int64)(totalPrice+(totalPrice*order_math_define.TaxRate)))})
	csvWriter.Flush()

}

//displayProductDisplay 商品の詳細ぶぶんの表示と書き込みを表示する
func displayProductDetail(csvWriter *csv.Writer, products []order_math_model.ProductInfo) {
	csvWriter.Write([]string{"商品名", "個数", "金額", "備考"})
	for _, product := range products {
		csvWriter.Write([]string{product.ProductName, fmt.Sprintf("%d", product.ProductLen), fmt.Sprintf("%d", product.Amount*product.ProductLen), product.Note})
		println(fmt.Sprintf("%s,%d,%d,%s", product.ProductName, product.ProductLen, product.Amount*product.ProductLen, product.Note))
	}
	csvWriter.Flush()
}
