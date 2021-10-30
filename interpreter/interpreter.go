package interpreter

import (
	"github.com/mingz2013/bsonfilter/interpreter/ast"
	"github.com/mingz2013/bsonfilter/interpreter/parser"
	"github.com/mingz2013/bsonfilter/interpreter/rawwrapper"
	"go.mongodb.org/mongo-driver/bson"
)

type Interpreter struct {
	Ast        ast.Node
	Parser     *parser.Parser
	RawWrapper *rawwrapper.RawWrapper
}

func New() *Interpreter {
	ins := &Interpreter{}
	ins.Parser = &parser.Parser{}
	return ins
}

func (interpreter *Interpreter) Init(query bson.D) *Interpreter {
	interpreter.Ast = interpreter.Parser.Parse(query)
	return interpreter
}

func (interpreter *Interpreter) Check(raw *bson.Raw) bool {
	interpreter.RawWrapper = &rawwrapper.RawWrapper{Raw: raw}
	return interpreter.Ast.Execute(interpreter.RawWrapper)
}
