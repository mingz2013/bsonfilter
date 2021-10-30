package rawwrapper

import "go.mongodb.org/mongo-driver/bson"

type RawWrapper struct {
	Raw *bson.Raw
}
