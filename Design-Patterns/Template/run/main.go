package main

import (
	"fmt"
	"xuyuji/template"
)

func main()  {
	p := getProcessor()
	run(p)
}

func getProcessor() template.Processor {
	return new(template.BatchTask)
}

func run(p template.Processor)  {
	p.Init()
	fmt.Println("----------")
	p.BeforeProcess()
	fmt.Println("----------")
	p.Process()
	fmt.Println("----------")
	p.AfterProcess()
}