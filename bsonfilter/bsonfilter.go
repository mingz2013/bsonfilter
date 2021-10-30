package bsonfilter

import (
	"fmt"
	"github.com/mingz2013/bsonfilter/interpreter"
	"github.com/mongodb/mongo-tools/common/db"
	"github.com/mongodb/mongo-tools/common/log"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type BSONFilter struct {
	options      *Options
	InputSource  *db.BSONSource
	OutputWriter io.WriteCloser
	interpreter  *interpreter.Interpreter
}

func (bf *BSONFilter) Init() {
	bf.InputSource.SetMaxBSONSize(16 * 1024 * 1024)
}

func (bf *BSONFilter) Close() error {
	_ = bf.InputSource.Close()
	return bf.OutputWriter.Close()
}

func (bf *BSONFilter) Check(raw *bson.Raw) bool {
	//return bf.check(raw, &bf.options.Query)
	return bf.interpreter.Check(raw)
}

func (bf *BSONFilter) Run() (numAll, numFound int, err error) {

	log.Logv(log.Always, "Run...")

	for {
		result := bson.Raw(bf.InputSource.LoadNext())
		if result == nil {
			break
		}

		var ok bool

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
		break

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

	filter.interpreter = filter.options.getInterpreter()

	return filter, nil
}
