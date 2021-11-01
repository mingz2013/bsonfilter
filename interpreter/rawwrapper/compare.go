package rawwrapper

import (
	"bytes"
	"github.com/mongodb/mongo-tools/common/log"
	"github.com/mongodb/mongo-tools/common/util"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

const eq = 0
const greater = 1
const less = -1
const notSupport = -2

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

//func checkTypeSupport(value bson.RawValue) bool {
//	switch value.Type {
//	case bson.TypeInt32, bson.TypeInt64, bson.TypeDecimal128, bson.TypeDouble:
//		return true
//	default:
//		log.Logvf(log.Always, "type support error %v", value.Type)
//		return false
//	}
//}

func compareDouble(left, right bson.RawValue) int {

	switch right.Type {
	case bson.TypeDouble:
		if left.Double() > right.Double() {
			return greater
		} else if left.Double() < right.Double() {
			return less
		} else {
			return eq
		}
	case bson.TypeString:
		return notSupport
	case bson.TypeEmbeddedDocument:
		return notSupport
	case bson.TypeArray:
		return notSupport
	case bson.TypeBinary:
		return notSupport
	case bson.TypeUndefined:
		return notSupport
	case bson.TypeObjectID:
		return notSupport
	case bson.TypeBoolean:
		return notSupport
	case bson.TypeDateTime:
		return notSupport
	case bson.TypeNull:
		return notSupport
	case bson.TypeRegex:
		return notSupport
	case bson.TypeDBPointer:
		return notSupport
	case bson.TypeJavaScript:
		return notSupport
	case bson.TypeSymbol:
		return notSupport
	case bson.TypeCodeWithScope:
		return notSupport
	case bson.TypeInt32:
		l := left.Double()
		r := float64(right.Int32())
		if l > r {
			return greater
		} else if l < r {
			return less
		} else {
			return eq
		}
	case bson.TypeTimestamp:
		return notSupport
	case bson.TypeInt64:
		l := left.Double()
		r := float64(right.Int64())
		if l > r {
			return greater
		} else if l < r {
			return less
		} else {
			return eq
		}
	case bson.TypeDecimal128:
		return notSupport

	case bson.TypeMinKey:
		return notSupport
	case bson.TypeMaxKey:
		return notSupport
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareString(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		return notSupport
	case bson.TypeString:
		return bytes.Compare(left.Value, right.Value)

	case bson.TypeEmbeddedDocument:
		return notSupport
	case bson.TypeArray:
		return notSupport
	case bson.TypeBinary:
		return notSupport
	case bson.TypeUndefined:
		return notSupport
	case bson.TypeObjectID:
		return notSupport
	case bson.TypeBoolean:
		return notSupport
	case bson.TypeDateTime:
		return notSupport
	case bson.TypeNull:
		return notSupport
	case bson.TypeRegex:
		return notSupport
	case bson.TypeDBPointer:
		return notSupport
	case bson.TypeJavaScript:
		return notSupport
	case bson.TypeSymbol:
		return notSupport
	case bson.TypeCodeWithScope:
		return notSupport
	case bson.TypeInt32:
		return notSupport
	case bson.TypeTimestamp:
		return notSupport
	case bson.TypeInt64:
		return notSupport
	case bson.TypeDecimal128:
		return notSupport
	case bson.TypeMinKey:
		return notSupport
	case bson.TypeMaxKey:
		return notSupport
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareEmbeddedDocument(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		return notSupport
	case bson.TypeString:
		return notSupport
	case bson.TypeEmbeddedDocument:
		return notSupport
	case bson.TypeArray:
		return notSupport
	case bson.TypeBinary:
		return notSupport
	case bson.TypeUndefined:
		return notSupport
	case bson.TypeObjectID:
		return notSupport
	case bson.TypeBoolean:
		return notSupport
	case bson.TypeDateTime:
		return notSupport
	case bson.TypeNull:
		return notSupport
	case bson.TypeRegex:
		return notSupport
	case bson.TypeDBPointer:
		return notSupport
	case bson.TypeJavaScript:
		return notSupport
	case bson.TypeSymbol:
		return notSupport
	case bson.TypeCodeWithScope:
		return notSupport
	case bson.TypeInt32:
		return notSupport
	case bson.TypeTimestamp:
		return notSupport
	case bson.TypeInt64:
		return notSupport
	case bson.TypeDecimal128:
		return notSupport
	case bson.TypeMinKey:
		return notSupport
	case bson.TypeMaxKey:
		return notSupport
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareArray(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		return notSupport
	case bson.TypeString:
		return notSupport
	case bson.TypeEmbeddedDocument:
		return notSupport
	case bson.TypeArray:
		return notSupport
	case bson.TypeBinary:
		return notSupport
	case bson.TypeUndefined:
		return notSupport
	case bson.TypeObjectID:
		return notSupport
	case bson.TypeBoolean:
		return notSupport
	case bson.TypeDateTime:
		return notSupport
	case bson.TypeNull:
		return notSupport
	case bson.TypeRegex:
		return notSupport
	case bson.TypeDBPointer:
		return notSupport
	case bson.TypeJavaScript:
		return notSupport
	case bson.TypeSymbol:
		return notSupport
	case bson.TypeCodeWithScope:
		return notSupport
	case bson.TypeInt32:
		return notSupport
	case bson.TypeTimestamp:
		return notSupport
	case bson.TypeInt64:
		return notSupport
	case bson.TypeDecimal128:
		return notSupport
	case bson.TypeMinKey:
		return notSupport
	case bson.TypeMaxKey:
		return notSupport
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareBinary(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		return notSupport
	case bson.TypeString:
		return notSupport
	case bson.TypeEmbeddedDocument:
		return notSupport
	case bson.TypeArray:
		return notSupport
	case bson.TypeBinary:
		return notSupport
	case bson.TypeUndefined:
		return notSupport
	case bson.TypeObjectID:
		return notSupport
	case bson.TypeBoolean:
		return notSupport
	case bson.TypeDateTime:
		return notSupport
	case bson.TypeNull:
		return notSupport
	case bson.TypeRegex:
		return notSupport
	case bson.TypeDBPointer:
		return notSupport
	case bson.TypeJavaScript:
		return notSupport
	case bson.TypeSymbol:
		return notSupport
	case bson.TypeCodeWithScope:
		return notSupport
	case bson.TypeInt32:
		return notSupport
	case bson.TypeTimestamp:
		return notSupport
	case bson.TypeInt64:
		return notSupport
	case bson.TypeDecimal128:
		return notSupport
	case bson.TypeMinKey:
		return notSupport
	case bson.TypeMaxKey:
		return notSupport
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareUndefined(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		return notSupport
	case bson.TypeString:
		return notSupport
	case bson.TypeEmbeddedDocument:
		return notSupport
	case bson.TypeArray:
		return notSupport
	case bson.TypeBinary:
		return notSupport
	case bson.TypeUndefined:
		return notSupport
	case bson.TypeObjectID:
		return notSupport
	case bson.TypeBoolean:
		return notSupport
	case bson.TypeDateTime:
		return notSupport
	case bson.TypeNull:
		return notSupport
	case bson.TypeRegex:
		return notSupport
	case bson.TypeDBPointer:
		return notSupport
	case bson.TypeJavaScript:
		return notSupport
	case bson.TypeSymbol:
		return notSupport
	case bson.TypeCodeWithScope:
		return notSupport
	case bson.TypeInt32:
		return notSupport
	case bson.TypeTimestamp:
		return notSupport
	case bson.TypeInt64:
		return notSupport
	case bson.TypeDecimal128:
		return notSupport
	case bson.TypeMinKey:
		return notSupport
	case bson.TypeMaxKey:
		return notSupport
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareObjectID(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		return notSupport
	case bson.TypeString:
		return notSupport
	case bson.TypeEmbeddedDocument:
		return notSupport
	case bson.TypeArray:
		return notSupport
	case bson.TypeBinary:
		return notSupport
	case bson.TypeUndefined:
		return notSupport
	case bson.TypeObjectID:
		return notSupport
	case bson.TypeBoolean:
		return notSupport
	case bson.TypeDateTime:
		return notSupport
	case bson.TypeNull:
		return notSupport
	case bson.TypeRegex:
		return notSupport
	case bson.TypeDBPointer:
		return notSupport
	case bson.TypeJavaScript:
		return notSupport
	case bson.TypeSymbol:
		return notSupport
	case bson.TypeCodeWithScope:
		return notSupport
	case bson.TypeInt32:
		return notSupport
	case bson.TypeTimestamp:
		return notSupport
	case bson.TypeInt64:
		return notSupport
	case bson.TypeDecimal128:
		return notSupport
	case bson.TypeMinKey:
		return notSupport
	case bson.TypeMaxKey:
		return notSupport
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareBoolean(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		return notSupport
	case bson.TypeString:
		return notSupport
	case bson.TypeEmbeddedDocument:
		return notSupport
	case bson.TypeArray:
		return notSupport
	case bson.TypeBinary:
		return notSupport
	case bson.TypeUndefined:
		return notSupport
	case bson.TypeObjectID:
		return notSupport
	case bson.TypeBoolean:
		return notSupport
	case bson.TypeDateTime:
		return notSupport
	case bson.TypeNull:
		return notSupport
	case bson.TypeRegex:
		return notSupport
	case bson.TypeDBPointer:
		return notSupport
	case bson.TypeJavaScript:
		return notSupport
	case bson.TypeSymbol:
		return notSupport
	case bson.TypeCodeWithScope:
		return notSupport
	case bson.TypeInt32:
		return notSupport
	case bson.TypeTimestamp:
		return notSupport
	case bson.TypeInt64:
		return notSupport
	case bson.TypeDecimal128:
		return notSupport
	case bson.TypeMinKey:
		return notSupport
	case bson.TypeMaxKey:
		return notSupport
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareDateTime(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		break
	case bson.TypeString:
		break
	case bson.TypeEmbeddedDocument:
		break
	case bson.TypeArray:
		break
	case bson.TypeBinary:
		break
	case bson.TypeUndefined:
		break
	case bson.TypeObjectID:
		break
	case bson.TypeBoolean:
		break
	case bson.TypeDateTime:
		break
	case bson.TypeNull:
		break
	case bson.TypeRegex:
		break
	case bson.TypeDBPointer:
		break
	case bson.TypeJavaScript:
		break
	case bson.TypeSymbol:
		break
	case bson.TypeCodeWithScope:
		break
	case bson.TypeInt32:
		break
	case bson.TypeTimestamp:
		break
	case bson.TypeInt64:
		break
	case bson.TypeDecimal128:
		break
	case bson.TypeMinKey:
		break
	case bson.TypeMaxKey:
		break
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareNull(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		break
	case bson.TypeString:
		break
	case bson.TypeEmbeddedDocument:
		break
	case bson.TypeArray:
		break
	case bson.TypeBinary:
		break
	case bson.TypeUndefined:
		break
	case bson.TypeObjectID:
		break
	case bson.TypeBoolean:
		break
	case bson.TypeDateTime:
		break
	case bson.TypeNull:
		break
	case bson.TypeRegex:
		break
	case bson.TypeDBPointer:
		break
	case bson.TypeJavaScript:
		break
	case bson.TypeSymbol:
		break
	case bson.TypeCodeWithScope:
		break
	case bson.TypeInt32:
		break
	case bson.TypeTimestamp:
		break
	case bson.TypeInt64:
		break
	case bson.TypeDecimal128:
		break
	case bson.TypeMinKey:
		break
	case bson.TypeMaxKey:
		break
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareRegex(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		break
	case bson.TypeString:
		break
	case bson.TypeEmbeddedDocument:
		break
	case bson.TypeArray:
		break
	case bson.TypeBinary:
		break
	case bson.TypeUndefined:
		break
	case bson.TypeObjectID:
		break
	case bson.TypeBoolean:
		break
	case bson.TypeDateTime:
		break
	case bson.TypeNull:
		break
	case bson.TypeRegex:
		break
	case bson.TypeDBPointer:
		break
	case bson.TypeJavaScript:
		break
	case bson.TypeSymbol:
		break
	case bson.TypeCodeWithScope:
		break
	case bson.TypeInt32:
		break
	case bson.TypeTimestamp:
		break
	case bson.TypeInt64:
		break
	case bson.TypeDecimal128:
		break
	case bson.TypeMinKey:
		break
	case bson.TypeMaxKey:
		break
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareDBPointer(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		break
	case bson.TypeString:
		break
	case bson.TypeEmbeddedDocument:
		break
	case bson.TypeArray:
		break
	case bson.TypeBinary:
		break
	case bson.TypeUndefined:
		break
	case bson.TypeObjectID:
		break
	case bson.TypeBoolean:
		break
	case bson.TypeDateTime:
		break
	case bson.TypeNull:
		break
	case bson.TypeRegex:
		break
	case bson.TypeDBPointer:
		break
	case bson.TypeJavaScript:
		break
	case bson.TypeSymbol:
		break
	case bson.TypeCodeWithScope:
		break
	case bson.TypeInt32:
		break
	case bson.TypeTimestamp:
		break
	case bson.TypeInt64:
		break
	case bson.TypeDecimal128:
		break
	case bson.TypeMinKey:
		break
	case bson.TypeMaxKey:
		break
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareJavaScript(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		break
	case bson.TypeString:
		break
	case bson.TypeEmbeddedDocument:
		break
	case bson.TypeArray:
		break
	case bson.TypeBinary:
		break
	case bson.TypeUndefined:
		break
	case bson.TypeObjectID:
		break
	case bson.TypeBoolean:
		break
	case bson.TypeDateTime:
		break
	case bson.TypeNull:
		break
	case bson.TypeRegex:
		break
	case bson.TypeDBPointer:
		break
	case bson.TypeJavaScript:
		break
	case bson.TypeSymbol:
		break
	case bson.TypeCodeWithScope:
		break
	case bson.TypeInt32:
		break
	case bson.TypeTimestamp:
		break
	case bson.TypeInt64:
		break
	case bson.TypeDecimal128:
		break
	case bson.TypeMinKey:
		break
	case bson.TypeMaxKey:
		break
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareSymbol(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		break
	case bson.TypeString:
		break
	case bson.TypeEmbeddedDocument:
		break
	case bson.TypeArray:
		break
	case bson.TypeBinary:
		break
	case bson.TypeUndefined:
		break
	case bson.TypeObjectID:
		break
	case bson.TypeBoolean:
		break
	case bson.TypeDateTime:
		break
	case bson.TypeNull:
		break
	case bson.TypeRegex:
		break
	case bson.TypeDBPointer:
		break
	case bson.TypeJavaScript:
		break
	case bson.TypeSymbol:
		break
	case bson.TypeCodeWithScope:
		break
	case bson.TypeInt32:
		break
	case bson.TypeTimestamp:
		break
	case bson.TypeInt64:
		break
	case bson.TypeDecimal128:
		break
	case bson.TypeMinKey:
		break
	case bson.TypeMaxKey:
		break
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareCodeWithScope(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		break
	case bson.TypeString:
		break
	case bson.TypeEmbeddedDocument:
		break
	case bson.TypeArray:
		break
	case bson.TypeBinary:
		break
	case bson.TypeUndefined:
		break
	case bson.TypeObjectID:
		break
	case bson.TypeBoolean:
		break
	case bson.TypeDateTime:
		break
	case bson.TypeNull:
		break
	case bson.TypeRegex:
		break
	case bson.TypeDBPointer:
		break
	case bson.TypeJavaScript:
		break
	case bson.TypeSymbol:
		break
	case bson.TypeCodeWithScope:
		break
	case bson.TypeInt32:
		break
	case bson.TypeTimestamp:
		break
	case bson.TypeInt64:
		break
	case bson.TypeDecimal128:
		break
	case bson.TypeMinKey:
		break
	case bson.TypeMaxKey:
		break
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareInt32(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		l := float64(left.Int32())
		r := right.Double()
		if l > r {
			return greater
		} else if l < r {

			return less
		} else {
			return eq
		}
	case bson.TypeString:
		break
	case bson.TypeEmbeddedDocument:
		break
	case bson.TypeArray:
		break
	case bson.TypeBinary:
		break
	case bson.TypeUndefined:
		break
	case bson.TypeObjectID:
		break
	case bson.TypeBoolean:
		break
	case bson.TypeDateTime:
		break
	case bson.TypeNull:
		break
	case bson.TypeRegex:
		break
	case bson.TypeDBPointer:
		break
	case bson.TypeJavaScript:
		break
	case bson.TypeSymbol:
		break
	case bson.TypeCodeWithScope:
		break
	case bson.TypeInt32:
		l := left.Int32()
		r := left.Int32()
		if l > r {
			return greater
		} else if l < r {

			return less
		} else {
			return eq
		}
	case bson.TypeTimestamp:
		break
	case bson.TypeInt64:
		l := int64(left.Int32())
		r := left.Int64()
		if l > r {
			return greater
		} else if l < r {

			return less
		} else {
			return eq
		}
	case bson.TypeDecimal128:
		break
	case bson.TypeMinKey:
		break
	case bson.TypeMaxKey:
		break
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareTimestamp(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		break
	case bson.TypeString:
		break
	case bson.TypeEmbeddedDocument:
		break
	case bson.TypeArray:
		break
	case bson.TypeBinary:
		break
	case bson.TypeUndefined:
		break
	case bson.TypeObjectID:
		break
	case bson.TypeBoolean:
		break
	case bson.TypeDateTime:
		break
	case bson.TypeNull:
		break
	case bson.TypeRegex:
		break
	case bson.TypeDBPointer:
		break
	case bson.TypeJavaScript:
		break
	case bson.TypeSymbol:
		break
	case bson.TypeCodeWithScope:
		break
	case bson.TypeInt32:
		break
	case bson.TypeTimestamp:
		//l := left.Timestamp()
		break
	case bson.TypeInt64:
		break
	case bson.TypeDecimal128:
		break
	case bson.TypeMinKey:
		break
	case bson.TypeMaxKey:
		break
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareInt64(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		l := float64(left.Int64())
		r := left.Double()
		if l > r {
			return greater
		} else if l < r {

			return less
		} else {
			return eq
		}
	case bson.TypeString:
		break
	case bson.TypeEmbeddedDocument:
		break
	case bson.TypeArray:
		break
	case bson.TypeBinary:
		break
	case bson.TypeUndefined:
		break
	case bson.TypeObjectID:
		break
	case bson.TypeBoolean:
		break
	case bson.TypeDateTime:
		break
	case bson.TypeNull:
		break
	case bson.TypeRegex:
		break
	case bson.TypeDBPointer:
		break
	case bson.TypeJavaScript:
		break
	case bson.TypeSymbol:
		break
	case bson.TypeCodeWithScope:
		break
	case bson.TypeInt32:
		l := left.Int64()
		r := int64(left.Int32())
		if l > r {
			return greater
		} else if l < r {

			return less
		} else {
			return eq
		}
	case bson.TypeTimestamp:
		break
	case bson.TypeInt64:
		l := left.Int64()
		r := left.Int64()
		if l > r {
			return greater
		} else if l < r {

			return less
		} else {
			return eq
		}
	case bson.TypeDecimal128:
		break
	case bson.TypeMinKey:
		break
	case bson.TypeMaxKey:
		break
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareDecimal128(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		break
	case bson.TypeString:
		break
	case bson.TypeEmbeddedDocument:
		break
	case bson.TypeArray:
		break
	case bson.TypeBinary:
		break
	case bson.TypeUndefined:
		break
	case bson.TypeObjectID:
		break
	case bson.TypeBoolean:
		break
	case bson.TypeDateTime:
		break
	case bson.TypeNull:
		break
	case bson.TypeRegex:
		break
	case bson.TypeDBPointer:
		break
	case bson.TypeJavaScript:
		break
	case bson.TypeSymbol:
		break
	case bson.TypeCodeWithScope:
		break
	case bson.TypeInt32:
		break
	case bson.TypeTimestamp:
		break
	case bson.TypeInt64:
		break
	case bson.TypeDecimal128:
		break
	case bson.TypeMinKey:
		break
	case bson.TypeMaxKey:
		break
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareMinKey(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		break
	case bson.TypeString:
		break
	case bson.TypeEmbeddedDocument:
		break
	case bson.TypeArray:
		break
	case bson.TypeBinary:
		break
	case bson.TypeUndefined:
		break
	case bson.TypeObjectID:
		break
	case bson.TypeBoolean:
		break
	case bson.TypeDateTime:
		break
	case bson.TypeNull:
		break
	case bson.TypeRegex:
		break
	case bson.TypeDBPointer:
		break
	case bson.TypeJavaScript:
		break
	case bson.TypeSymbol:
		break
	case bson.TypeCodeWithScope:
		break
	case bson.TypeInt32:
		break
	case bson.TypeTimestamp:
		break
	case bson.TypeInt64:
		break
	case bson.TypeDecimal128:
		break
	case bson.TypeMinKey:
		break
	case bson.TypeMaxKey:
		break
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
}
func compareMaxKey(left, right bson.RawValue) int {
	switch right.Type {
	case bson.TypeDouble:
		break
	case bson.TypeString:
		break
	case bson.TypeEmbeddedDocument:
		break
	case bson.TypeArray:
		break
	case bson.TypeBinary:
		break
	case bson.TypeUndefined:
		break
	case bson.TypeObjectID:
		break
	case bson.TypeBoolean:
		break
	case bson.TypeDateTime:
		break
	case bson.TypeNull:
		break
	case bson.TypeRegex:
		break
	case bson.TypeDBPointer:
		break
	case bson.TypeJavaScript:
		break
	case bson.TypeSymbol:
		break
	case bson.TypeCodeWithScope:
		break
	case bson.TypeInt32:
		break
	case bson.TypeTimestamp:
		break
	case bson.TypeInt64:
		break
	case bson.TypeDecimal128:
		break
	case bson.TypeMinKey:
		break
	case bson.TypeMaxKey:
		break
	default:
		log.Logvf(log.Always, "unknown type %v", right.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport
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

	//if checkTypeSupport(left) && checkTypeSupport(right) {
	//	return false
	//}

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

	switch left.Type {
	case bson.TypeDouble:
		return compareDouble(left, right)
	case bson.TypeString:
		return compareString(left, right)
	case bson.TypeEmbeddedDocument:
		return compareEmbeddedDocument(left, right)
	case bson.TypeArray:
		return compareArray(left, right)
	case bson.TypeBinary:
		return compareBinary(left, right)
	case bson.TypeUndefined:
		return compareUndefined(left, right)
	case bson.TypeObjectID:
		return compareObjectID(left, right)
	case bson.TypeBoolean:
		return compareBoolean(left, right)
	case bson.TypeDateTime:
		return compareDateTime(left, right)
	case bson.TypeNull:
		return compareNull(left, right)
	case bson.TypeRegex:
		return compareRegex(left, right)
	case bson.TypeDBPointer:
		return compareDBPointer(left, right)
	case bson.TypeJavaScript:
		return compareJavaScript(left, right)
	case bson.TypeSymbol:
		return compareSymbol(left, right)
	case bson.TypeCodeWithScope:
		return compareCodeWithScope(left, right)
	case bson.TypeInt32:
		return compareInt32(left, right)
	case bson.TypeTimestamp:
		return compareTimestamp(left, right)
	case bson.TypeInt64:
		return compareInt64(left, right)
	case bson.TypeDecimal128:
		return compareDecimal128(left, right)
	case bson.TypeMinKey:
		return compareMinKey(left, right)
	case bson.TypeMaxKey:
		return compareMaxKey(left, right)
	default:
		log.Logvf(log.Always, "unknown type %v", left.Type)
		os.Exit(util.ExitFailure)
	}

	return notSupport

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
