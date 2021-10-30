package bsonfilter

import (
	"flag"
	"fmt"
	interpreter2 "github.com/mingz2013/bsonfilter/interpreter"
	"github.com/mongodb/mongo-tools/bsondump"
	"github.com/mongodb/mongo-tools/common/log"
	"github.com/mongodb/mongo-tools/common/util"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"os"
)

type Options struct {
	IFileName string
	OFileName string
	//UserIdMap map[int64]bool
	//IsOplog bool
	Query bson.D
}

func (options *Options) getBSONReader() (io.ReadCloser, error) {
	if options.IFileName != "" {
		file, err := os.Open(util.ToUniversalPath(options.IFileName))
		if err != nil {
			return nil, fmt.Errorf("couldn't open BSON file: %v", err)
		}
		return file, nil
	}
	return bsondump.ReadNopCloser{os.Stdin}, nil
}

func (options *Options) getWriter() (io.WriteCloser, error) {
	if options.OFileName != "" {
		file, err := os.Create(util.ToUniversalPath(options.OFileName))
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	return bsondump.WriteNopCloser{os.Stdout}, nil
}

//func (options *Options) CheckUserId(userId int64) bool {
//	if _, ok := options.UserIdMap[userId]; ok {
//		return true
//	}
//	return false
//}

func (options *Options) getInterpreter() *interpreter2.Interpreter {
	interpreter := interpreter2.New()
	interpreter.Init(options.Query)
	return interpreter
}

func ParseFlag() (*Options, error) {
	oFilePtr := flag.String("o", "", "output file")
	iFilePtr := flag.String("i", "", "input file")
	//userIdListStrPtr := flag.String("userids", "", "UserId List, split by ','")
	//isOplog := flag.Bool("isOplog", false, "is oplog")
	queryPtr := flag.String("q", "", "query filter, as a json string")
	flag.Parse()
	log.Logvf(log.Always, "input file: %v", *iFilePtr)
	log.Logvf(log.Always, "output file: %v", *oFilePtr)
	//log.Logvf(log.Always, "user id list: %v", *userIdListStrPtr)
	//log.Logvf(log.Always, "is oplog: %v", *isOplog)
	log.Logvf(log.Always, "query: %v", *queryPtr)

	//jsonStr := "[" + *userIdListStrPtr + "]"
	//log.Logvf(log.Always, "jsonStr: %v", jsonStr)
	//var userIdList []int64
	//err := json.Unmarshal([]byte(jsonStr), &userIdList)
	//if err != nil {
	//	return nil, err
	//}
	//log.Logvf(log.Always, "userIdList: %v", userIdList)
	//userIdMap := make(map[int64]bool)
	//
	//for _, userId := range userIdList {
	//	userIdMap[userId] = true
	//}
	//
	//log.Logvf(log.Always, "userIdMap: %v", userIdMap)

	var query bson.D
	err := bson.UnmarshalExtJSON([]byte(*queryPtr), false, &query)
	if err != nil {
		return nil, fmt.Errorf("error parsing query as Extended JSON: %v", err)
	}

	return &Options{
		IFileName: *iFilePtr,
		OFileName: *oFilePtr,
		//UserIdMap: userIdMap,
		//IsOplog: *isOplog,
		Query: query,
	}, nil
}
