package ast

type Node interface {
	// 节点
	Execute()
}

type EndNode struct {
	// 叶子结点
	Node
}

type StringLiteral struct {
	// 字符串字面量
	EndNode
}

type DigitLiteral struct {
	// 数字字面值
	EndNode
}

type Integer struct {
	// 整数
	EndNode
}

type FloatNumber struct {
	// 浮点数
	EndNode
}

type Identifier struct {
	// 标识符
	EndNode
}

type Expression interface {
	// 表达式
	Node
}

type UnaryExpression struct {
	// 一元运算符
	Expression
	Value *Node
}

type BinaryOperationExpression struct {
	// 二元运算符
	Expression
	Value1 *Node
	Value2 *Node
}

type ComparisonExpression struct {
	// 比较运算
	BinaryOperationExpression
}

type EqualExpression struct {
	// 等于运算符
	ComparisonExpression
}

type NotEqualExpression struct {
	// 不等于运算符
	ComparisonExpression
}

type LessThanExpression struct {
	// 小于
	ComparisonExpression
}

type LessThanOrEqualExpression struct {
	// 小于等于
	ComparisonExpression
}

type GreaterThanExpression struct {
	// 大于
	ComparisonExpression
}

type GreaterThanOrEqualExpression struct {
	// 大于等于
	ComparisonExpression
}

type InExpression struct {
	ComparisonExpression
}

type BooleanExpression struct {
	BinaryOperationExpression
}

type OrExpression struct {
	BooleanExpression
}

type AndExpression struct {
	BooleanExpression
}

type NotExpression struct {
	UnaryExpression
}
