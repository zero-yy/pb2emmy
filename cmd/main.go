package main

import (
	"flag"
	"fmt"
	"github.com/zero-yy/pb2emmy"
)

var (
	inputDir = flag.String("inputdir", "", "")
	outputName = flag.String("output", "", "")
)

func main() {
	flag.Parse()
	if len(*inputDir) == 0 || len(*outputName) == 0 {
		panic("len(*inputDir) == 0 || len(*outputName) == 0")
	}
	c := &pb2emmy.Config{
		InputDir: *inputDir,
		OutputName: *outputName,
	}
	fmt.Printf("====>pb2emmy start: %v\n", c)
	//c := &pb2emmy.Config{
	//	InputDir: "./test2/",
	//	OutputName: "./output.lua",
	//}

	p := pb2emmy.NewPb2Emmy(c)
	p.Do()
	fmt.Println("====>pb2emmy done")
}
