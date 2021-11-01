package rawwrapper

import (
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type RawWrapper struct {
	Raw *bson.Raw

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

//func (wrapper*RawWrapper)getNumber(key string)  {
//	wrapper.GetKey(key).IsNumber()
//}
