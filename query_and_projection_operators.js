// https://docs.mongodb.com/manual/reference/operator/query/

// init
db.data.insert({"_id": 1})
db.data.insert({"_id": 2})
db.data.insert({"_id": 3})

// 查询选择器, 只用在值的处理上
// $eq
db.data.find({"_id": 1})
db.data.find({"_id": {"$eq": 1}})
// $gt
db.data.find({"_id": {"$gt": 1}})
// $gte
db.data.find({"_id": {"$gte": 1}})
// $in
db.data.find({"_id": {"$in": [1,2,3]}})
// $lt
db.data.find({"_id": {"$lt": 1}})
// $lte
db.data.find({"_id": {"$lte": 1}})
// $ne
db.data.find({"_id": {"$ne": 1}})
// $nin
db.data.find({"_id": {"$nin": [1,2]}})

// 逻辑运算符，不能用到字段的后面，只能用到最外层。
// $and
db.data.find({"_id": { "$gt":1, "lt":2 }})
db.data.find({"$and": [{"_id": {"$gt": 1}}, {"_id": {"$lt": 3}}]})
// $not, 用到值后面
db.data.find({"_id": {"$not": {"$gt": 2}}})

// $nor [], 所有都为false
db.data.find({"$nor": [{"_id": 1}, {"_id": 2}]})
// $or
db.data.find({"$or": [{"_id": 1}, {"_id": 2}]})


// 元素判断，只能用到值上
// $exists
// $type  类型判断


// TODO
// 评估，
// 地理空间
// 数组
// 按位
// 投影运算符
// 杂项




db.data.find({
    "_id": {
            "$gt": 1,
            "$lt": 3
    }
})

db.data.find({
    "_id": {
        "$in": [1,2]
    }
})


