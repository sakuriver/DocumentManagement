# 編集中

## 売上サンプル翻訳してみた(business program sample)

## コメント欄に記載されている、システム背景より抜粋
```mermaid

classDiagram
    class ServiceSaleTransaction
    ServiceSaleTransaction : +String sale-name
    ServiceSaleTransaction : +String sexual
    ServiceSaleTransaction : +deposit(amount) age
    ServiceSaleTransaction : +deposit(amount) sale



```

## コンビニ売上や、ネット決済などのサービス名が追加される(sale payment service name add tax or business)
```mermaid

classDiagram
    class ServiceTransactionMust
    ServiceTransactionMust : +String sale-name
    ServiceTransactionMust : +String sexual
    ServiceTransactionMust : +deposit(amount) ages
    ServiceTransactionMust : +deposit(amount) sale
    ServiceTransactionMust : +String service-name 


```
## 売上、消費税の計算結果時のメッセージ一覧(messages)


|  メッセージID(MSGID)  |  メッセージ(Message)  |
| ---- | ---- |
|  KFSB05200-E |  SERVER:領域確保に失敗しました。  |  TD  |  TD  |
|  KFSB05300-E | SERVER:タイマトランザクション起動に失敗しました。 | 
|  KFSB05400-E | SERVER:ROLLBACKに失敗しました。 | 
|  KFSB06000-E | SERVER:DBキューのメッセージ書込みに失敗しました。 | 


## データ変換処理(cash register , e money ) → digital transaction data
```mermaid
sequenceDiagram
    participant ConvertPc
    ConvertPc->>ConvertPc: input name to database name
    ConvertPc->>ConvertPc: input sexual to database sexual    
    ConvertPc->>ConvertPc: input age to database age    
    ConvertPc->>ConvertPc: input sale to database sale
    ConvertPc->>ConvertPc: input servicename to database servicename

```

## データ巻き戻し(Rollback)
```mermaid
flowchart TB
    messageerr-->errorset
    errorset-->arg1set
    arg1set-->arg2set
    messageok-->errorend
```

## 保存処理への送信(send to message)
```mermaid
flowchart TB
    start-->senddatasetting
    senddatasetting-->flagset
    flagset-->putdatalen
    putdata-->putdatabody
```

## データ発行
```mermaid
flowchart TB
    start-->messagesend
    messagesend-->send-resultcheck
    messagesend-->send-rollback
    send-resultcheck-->append
    send-rollback-->append

```

## 集計用コンピュータの呼び出し(Recmote Procedure Call)
```mermaid
flowchart TB
    strart-->indata-setting
    indata-setting-->indata-lensetting
    indata-lensetting-->rpc-flag-set
    rpc-flag-set-->rpc-svname
    rpc-svname-->rpc-svgroup
    rpc-svgroup-->rpc-data-len
    rpc-data-len-->call-rpc
    call-rpc-->send-resultcheck
    call-rpc-->send-rollback
    send-resultcheck-->append
    send-rollback-->append
```



