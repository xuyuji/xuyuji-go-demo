package template

import "fmt"

type BaseEmail struct {
	Base
	EmailWapper
}

type EmailWapper interface {
	SendEmail()
}

func (b *BaseEmail) Init() {
	b.Base.Init()
	fmt.Println("BaseEmail.Init()")
	b.EmailWapper = b	//这里不设置，b.EmailWapper.SendEmail()调用会报空指针
}

func (b *BaseEmail) BeforeProcess() {
	fmt.Println("BaseEmail.BeforeProcess()")
	b.SendEmail()
}

func (b *BaseEmail) Process() {
	b.Base.Process()
	fmt.Println("BaseEmail.Process()")
	b.EmailWapper.SendEmail()
}

func (b *BaseEmail) SendEmail() {
	fmt.Println("BaseEmail.SendEmail()")
}