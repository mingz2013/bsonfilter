./bin/bsonfilter.darwin -i oplog.enc.bson -o tmp.bson -v -q '{"$and": [{"ns": {"$regex": "data.user_"}}, {"$or": [{"op": "u", "o2._id": {"$in": [45921229, 5001059177]}},{"op": "i", "o._id": {"$in": [45921229, 5001059177]}},{"op": "d", "o._id": {"$in": [45921229, 5001059177]}}]}]}'
