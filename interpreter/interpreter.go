package interpreter

import (
	"github.com/mingz2013/bsonfilter/interpreter/ast"
	"github.com/mingz2013/bsonfilter/interpreter/parser"
	"github.com/mingz2013/bsonfilter/interpreter/rawwrapper"
	"github.com/mongodb/mongo-tools/common/log"
	"go.mongodb.org/mongo-driver/bson"
)

type Interpreter struct {
	Ast        ast.Expression
	Parser     *parser.Parser
	RawWrapper *rawwrapper.RawWrapper
}

func New() *Interpreter {
	ins := &Interpreter{}
	ins.Parser = &parser.Parser{}
	return ins
}

func (interpreter *Interpreter) Init(query bson.Raw) *Interpreter {
	interpreter.Ast = interpreter.Parser.Parse(query)
	log.Logvf(log.DebugHigh, "Init..., %v", interpreter.Ast)
	return interpreter
}

func (interpreter *Interpreter) Check(raw *bson.Raw) bool {
	interpreter.RawWrapper = &rawwrapper.RawWrapper{Raw: raw}
	return interpreter.Ast.Execute(interpreter.RawWrapper)
}
