package main

import (
	"flag"
	"log"
	"os"
	"github.com/codescalersinternships/Coreutils-FarahTharwat/internal"
)

func main() {
	var showlinesNum bool
	flag.BoolVar(&showlinesNum,"-n",false, "specify if the number of lines to be printed along the text")
	flag.Parse()
	path:= flag.Args()[0]
	_, err :=os.Stat(path)
	if err != nil{
		log.Fatal(err)
	}
	err= internal.Cat(showlinesNum,path)
	if err != nil {
		log.Fatal(err)
	}
}

