package interpreter

import (
	"github.com/mingz2013/bsonfilter/interpreter/ast"
	"github.com/mingz2013/bsonfilter/interpreter/parser"
)

type Interpreter struct {
	Ast    *ast.Node
	Parser *parser.Parser
}
