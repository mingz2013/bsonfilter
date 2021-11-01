package rawwrapper

import (
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type RawWrapper struct {
	Raw *bson.Raw

	// cache
	keyMap map[string]bson.RawValue
}

func (wrapper *RawWrapper) GetRawValue(key string) bson.RawValue {

	if wrapper.keyMap == nil {
		wrapper.keyMap = map[string]bson.RawValue{}
	}

	if rawValue, ok := wrapper.keyMap[key]; ok {
		return rawValue
	}

	keyList := strings.Split(key, ".")

	rawValue := wrapper.Raw.Lookup(keyList...)

	wrapper.keyMap[key] = rawValue
	return rawValue

}
