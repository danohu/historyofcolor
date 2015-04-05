package main

import (
	"fmt"
	"log"
	"os"
	_ "image"

	_ "image/jpeg"
)

func getwd() string{
	// find the base directory of the package
	// NB AFAICT golang has no decent way of handling packaged
	// binary data -- see e.g. the fruitless discussion at
	// https://groups.google.com/forum/#!topic/golang-nuts/rdk-HdpxQps

	var goroot string = "/home/src/golang"
	var packagename string = "github.com/danohuiginn/historyofcolor"
	return goroot + "/src/" + packagename + "/"
}

func decodeimage(fn string){
	var path string = getwd()
	fn = path + fn
	fmt.Printf(fn);
	reader, err := os.Open(fn)
	if err != nil{
		log.Fatal(err)
	}
	_ = reader
}

func main(){
	decodeimage("testdata/albertin_virgin_c.jpg");
}
