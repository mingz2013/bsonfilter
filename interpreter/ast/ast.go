package ast

import (
	"github.com/mingz2013/bsonfilter/interpreter/rawwrapper"
)

type Node interface {
	// 节点
	Execute(wrapper *rawwrapper.RawWrapper) bool
}

type Literal struct {
	// 字面量

	Value interface{}
}

func (node Literal) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type Identifier struct {
	// 标识符
	Key string
}

func (node Identifier) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type UnaryExpression struct {
	// 一元运算符

	Value Node
}

func (node UnaryExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type BinaryOperationExpression struct {
	// 二元运算符

	Value1 Node
	Value2 Node
}

func (node BinaryOperationExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type ComparisonExpression struct {
	// 比较运算
	Value1 Node
	Value2 Node
}

func (node ComparisonExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type EqualExpression struct {
	// 等于运算符
	Value1 Node
	Value2 Node
}

func (node EqualExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type NotEqualExpression struct {
	// 不等于运算符
	Value1 Node
	Value2 Node
}

func (node NotEqualExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type LessThanExpression struct {
	// 小于
	Value1 Node
	Value2 Node
}

func (node LessThanExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type LessThanOrEqualExpression struct {
	// 小于等于
	Value1 Node
	Value2 Node
}

func (node LessThanOrEqualExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type GreaterThanExpression struct {
	// 大于
	Value1 Node
	Value2 Node
}

func (node GreaterThanExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type GreaterThanOrEqualExpression struct {
	// 大于等于
	Value1 Node
	Value2 Node
}

func (node GreaterThanOrEqualExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type InExpression struct {
	Value1 Node
	Value2 Node
}

func (node InExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type BooleanExpression struct {
	BinaryOperationExpression
}

func (node BooleanExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type OrExpression struct {
	Value1 Node
	Value2 Node
}

func (node OrExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type AndExpression struct {
	//BooleanExpression
	Value1 Node
	Value2 Node
}

func (node AndExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type NotExpression struct {
	UnaryExpression
}

func (node NotExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}

type ExistsExpression struct {
	Value1 Node
	Value2 Node
}

func (node ExistsExpression) Execute(wrapper *rawwrapper.RawWrapper) bool {
	return false
}
