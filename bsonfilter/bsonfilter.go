package bsonfilter

import (
	"fmt"
	"github.com/mongodb/mongo-tools/common/db"
	"github.com/mongodb/mongo-tools/common/log"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type BSONFilter struct {
	options      *Options
	InputSource  *db.BSONSource
	OutputWriter io.WriteCloser
}

func (bf *BSONFilter) Init() {
	bf.InputSource.SetMaxBSONSize(16 * 1024 * 1024)
}

func (bf *BSONFilter) Close() error {
	_ = bf.InputSource.Close()
	return bf.OutputWriter.Close()
}

//
//func (bf *BSONFilter) filter(raw bson.Raw) bool {
//	//log.Logvf(log.Always, "%s", raw)
//	//fmt.Println(raw.Elements())
//	//fmt.Println(raw.Values())
//	//fmt.Println(raw.Lookup("_id").IsNumber())
//	//fmt.Println(raw.Lookup("_id").AsInt64OK())
//
//	id_ := raw.Lookup("_id")
//	if id_.IsNumber() {
//		userId, ok := id_.AsInt64OK()
//		if ok {
//			if bf.CheckUserId(userId) {
//				return true
//			}
//		}
//	}
//
//	return false
//}

//func (bf *BSONFilter) filterOplog(raw bson.Raw) bool {
//	oplog := &db.Oplog{}
//	err := bson.Unmarshal(raw, oplog)
//	if err != nil {
//		log.Logvf(log.Always, "bson unmarshal error: %v, err: %v", oplog, err)
//		return false
//	}
//	if oplog.Operation == "i" || oplog.Operation == "d" {
//		if result, ok := oplog.Object.Map()["_id"]; ok {
//			if userId, ok := result.(int64); ok {
//				if bf.CheckUserId(userId) {
//					return true
//				}
//			}
//
//		}
//	} else if oplog.Operation == "u" {
//		if result, ok := oplog.Query.Map()["_id"]; ok {
//			if userId, ok := result.(int64); ok {
//				if bf.CheckUserId(userId) {
//					return true
//				}
//			}
//
//		}
//	} else if oplog.Operation == "c" {
//		log.Logvf(log.Always, "un except operation: %v", oplog)
//	}
//
//	return false
//
//}

func (bf *BSONFilter) check(raw *bson.Raw, filter *bson.D) bool {

	return true
}

func (bf *BSONFilter) Check(raw *bson.Raw) bool {
	return bf.check(raw, &bf.options.Query)
}

func (bf *BSONFilter) Run() (numAll, numFound int, err error) {

	log.Logv(log.Always, "Run...")

	for {
		result := bson.Raw(bf.InputSource.LoadNext())
		if result == nil {
			break
		}

		var ok bool

		//if bf.IsOplog{
		//	ok = bf.filterOplog(result)
		//}else {
		//	ok = bf.filter(result)
		//}

		ok = bf.Check(&result)

		if ok {
			n, err := bf.OutputWriter.Write(result)
			if err != nil {
				log.Logvf(log.Always, "output write err: %v, n: %v", err, n)
				break
			}
			numFound++
		}

		numAll++

		//if numFound > 1{
		//	break
		//}

	}

	log.Logv(log.Always, "Run...end")

	if err := bf.InputSource.Err(); err != nil {
		return numAll, numFound, err
	}

	return numAll, numFound, nil

}

func New(options *Options) (*BSONFilter, error) {

	filter := &BSONFilter{
		options: options,
	}

	reader, err := filter.options.getBSONReader()
	if err != nil {
		return nil, fmt.Errorf("getting BSON reader failed: %v", err)
	}
	filter.InputSource = db.NewBSONSource(reader)

	writer, err := filter.options.getWriter()
	if err != nil {
		_ = filter.InputSource.Close()
		return nil, fmt.Errorf("getting Writer failed: %v", err)
	}
	filter.OutputWriter = writer

	filter.Init()

	return filter, nil
}
