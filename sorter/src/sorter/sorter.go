package main

import (
	"fmt"
	"flag"
)

var infile *string = flag.String("i","infile","File contains value for sorting")
var outfile *string = flag.String("o","outfile","File to receive sorted values")
var algorithm *string = flag.String( "a", "qsort","Sort algorithm")

func main()  {
	flag.Parse()
	if infile != nil{
		fmt.Println("Infile=",*infile," Outfile=", *outfile, " algorithm=", *algorithm)
	}
}

