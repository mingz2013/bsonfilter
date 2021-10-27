package bsonfilter

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/mongodb/mongo-tools/bsondump"
	"github.com/mongodb/mongo-tools/common/log"
	"github.com/mongodb/mongo-tools/common/util"
	"io"
	"os"
)

type Options struct {
	IFileName string
	OFileName string
	UserIdMap map[int64]bool
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

func (options *Options) CheckUserId(userId int64) bool {
	if _, ok := options.UserIdMap[userId]; ok {
		return true
	}
	return false
}

func ParseFlag() (*Options, error) {
	oFilePtr := flag.String("o", "", "output file")
	iFilePtr := flag.String("i", "", "input file")
	userIdListStrPtr := flag.String("userids", "", "UserId List, split by ','")

	flag.Parse()
	log.Logvf(log.Always, "input file: %s", *iFilePtr)
	log.Logvf(log.Always, "output file: %s", *oFilePtr)
	log.Logvf(log.Always, "user id list: %v", *userIdListStrPtr)
	jsonStr := "[" + *userIdListStrPtr + "]"
	log.Logvf(log.Always, "jsonStr: %v", jsonStr)
	var userIdList []int64
	err := json.Unmarshal([]byte(jsonStr), &userIdList)
	if err != nil {
		return nil, err
	}
	log.Logvf(log.Always, "userIdList: %v", userIdList)
	userIdMap := make(map[int64]bool)

	for _, userId := range userIdList {
		userIdMap[userId] = true
	}

	log.Logvf(log.Always, "userIdMap: %v", userIdMap)

	return &Options{
		IFileName: *iFilePtr,
		OFileName: *oFilePtr,
		UserIdMap: userIdMap,
	}, nil
}
