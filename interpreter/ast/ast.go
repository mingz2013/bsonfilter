package ast

import (
	"github.com/mingz2013/bsonfilter/interpreter/rawwrapper"
	"go.mongodb.org/mongo-driver/bson"
)

type Node interface {
	// 节点
	//Execute(wrapper *rawwrapper.RawWrapper) bool
}

//
//type EndNode interface {
//	IsNumber() bool
//	IsString() bool
//	IsList() bool
//	GetNumber()
//	GetString() (string, error)
//	GetList() (types.Slice, error)
//}

type Literal struct {
	// 字面量

	Value bson.RawValue
}

func (literal *Literal) GetRawValue() bson.RawValue {
	return literal.Value
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

type EqualExpression struct {
	// 等于运算符
	Value1 Identifier
	Value2 Literal
}

func (node EqualExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	//return node.Value1.Execute(wrapper) == node.Value2.Execute(wrapper)
	//return node.Value1.GetRawValue(wrapper).Equal(node.Value2.GetRawValue())
	//return false
	return rawwrapper.IsEqual(node.Value1.GetRawValue(wrapper), node.Value2.GetRawValue())

}

type LessThanExpression struct {
	// 小于
	Value1 Identifier
	Value2 Literal
}

func (node LessThanExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	//return node.Value1.Execute(wrapper) < node.Value2.Execute(wrapper)
	return rawwrapper.IsLess(node.Value1.GetRawValue(wrapper), node.Value2.GetRawValue())

}

type LessThanOrEqualExpression struct {
	// 小于等于
	Value1 Identifier
	Value2 Literal
}

func (node LessThanOrEqualExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	//return node.Value1.Execute(wrapper) <= node.Value2.Execute(wrapper)
	return rawwrapper.IsLessOrEq(node.Value1.GetRawValue(wrapper), node.Value2.GetRawValue())

}

type GreaterThanExpression struct {
	// 大于
	Value1 Identifier
	Value2 Literal
}

func (node GreaterThanExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	//return node.Value1.Execute(wrapper) > node.Value2.Execute(wrapper)
	return rawwrapper.IsGreater(node.Value1.GetRawValue(wrapper), node.Value2.GetRawValue())

}

type GreaterThanOrEqualExpression struct {
	// 大于等于
	Value1 Identifier
	Value2 Literal
}

func (node GreaterThanOrEqualExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	//return node.Value1.Execute(wrapper) >= node.Value2.Execute(wrapper)
	return rawwrapper.IsGreaterOrEq(node.Value1.GetRawValue(wrapper), node.Value2.GetRawValue())

}

type InExpression struct {
	Value1 Identifier
	Value2 Literal
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
	//BooleanExpression
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
	Value1 Identifier
	Value2 Literal
}

func (node ExistsExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	if node.Value2.GetRawValue().Boolean() {
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
