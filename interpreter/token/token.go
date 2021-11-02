package token

const (
	KeywordEq  = "$eq"
	KeywordGt  = "$gt"
	KeywordGte = "$gte"
	KeywordIn  = "$in"
	KeywordLt  = "$lt"
	KeywordLte = "$lte"
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
