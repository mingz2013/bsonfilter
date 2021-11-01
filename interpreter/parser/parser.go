package parser

import (
	"github.com/mingz2013/bsonfilter/interpreter/ast"
	"github.com/mingz2013/bsonfilter/interpreter/token"
	"github.com/mongodb/mongo-tools/common/log"
	"github.com/mongodb/mongo-tools/common/util"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

type Parser struct {
	//src *map[string] interface{}// query
}

//func (parser *Parser) Parse(d bson.D) (node ast.Expression) {
//	return parser.parseD(d)
//}

func (parser *Parser) Parse(raw bson.Raw) (node ast.Expression) {
	return parser.ParseRaw(raw)
}

func (parser *Parser) ParseRaw(raw bson.Raw) (node ast.Expression) {
	rawElems, err := raw.Elements()
	if err != nil {
		log.Logvf(log.Always, "err: %v", err)
		os.Exit(util.ExitFailure)
	}

	length := len(rawElems)
	if length == 0 {
		return nil
	} else if length == 1 {
		return parser.parseExpression(rawElems[0])
	} else {
		return parser.parseAnd(raw)
	}

}

//func (parser *Parser) parseD(d bson.D) (node ast.Expression) {
//
//	//if len == 1{
//	//	return parser.parseAnd(src)
//	//}
//
//	length := len(d)
//	if length == 0 {
//		return nil
//	} else if length == 1 {
//		return parser.parseExpression(d[0])
//	} else {
//		// > 1
//
//		return parser.parseAnd(d)
//	}
//
//}

func (parser *Parser) parseExpression(e bson.RawElement) ast.Expression {
	switch e.Key() {
	case token.KeywordAnd:
		return parser.parseAnd(e.Value().Array())
	case token.KeywordOr:
		return parser.parseOr(e.Value().Array())
	case token.KeywordNor:
		return parser.parseNor(e.Value().Array())
	default:
		// value key
		return parser.parseKey(e.Key(), e.Value())
	}
}

//func (parser *Parser) parseExpression(e bson.E) ast.Expression {
//
//	switch e.Key {
//	case token.KeywordAnd:
//		return parser.parseAnd(e.Value.([]bson.D)...)
//	case token.KeywordOr:
//		return parser.parseOr(e.Value.([]bson.D)...)
//	case token.KeywordNor:
//		return parser.parseNor(e.Value.([]bson.D)...)
//	default:
//		// value key
//		return parser.parseKey(e.Key, e.Value)
//	}
//}
func (parser *Parser) parseAnd(raw bson.Raw) ast.Expression {

	var begin ast.AndExpression
	var tmp ast.AndExpression
	rawElems, err := raw.Elements()
	if err != nil {
		log.Logvf(log.Always, "err: %v", err)
		os.Exit(util.ExitFailure)
	}

	for _, rawElem := range rawElems {

		node := ast.AndExpression{}
		node.Value1 = parser.parseExpression(rawElem)

		if &begin == nil {
			begin = node
		}
		if &tmp != nil {
			tmp.Value2 = node
		}

		tmp = node

	}
	return &begin
}

//func (parser *Parser) parseAnd(args ...bson.D) ast.Expression {
//
//	var begin ast.AndExpression
//	var tmp ast.AndExpression
//	for _, d := range args {
//
//		node := ast.AndExpression{}
//		node.Value1 = parser.parseD(d)
//
//		if &begin == nil {
//			begin = node
//		}
//		if &tmp != nil {
//			tmp.Value2 = node
//		}
//
//		tmp = node
//
//	}
//	return &begin
//}

func (parser *Parser) parseOr(raw bson.Raw) ast.Expression {

	var begin ast.OrExpression
	var tmp ast.OrExpression
	rawElems, err := raw.Elements()
	if err != nil {
		log.Logvf(log.Always, "err: %v", err)
		os.Exit(util.ExitFailure)
	}

	for _, d := range rawElems {

		node := ast.OrExpression{}
		node.Value1 = parser.parseExpression(d)

		if &begin == nil {
			begin = node
		}
		if &tmp != nil {
			tmp.Value2 = node
		}

		tmp = node

	}
	return begin
}

//func (parser *Parser) parseOr(args ...bson.D) ast.Expression {
//
//	var begin ast.OrExpression
//	var tmp ast.OrExpression
//	for _, d := range args {
//
//		node := ast.OrExpression{}
//		node.Value1 = parser.parseD(d)
//
//		if &begin == nil {
//			begin = node
//		}
//		if &tmp != nil {
//			tmp.Value2 = node
//		}
//
//		tmp = node
//
//	}
//	return begin
//}

func (parser *Parser) parseNor(raw bson.Raw) ast.Expression {
	var begin ast.OrExpression
	var tmp ast.OrExpression
	rawElems, err := raw.Elements()
	if err != nil {
		log.Logvf(log.Always, "err: %v", err)
		os.Exit(util.ExitFailure)
	}
	for _, d := range rawElems {

		node := ast.OrExpression{}
		not := ast.NotExpression{}
		node.Value1 = not
		not.Value = parser.parseExpression(d)

		if &begin == nil {
			begin = node
		}
		if &tmp != nil {
			tmp.Value2 = node
		}

		tmp = node

	}
	return begin
}

//func (parser *Parser) parseNor(args ...bson.D) ast.Expression {
//	var begin ast.OrExpression
//	var tmp ast.OrExpression
//	for _, d := range args {
//
//		node := ast.OrExpression{}
//		not := ast.NotExpression{}
//		node.Value1 = not
//		not.Value = parser.parseD(d)
//
//		if &begin == nil {
//			begin = node
//		}
//		if &tmp != nil {
//			tmp.Value2 = node
//		}
//
//		tmp = node
//
//	}
//	return begin
//}

func (parser *Parser) parseKey(key string, value bson.RawValue) ast.Expression {

	keyNode := ast.Identifier{}
	keyNode.Key = key

	//var d bson.D
	//d, ok := obj.(bson.D)

	raw, ok := value.ArrayOK()

	if !ok {
		node := ast.EqualExpression{}
		node.Value1 = keyNode
		node.Value2 = ast.Literal{Value: value}
		return node
	}
	//
	var begin ast.AndExpression
	var tmp ast.AndExpression

	elems, err := raw.Elements()
	if err != nil {
		log.Logvf(log.Always, "err: %v", err)
		os.Exit(util.ExitFailure)
	}

	for _, e := range elems {

		node := ast.AndExpression{}
		node.Value1 = parser.parseKeyExpression(keyNode, e)

		if &begin == nil {
			begin = node
		}
		if &tmp != nil {
			tmp.Value2 = node
		}

		tmp = node

	}

	return begin
}

func (parser *Parser) parseKeyExpression(keyNode ast.Identifier, e bson.RawElement) ast.Expression {
	//var node ast.Node

	switch e.Key() {
	case token.KeywordEq:
		node := ast.EqualExpression{}
		node.Value1 = keyNode
		node.Value2 = ast.Literal{Value: e.Value()}
		return node
	case token.KeywordNe:
		node := ast.NotExpression{}

		node2 := ast.EqualExpression{}
		node2.Value1 = keyNode
		node2.Value2 = ast.Literal{Value: e.Value()}
		node.Value = node2
		return node

	case token.KeywordGt:
		node := ast.GreaterThanExpression{}
		node.Value1 = keyNode
		node.Value2 = ast.Literal{Value: e.Value()}
		return node
	case token.KeywordGte:
		node := ast.GreaterThanOrEqualExpression{}
		node.Value1 = keyNode
		node.Value2 = ast.Literal{Value: e.Value()}
		return node
	case token.KeywordLt:
		node := ast.LessThanExpression{}
		node.Value1 = keyNode
		node.Value2 = ast.Literal{Value: e.Value()}
		return node
	case token.KeywordLte:
		node := ast.LessThanOrEqualExpression{}
		node.Value1 = keyNode
		node.Value2 = ast.Literal{Value: e.Value()}
		return node
	case token.KeywordIn:
		node := ast.InExpression{}
		node.Value1 = keyNode
		node.Value2 = ast.Literal{Value: e.Value()}
		return node
	case token.KeywordNin:
		nodeIn := ast.EqualExpression{}
		nodeIn.Value1 = keyNode
		nodeIn.Value2 = ast.Literal{Value: e.Value()}

		node := ast.NotExpression{}
		node.Value = nodeIn
		return node

	case token.KeywordExists:
		node := ast.ExistsExpression{}
		node.Value1 = keyNode
		node.Value2 = ast.Literal{Value: e.Value()}
		return node
	default:
		log.Logvf(log.Always, "error: %v, %v", keyNode, e)
		os.Exit(util.ExitFailure)

	}
	return nil
}
