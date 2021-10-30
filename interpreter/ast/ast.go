package ast

import "github.com/mingz2013/bsonfilter/interpreter"

type Node interface {
	// 节点
	Execute(intepreterInterface interpreter.IntepreterInterface) bool
}

type Literal struct {
	// 字面量

	Value interface{}
}

func (node Literal) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type Identifier struct {
	// 标识符
	Key string
}

func (node Identifier) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type UnaryExpression struct {
	// 一元运算符

	Value Node
}

func (node UnaryExpression) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type BinaryOperationExpression struct {
	// 二元运算符

	Value1 Node
	Value2 Node
}

func (node BinaryOperationExpression) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type ComparisonExpression struct {
	// 比较运算
	Value1 Node
	Value2 Node
}

func (node ComparisonExpression) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type EqualExpression struct {
	// 等于运算符
	Value1 Node
	Value2 Node
}

func (node EqualExpression) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type NotEqualExpression struct {
	// 不等于运算符
	Value1 Node
	Value2 Node
}

func (node NotEqualExpression) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type LessThanExpression struct {
	// 小于
	Value1 Node
	Value2 Node
}

func (node LessThanExpression) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type LessThanOrEqualExpression struct {
	// 小于等于
	Value1 Node
	Value2 Node
}

func (node LessThanOrEqualExpression) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type GreaterThanExpression struct {
	// 大于
	Value1 Node
	Value2 Node
}

func (node GreaterThanExpression) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type GreaterThanOrEqualExpression struct {
	// 大于等于
	Value1 Node
	Value2 Node
}

func (node GreaterThanOrEqualExpression) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type InExpression struct {
	Value1 Node
	Value2 Node
}

func (node InExpression) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type BooleanExpression struct {
	BinaryOperationExpression
}

func (node BooleanExpression) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type OrExpression struct {
	Value1 Node
	Value2 Node
}

func (node OrExpression) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type AndExpression struct {
	//BooleanExpression
	Value1 Node
	Value2 Node
}

func (node AndExpression) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type NotExpression struct {
	UnaryExpression
}

func (node NotExpression) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}

type ExistsExpression struct {
	Value1 Node
	Value2 Node
}

func (node ExistsExpression) Execute(intepreterInterface interpreter.IntepreterInterface) bool {
	return false
}
