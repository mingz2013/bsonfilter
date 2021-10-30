package main

import (
	"github.com/mingz2013/bsonfilter/bsonfilter"
	interpreter2 "github.com/mingz2013/bsonfilter/interpreter"
	"github.com/mingz2013/bsonfilter/interpreter/parser"
	"github.com/mongodb/mongo-tools/common/log"
	"github.com/mongodb/mongo-tools/common/util"
	"os"
)

//var oFile = "./data.o.bson"
//var iFile = "./data.i.bson"
//var oFile = ""
//var iFile = ""

func run() {

	options, err := bsonfilter.ParseFlag()
	if err != nil {
		log.Logv(log.Always, err.Error())
		os.Exit(util.ExitFailure)
	}
	filter, err := bsonfilter.New(options)

	if err != nil {
		log.Logv(log.Always, err.Error())
		os.Exit(util.ExitFailure)
	}
	defer func() {
		err := filter.Close()
		if err != nil {
			log.Logvf(log.Always, "error cleaning up: %v", err)
			os.Exit(util.ExitFailure)
		}
	}()

	numAll, numFound, err := filter.Run()
	log.Logvf(log.Always, "numAll: %v, numFound: %v", numAll, numFound)
	if err != nil {
		log.Logv(log.Always, err.Error())
		os.Exit(util.ExitFailure)
	}
}

func run2() {
	options, err := bsonfilter.ParseFlag()
	if err != nil {
		log.Logvf(log.Always, err.Error())
		os.Exit(util.ExitFailure)
	}

	p := parser.Parser{}
	node := p.Parse(options.Query)
	log.Logvf(log.Always, "node: %v", node)
}

func run3() {
	options, err := bsonfilter.ParseFlag()
	if err != nil {
		log.Logvf(log.Always, err.Error())
		os.Exit(util.ExitFailure)
	}

	interpreter := interpreter2.New()
	interpreter.Init(options.Query)

}

func main() {

	//runtime.GOMAXPROCS(runtime.NumCPU())

	//run()
	run2()

}
