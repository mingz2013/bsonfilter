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
	Query     string
	IsDebug   bool
}

func (options *Options) getBSONReader() (io.ReadCloser, error) {
	if options.IFileName != "" {
		file, err := os.Open(util.ToUniversalPath(options.IFileName))
		if err != nil {
			return nil, fmt.Errorf("couldn't open BSON file: %v", err)
		}
		return file, nil
	}
	return bsondump.ReadNopCloser{Reader: os.Stdin}, nil
}

func (options *Options) getWriter() (io.WriteCloser, error) {
	if options.OFileName != "" {
		file, err := os.Create(util.ToUniversalPath(options.OFileName))
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	return bsondump.WriteNopCloser{Writer: os.Stdout}, nil
}

func (options *Options) getInterpreter() *interpreter2.Interpreter {
	log.Logv(log.Always, "getInterpreter...")
	interpreter := interpreter2.New()

	var query bson.D
	err := bson.UnmarshalExtJSON([]byte(options.Query), false, &query)
	if err != nil {
		log.Logvf(log.Always, "error parsing query as Extended JSON: %v", err)
		//os.Exit(util.ExitFailure)
		panic(err)
	}

	b, err := bson.Marshal(query)

	if err != nil {
		log.Logvf(log.Always, "error marshal bson query, %v", options.Query)
		//os.Exit(util.ExitFailure)
		panic(err)
	}

	raw := bson.Raw(b)
	log.Logvf(log.Always, "query raw: %v", raw)

	interpreter.Init(raw)
	return interpreter
}

func ParseFlag() (*Options, error) {
	isHelp := flag.Bool("h", false, "show help")
	isDebug := flag.Bool("v", false, "show details info")
	oFilePtr := flag.String("o", "", "output file")
	iFilePtr := flag.String("i", "", "input file")
	queryPtr := flag.String("q", "", "query filter, as a json string")

	flag.Parse()
	//log.Logvf(log.Always, "show help: %v", isHelp)
	//log.Logvf(log.Always, "input file: %v", *iFilePtr)
	//log.Logvf(log.Always, "output file: %v", *oFilePtr)
	//log.Logvf(log.Always, "query: %v", *queryPtr)

	if *isHelp || *oFilePtr == "" || *iFilePtr == "" || *queryPtr == "" {
		//log.Logv(log.Always, "Usage for bsonfilter:")
		flag.Usage()
		return nil, nil
	}

	return &Options{
		IFileName: *iFilePtr,
		OFileName: *oFilePtr,
		Query:     *queryPtr,
		IsDebug:   *isDebug,
	}, nil
}
