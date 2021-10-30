package interpreter

import (
	"github.com/mingz2013/bsonfilter/interpreter/ast"
	"github.com/mingz2013/bsonfilter/interpreter/parser"
	"go.mongodb.org/mongo-driver/bson"
)

type RawWrapper struct {
	raw *bson.Raw
}

type IntepreterInterface interface {
	GetRawWrapper() *RawWrapper
}

type Interpreter struct {
	Ast        ast.Node
	Parser     *parser.Parser
	RawWrapper *RawWrapper
}

func New() *Interpreter {
	ins := &Interpreter{}
	ins.Parser = &parser.Parser{}
	return ins
}

func (interpreter *Interpreter) GetRawWrapper() *RawWrapper {
	return interpreter.RawWrapper
}

func (interpreter *Interpreter) Init(query bson.D) *Interpreter {
	interpreter.Ast = interpreter.Parser.Parse(query)
	return interpreter
}

func (interpreter *Interpreter) Check(raw *bson.Raw) bool {
	interpreter.RawWrapper = &RawWrapper{raw: raw}
	return interpreter.Ast.Execute(interpreter)
}
