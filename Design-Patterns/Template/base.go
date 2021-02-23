package template

import "fmt"

type Base struct {
}

type Processor interface {
	Init()
	BeforeProcess()
	Process()
	AfterProcess()
}

func (b *Base) Init() {
	fmt.Println("Base.Init()")
}

func (b *Base) BeforeProcess() {
	fmt.Println("Base.BeforeProcess()")
}

func (b *Base) Process() {
	fmt.Println("Base.Process()")
}

func (b *Base) AfterProcess() {
	fmt.Println("Base.AfterProcess()")
}