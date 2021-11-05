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

	KeywordExpr       = "$expr"
	KeywordJsonSchema = "$jsonSchema"
	KeywordMod        = "$mod"
	KeywordRegex      = "$regex"
	KeywordText       = "$text"
	KeywordWhere      = "$where"

	KeywordGeoIntersects = "$geoIntersects"
	KeywordGeoWithin     = "$geoWithin"
	KeywordNear          = "$near"
	KeywordNearSphere    = "$nearSphere"

	KeywordAll       = "$all"
	KeywordElemMatch = "$elemMatch"
	KeywordSize      = "$size"

	KeywordBitsAllClear = "$bitsAllClear"
	KeywordBitsAllSet   = "$bitsAllSet"
	KeywordBitsAnyClear = "$bitsAnyClear"
	KeywordBitsAnySet   = "$bitsAnySet"

	KeywordComment = "$comment"

	KeywordMeta  = "$meta"
	KeywordSlice = "$slice"
)

type Token struct {
}
