package rawwrapper

import (
	"bytes"
	"github.com/mongodb/mongo-tools/common/log"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	eq         = 0
	greater    = 1
	less       = -1
	notSupport = -2
)

func compareDouble(left, right bson.RawValue) int {

	switch right.Type {
	case bson.TypeDouble:
		l := left.Double()
		r := right.Double()
		if l > r {
			return greater
		} else if l < r {
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		//os.Exit(util.ExitFailure)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		right.ObjectID()
		return bytes.Compare([]byte(left.ObjectID().Hex()), []byte(right.ObjectID().Hex()))
		//return notSupport
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		l := left.DateTime()
		r := right.DateTime()
		if l > r {
			return greater
		} else if l < r {
			return less
		} else {
			return eq
		}
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		lt, li := left.Timestamp()
		rt, ri := right.Timestamp()
		if lt > rt {
			return greater
		} else if lt < rt {
			return less
		} else {
			if li > ri {
				return greater
			} else if li < ri {
				return less
			} else {
				return eq
			}
		}

	case bson.TypeInt64:
		break
	case bson.TypeDecimal128:
		break
	case bson.TypeMinKey:
		break
	case bson.TypeMaxKey:
		break
	default:
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
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
		log.Logvf(log.Info, "unknown type %v", right.Type)
		panic(right.Type)
	}

	return notSupport
}

func Compare(left, right bson.RawValue) int {

	log.Logvf(log.DebugHigh, "Compare << : left %v, right %v", left, right)

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
		log.Logvf(log.Info, "unknown type %v, left: %v", left.Type, left)
		//panic(left.Type)
		return notSupport
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
	raw := right.Array()
	elems, err := raw.Values()
	if err != nil {
		log.Logvf(log.Info, "error %v", err)
		panic(err)
	}
	for _, elem := range elems {
		if Compare(left, elem) == eq {
			return true
		}
	}
	return false

}
