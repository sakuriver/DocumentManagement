# 目的

## 各パッケージのモジュール仕様書

### mainパッケージ

|  列番号(1始まり)  |  関数名 |  概要  |  パラメータ引数  |  戻り値  |
| ---- | ---- | ---- | ---- | ---- |
|  1  |  main  | go run コマンド時の初期実行関数   |  なし  |  なし  |
|  2  |  writeSaleSlipHeader  |  売上伝票のヘッダー情報を出力する  |  csvWriter *csv.Writer  |  なし  |
|  3  |  displayProductDisplay  |  商品詳細情報の価格を出力する  |  csvWriter *csv.Writer, products []  |  なし  |
|  4  |  displaySaleSlipFooter  |  売上伝票のフッター情報を出力する  |  csvWriter *csv.Writer, totalPrice float64  |  なし  |

### ordermath(商品金額関連)パッケージ

#### 関数一覧

|  列番号(1始まり)  |  関数名  |  概要  |  パラメータ引数  |  戻り値  |
| ---- | ---- | ---- | ---- | ---- |
|  1  |  calcTotalPrice  |  商品の合計金額を計算して出力する  |  ProductInfoの配列  |  float  |
|  2  |    |    |    |    |
|  3  |    |    |    |    |
|  4  |    |    |    |    |

#### 計算式定義

|  計算式名称  |  概要  |  計算式  |
| ---- | ---- | ---- |
|  1注文ないの商品金額  |  商品１行当たりの金額  |  商品.単価*商品.個数  |


