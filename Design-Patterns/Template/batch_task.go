package template

import "fmt"

type BatchTask struct {
	BaseEmail
}

func (b *BatchTask) Init() {
	b.BaseEmail.Init()
	fmt.Println("BatchTask.Init()")
	b.EmailWapper = b
}

func (b *BatchTask) BeforeProcess() {
	b.BaseEmail.BeforeProcess()
	fmt.Println("BatchTask.BeforeProcess()")
}

func (b *BatchTask) Process() {
	b.BaseEmail.Process()
	fmt.Println("BatchTask.Process()")
}

func (b *BatchTask) SendEmail() {
	fmt.Println("BatchTask.SendEmail()")
}