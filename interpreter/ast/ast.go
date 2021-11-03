package ast

import (
	"github.com/mingz2013/bsonfilter/interpreter/rawwrapper"
	"go.mongodb.org/mongo-driver/bson"
)

type Node interface {
	// 节点
	//Execute(wrapper *rawwrapper.RawWrapper) bool
}

type Literal struct {
	// 字面量

	Value bson.RawValue
}

func (literal *Literal) GetRawValue() bson.RawValue {
	return literal.Value
}

func (literal *Literal) GetBooleanOk() (value bool, ok bool) {
	switch literal.Value.Type {
	case bson.TypeBoolean:
		return literal.Value.Boolean(), true
	case bson.TypeInt32:
		i := literal.Value.Int32()
		if i == 0 {
			return false, true
		} else {
			return true, true
		}
	default:
		return false, false
	}
}

//func (node Literal) Execute(wrapper *rawwrapper.RawWrapper) bool {
//	return false
//}

type Identifier struct {
	// 标识符
	Key string
}

func (identifier *Identifier) GetRawValue(wrapper *rawwrapper.RawWrapper) bson.RawValue {
	return wrapper.GetRawValue(identifier.Key)
}

//func (node Identifier) Execute(wrapper *rawwrapper.RawWrapper) bool {
//	return false
//}

type Expression interface {
	Execute(wrapper *rawwrapper.RawWrapper) bool
}

type BinaryOperationExpression struct {
	// 二元运算符

	Value1 Expression
	Value2 Expression
}

type KeyExpression struct {
	Value1 Identifier
	Value2 Literal
}

type EqualExpression struct {
	// 等于运算符
	KeyExpression
}

func (node EqualExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {

	return rawwrapper.IsEqual(node.Value1.GetRawValue(wrapper), node.Value2.GetRawValue())

}

type LessThanExpression struct {
	// 小于
	KeyExpression
}

func (node LessThanExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {

	return rawwrapper.IsLess(node.Value1.GetRawValue(wrapper), node.Value2.GetRawValue())

}

type LessThanOrEqualExpression struct {
	// 小于等于
	KeyExpression
}

func (node LessThanOrEqualExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {

	return rawwrapper.IsLessOrEq(node.Value1.GetRawValue(wrapper), node.Value2.GetRawValue())

}

type GreaterThanExpression struct {
	// 大于
	KeyExpression
}

func (node GreaterThanExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {

	return rawwrapper.IsGreater(node.Value1.GetRawValue(wrapper), node.Value2.GetRawValue())

}

type GreaterThanOrEqualExpression struct {
	// 大于等于
	KeyExpression
}

func (node GreaterThanOrEqualExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return rawwrapper.IsGreaterOrEq(node.Value1.GetRawValue(wrapper), node.Value2.GetRawValue())

}

type InExpression struct {
	KeyExpression
}

func (node InExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	result := rawwrapper.IsIn(node.Value1.GetRawValue(wrapper), node.Value2.GetRawValue())
	return result
}

type OrExpression struct {
	BinaryOperationExpression
}

func (node OrExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return node.Value1.Execute(wrapper) || node.Value2.Execute(wrapper)
}

type AndExpression struct {
	BinaryOperationExpression
}

func (node AndExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return node.Value1.Execute(wrapper) && node.Value2.Execute(wrapper)
}

type NotExpression struct {
	Value Expression
}

func (node NotExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return !node.Value.Execute(wrapper)
}

type ExistsExpression struct {
	KeyExpression
}

func (node ExistsExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	isExists, ok := node.Value2.GetBooleanOk()
	if !ok {
		return false
	}

	if isExists {
		if node.Value1.GetRawValue(wrapper).Value != nil {
			return true
		}
	} else {
		if node.Value1.GetRawValue(wrapper).Value == nil {
			return true
		}
	}
	return false
}
