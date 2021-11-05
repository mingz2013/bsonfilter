# bsonfilter
用于从mongo的bson备份文件中查询相应的数据，简化备份文件。

使用场景，mongo的全量备份和增量备份都是bson格式的数据。当备份文件很大，且只需要恢复部分数据的时候，可以使用本工具对备份文件进行简化，以达到缩减时间的目的。


## help
```
$ ./bin/bsonfilter.darwin 
Usage of ./bin/bsonfilter.darwin:
  -h    show help
  -i string
        input file
  -o string
        output file
  -q string
        query filter, as a json string
  -v    show details info
  -vv
        show debug info 
```


## example
```shell
./bin/bsonfilter.darwin -i backup/data/data/data.bson -o backup/data_new/data/tmp.bson -q '{"_id": 1}'

```

## -q参数说明
-q参数 用来传递 mongo的query 语句。由于此部分功能没有找到现有实现，由本工程自己实现。但是支持不全面。

## 目前支持的指令

mongo全部指令参考：https://docs.mongodb.com/v4.0/reference/operator/query/#query-selectors


- $eq
- $ne
- $gt
- $gte
- $lt
- $lte
- $in
- $nin
- $and
- $not
- $nor
- $or
- $exists

## 目前支持的类型比较

全部类型参考：https://github.com/mongodb/mongo-go-driver/blob/master/bson/types.go

```go
const (
	TypeDouble           = bsontype.Double
	TypeString           = bsontype.String
	TypeEmbeddedDocument = bsontype.EmbeddedDocument
	TypeArray            = bsontype.Array
	TypeBinary           = bsontype.Binary
	TypeUndefined        = bsontype.Undefined
	TypeObjectID         = bsontype.ObjectID
	TypeBoolean          = bsontype.Boolean
	TypeDateTime         = bsontype.DateTime
	TypeNull             = bsontype.Null
	TypeRegex            = bsontype.Regex
	TypeDBPointer        = bsontype.DBPointer
	TypeJavaScript       = bsontype.JavaScript
	TypeSymbol           = bsontype.Symbol
	TypeCodeWithScope    = bsontype.CodeWithScope
	TypeInt32            = bsontype.Int32
	TypeTimestamp        = bsontype.Timestamp
	TypeInt64            = bsontype.Int64
	TypeDecimal128       = bsontype.Decimal128
	TypeMinKey           = bsontype.MinKey
	TypeMaxKey           = bsontype.MaxKey
)
```

目前的类型支持情况如下，

没有列出的默认不支持, 在query时默认为不匹配

| bson文档中的类型 | query语句中的类型 | 处理方案 |
|:---:|:---:|:---:|
| double | double | double比较 |
| double | int32 | float64 |
| double | int64 | float64 |
| string | string | bytes |
| ObjectId | ObjectId | []byte(Hex()) |
| DateTime | DateTime | DateTime |
| Int32 | Double | float64 |
| Int32 | Int32 | Int32 |
| Int32 | Int64 | Int64 |
| Timestamp | Timestamp | Timestamp |
| Int64 | Double | float64 |
| Int64 | Int32 | Int64 |
| Int64 | Int64 | Int64 |




# ref

