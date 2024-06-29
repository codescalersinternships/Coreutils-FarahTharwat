package main

import (
	"flag"
	"log"
	"github.com/codescalersinternships/Coreutils-FarahTharwat/internal"
)
func main() {
	var numOfLines int
	flag.IntVar(&numOfLines,"-n",10,"specify number of lines to be printed")
	flag.Parse()
    path:= flag.Args()[0]
	err := internal.Tail(numOfLines,path)
	if err != nil {
		log.Fatal(err)
	}
}
