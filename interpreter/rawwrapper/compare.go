package rawwrapper

import "go.mongodb.org/mongo-driver/bson"

const eq = 0
const greater = 1
const less = -1

func Compare(left, right bson.RawValue) int {
	// 0 ==
	// 1 >
	// -1 <

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
	return true
}
