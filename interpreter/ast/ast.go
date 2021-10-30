package ast

type Node interface {
	// 节点
	Execute()
}

type Literal struct {
	// 字面量

	Value interface{}
}

func (node Literal) Execute() {

}

type Identifier struct {
	// 标识符
	Key string
}

func (node Identifier) Execute() {

}

type Expression interface {
	// 表达式

}

type UnaryExpression struct {
	// 一元运算符

	Value Node
}

func (node UnaryExpression) Execute() {

}

type BinaryOperationExpression struct {
	// 二元运算符

	Value1 Node
	Value2 Node
}

func (node BinaryOperationExpression) Execute() {

}

type ComparisonExpression struct {
	// 比较运算
	Value1 Node
	Value2 Node
}

func (node ComparisonExpression) Execute() {

}

type EqualExpression struct {
	// 等于运算符
	Value1 Node
	Value2 Node
}

func (node EqualExpression) Execute() {

}

type NotEqualExpression struct {
	// 不等于运算符
	Value1 Node
	Value2 Node
}

func (node NotEqualExpression) Execute() {

}

type LessThanExpression struct {
	// 小于
	Value1 Node
	Value2 Node
}

func (node LessThanExpression) Execute() {

}

type LessThanOrEqualExpression struct {
	// 小于等于
	Value1 Node
	Value2 Node
}

func (node LessThanOrEqualExpression) Execute() {

}

type GreaterThanExpression struct {
	// 大于
	Value1 Node
	Value2 Node
}

func (node GreaterThanExpression) Execute() {

}

type GreaterThanOrEqualExpression struct {
	// 大于等于
	Value1 Node
	Value2 Node
}

func (node GreaterThanOrEqualExpression) Execute() {

}

type InExpression struct {
	Value1 Node
	Value2 Node
}

func (node InExpression) Execute() {

}

type BooleanExpression struct {
	BinaryOperationExpression
}

func (node BooleanExpression) Execute() {

}

type OrExpression struct {
	Value1 Node
	Value2 Node
}

func (node OrExpression) Execute() {

}

type AndExpression struct {
	//BooleanExpression
	Value1 Node
	Value2 Node
}

func (node AndExpression) Execute() {

}

type NotExpression struct {
	UnaryExpression
}

func (node NotExpression) Execute() {

}

type ExistsExpression struct {
	Value1 Node
	Value2 Node
}

func (node ExistsExpression) Execute() {

}
