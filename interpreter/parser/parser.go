package parser

import (
	"github.com/mingz2013/bsonfilter/interpreter/ast"
	"github.com/mingz2013/bsonfilter/interpreter/token"
	"github.com/mongodb/mongo-tools/common/log"
	"go.mongodb.org/mongo-driver/bson"
)

type Parser struct {
	//src *map[string] interface{}// query
}

func (parser *Parser) Parse(raw bson.Raw) (node ast.Expression) {
	return parser.parseRaw(raw)
}

func (parser *Parser) parseRaw(raw bson.Raw) (node ast.Expression) {
	log.Logvf(log.Always, "parseRaw...%v", raw)
	rawElems, err := raw.Elements()
	if err != nil {
		log.Logvf(log.Always, "err: %v", err)
		//os.Exit(util.ExitFailure)
		panic(err)
	}
	return parser.parseAndByRawElements(rawElems)
}

func (parser *Parser) parseAndByRawElements(rawElems []bson.RawElement) ast.Expression {

	length := len(rawElems)
	if length == 0 {
		log.Logvf(log.Always, "error rawElems: %v", rawElems)
		return nil
	} else if length == 1 {
		return parser.parseExpression(rawElems[0])
	} else {
		node := ast.AndExpression{}
		node.Value1 = parser.parseExpression(rawElems[0])
		node.Value2 = parser.parseAndByRawElements(rawElems[1:])
		return node
	}

}

func (parser *Parser) rawValueToRawValues(rawVaue bson.RawValue) []bson.RawValue {
	raw, ok := rawVaue.ArrayOK()
	if !ok {
		log.Logvf(log.Always, "error parse")
		panic(rawVaue)
	}
	rawValues, err := raw.Values()
	if err != nil {
		log.Logvf(log.Always, "error parse")
		panic(err)
	}
	return rawValues
}

func (parser *Parser) parseExpression(e bson.RawElement) ast.Expression {
	log.Logvf(log.Always, "parseExpression << e: %v, e.Key: %v, e.Value: %v", e, e.Key(), e.Value())
	switch e.Key() {
	case token.KeywordAnd:
		rawValues := parser.rawValueToRawValues(e.Value())
		return parser.parseAnd(rawValues)
	case token.KeywordOr:
		rawValues := parser.rawValueToRawValues(e.Value())
		return parser.parseOr(rawValues)
	case token.KeywordNor:
		rawValues := parser.rawValueToRawValues(e.Value())
		return parser.parseNor(rawValues)
	case token.KeywordNot:
		return parser.parseNot(e.Value().Document())

	default:
		// value key
		return parser.parseKey(e.Key(), e.Value())
	}
}

func (parser *Parser) parseNot(raw bson.Raw) ast.Expression {
	node := ast.NotExpression{}
	node.Value = parser.parseRaw(raw)
	return node
}

func (parser *Parser) parseAnd(rawValues []bson.RawValue) ast.Expression {
	log.Logvf(log.Always, "parseAnd << %v", rawValues)

	length := len(rawValues)
	if length == 0 {
		log.Logvf(log.Always, "parseAnd length 0, %v", rawValues)
		return nil
	} else if length == 1 {
		return parser.parseRaw(rawValues[0].Document())
	} else {
		node := ast.AndExpression{}
		node.Value1 = parser.parseRaw(rawValues[0].Document())
		node.Value2 = parser.parseAnd(rawValues[1:])
		return node
	}

}

func (parser *Parser) parseOr(rawValues []bson.RawValue) ast.Expression {
	log.Logvf(log.Always, "parseOr << %v", rawValues)

	length := len(rawValues)
	if length == 0 {
		log.Logvf(log.Always, "parseOr length 0, %v", rawValues)
		return nil
	} else if length == 1 {
		return parser.parseRaw(rawValues[0].Document())
	} else {
		node := ast.OrExpression{}
		node.Value1 = parser.parseRaw(rawValues[0].Document())
		node.Value2 = parser.parseOr(rawValues[1:])
		return node
	}
}

func (parser *Parser) parseNor(rawValues []bson.RawValue) ast.Expression {
	log.Logvf(log.Always, "parseNor << %v", rawValues)

	length := len(rawValues)
	if length == 0 {
		log.Logvf(log.Always, "parseNor length 0, %v", rawValues)
		return nil
	} else if length == 1 {
		node2 := ast.NotExpression{}
		node2.Value = parser.parseRaw(rawValues[0].Document())
		return node2
	} else {
		node := ast.OrExpression{}
		node2 := ast.NotExpression{}
		node2.Value = parser.parseRaw(rawValues[0].Document())
		node.Value1 = node2
		node.Value2 = parser.parseNor(rawValues[1:])
		return node
	}
}

func (parser *Parser) parseKey(key string, value bson.RawValue) ast.Expression {
	log.Logvf(log.Always, "parseKey << %v, %v", key, value)
	keyNode := ast.Identifier{}
	keyNode.Key = key

	raw, ok := value.DocumentOK()

	log.Logvf(log.Always, "parseKey, raw : %v, ok: %v", raw, ok)

	if !ok {
		node := ast.EqualExpression{}
		node.Value1 = keyNode
		node.Value2 = ast.Literal{Value: value}
		log.Logvf(log.Always, "parseKey node >> %v", node)
		return node
	}

	elems, err := raw.Elements()
	if err != nil {
		log.Logvf(log.Always, "err: %v", err)
		//os.Exit(util.ExitFailure)
		panic(err)
	}

	return parser.parseKeyAnd(keyNode, elems)

}

func (parser *Parser) parseKeyAnd(keyNode ast.Identifier, elems []bson.RawElement) ast.Expression {
	length := len(elems)
	if length == 0 {
		log.Logvf(log.Always, "parseKeyAnd length = 0, %v", elems)
		return nil
	} else if length == 1 {
		return parser.parseKeyExpression(keyNode, elems[0])
	} else {
		node := ast.AndExpression{}
		node.Value1 = parser.parseKeyExpression(keyNode, elems[0])
		node.Value2 = parser.parseKeyAnd(keyNode, elems[1:])
		return node
	}
}

func (parser *Parser) parseKeyExpression(keyNode ast.Identifier, e bson.RawElement) ast.Expression {
	log.Logvf(log.Always, "parseKeyExpression << keyNode: %v, e: %v", keyNode, e)
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
		log.Logvf(log.Always, "KeywordExists, e %v, e.Key, %v, e.Value, %s", e, e.Key(), e.Value())
		node := ast.ExistsExpression{}
		node.Value1 = keyNode
		node.Value2 = ast.Literal{Value: e.Value()}
		return node
	default:
		log.Logvf(log.Always, "error: %v, %v", keyNode, e)
		//os.Exit(util.ExitFailure)
		panic(e)

	}
	return nil
}
