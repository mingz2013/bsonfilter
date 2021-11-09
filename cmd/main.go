package main

import (
	"github.com/mingz2013/bsonfilter/bsonfilter"
	"github.com/mongodb/mongo-tools/common/log"
	"github.com/mongodb/mongo-tools/common/util"
	"os"
	"runtime"
)

func run() {

	//defer func() {
	//	log.Logvf(log.Always, "run exist ..")
	//	os.Exit(0)
	//}()

	options, err := bsonfilter.ParseFlag()
	if err != nil {
		log.Logv(log.Always, err.Error())
		os.Exit(util.ExitFailure)

	}

	if options == nil {
		log.Logvf(log.Info, "options is none, %v", options)
		os.Exit(0)
	}

	filter, err := bsonfilter.New(options)

	if err != nil {
		log.Logv(log.Always, err.Error())
		panic(err)
	}
	defer func() {
		err := filter.Close()
		if err != nil {
			log.Logvf(log.Always, "error cleaning up: %v", err)
			panic(err)
		}
	}()

	numAll, numFound := filter.Run()
	log.Logvf(log.Info, "numAll: %v, numFound: %v", numAll, numFound)

}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	run()

}
