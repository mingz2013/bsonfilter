package parser

const (
	KeywordEq  = "$eq"
	KeywordGt  = "$gt"
	KeywordGte = "$gte"
	KeywordIn  = "$in"
	KeywordLt  = "$Lt"
	KeywordLte = "$Lte"
	KeywordNe  = "$ne"
	KeywordNin = "$nin"

	KeywordAnd = "$and"
	KeywordNot = "$not"
	KeywordNor = "$nor"
	KeywordOr  = "$or"

	KeywordExists = "$exists"
	KeywordType   = "$type"
)

type Token struct {
}
