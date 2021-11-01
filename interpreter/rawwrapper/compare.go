package rawwrapper

import (
	"github.com/mongodb/mongo-tools/common/log"
	"go.mongodb.org/mongo-driver/bson"
)

const eq = 0
const greater = 1
const less = -1

//const (
//	TypeDouble           = bsontype.Double
//	TypeString           = bsontype.String
//	TypeEmbeddedDocument = bsontype.EmbeddedDocument
//	TypeArray            = bsontype.Array
//	TypeBinary           = bsontype.Binary
//	TypeUndefined        = bsontype.Undefined
//	TypeObjectID         = bsontype.ObjectID
//	TypeBoolean          = bsontype.Boolean
//	TypeDateTime         = bsontype.DateTime
//	TypeNull             = bsontype.Null
//	TypeRegex            = bsontype.Regex
//	TypeDBPointer        = bsontype.DBPointer
//	TypeJavaScript       = bsontype.JavaScript
//	TypeSymbol           = bsontype.Symbol
//	TypeCodeWithScope    = bsontype.CodeWithScope
//	TypeInt32            = bsontype.Int32
//	TypeTimestamp        = bsontype.Timestamp
//	TypeInt64            = bsontype.Int64
//	TypeDecimal128       = bsontype.Decimal128
//	TypeMinKey           = bsontype.MinKey
//	TypeMaxKey           = bsontype.MaxKey
//)

func checkTypeSupport(value bson.RawValue) bool {
	switch value.Type {
	case bson.TypeInt32, bson.TypeInt64, bson.TypeDecimal128, bson.TypeDouble:
		return true
	default:
		log.Logvf(log.Always, "type support error %v", value.Type)
		return false
	}
}

func Compare(left, right bson.RawValue) int {
	// 0 ==
	// 1 >
	// -1 <
	log.Logvf(log.Always, "Compare: left %v, right %v", left, right)
	//if left.IsNumber() && right.IsNumber(){
	//	if left.Double() > right.Double(){
	//		return greater
	//	}else if left.Double() < right.Double(){
	//		return less
	//	}else {
	//		return eq
	//	}
	//} else {
	//return bytes.Compare(left.Value, right.Value)
	//}

	//if left.IsNumber() && right.IsNumber(){
	//	left.
	//}

	if checkTypeSupport(left) && checkTypeSupport(right) {
		return false
	}

	switch left.Type {
	case bson.TypeDouble:

	}

	return eq

}

func IsEqual(left, right bson.RawValue) bool {
	result := Compare(left, right)
	return result == eq
}

func IsLess(left, right bson.RawValue) bool {
	result := Compare(left, right)
	return result == less
}

func IsLessOrEq(left, right bson.RawValue) bool {
	result := Compare(left, right)
	return result == eq || result == less
}

func IsGreater(left, right bson.RawValue) bool {
	result := Compare(left, right)
	return result == greater
}

func IsGreaterOrEq(left, right bson.RawValue) bool {
	result := Compare(left, right)
	return result == eq || result == greater
}

func IsIn(left, right bson.RawValue) bool {
	//raw := right.Array()
	raw := right.Array()
	elems, err := raw.Values()
	if err != nil {
		log.Logvf(log.Always, "error %v", err)
	}
	for _, elem := range elems {
		if Compare(left, elem) == eq {
			return true
		}
	}
	return false

}
