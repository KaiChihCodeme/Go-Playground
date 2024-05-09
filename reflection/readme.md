# Reflect

Go為靜態類型，也就是在編譯期只能有一個確定並已知的類型。所以在declaire之後就不能再更改。
雖然在運行之後可以賦予他不同類型的值，蛋改變的只是他內部的動態類型與動態值。

`reflect.TypeOf()` 用來取出interface的動態類型，return type `reflect.Type`
`reflect.ValueOf()` 則是取出interface的動態值，return type `reflect.Value`
`reflect.Kind()` 獲取最相似之種類
`reflect.Value.NumField()` 從struct reflect拿字段個數 (struct有幾個type)
`reflect.Value.Field(i)` 從struct reflect拿第i個reflect對象 (第i個type)
`reflect.Int()/reflect.String()...` 從reflect對象取出具體類型

要注意， `NumField()` and `Field()` 只能使用在struct, 否則會 `panic`