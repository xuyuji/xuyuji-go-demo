# Template

用Go实现模版模式demo

执行结果：

```
Base.Init()
BaseEmail.Init()
BatchTask.Init()
----------
BaseEmail.BeforeProcess()
BaseEmail.SendEmail()
BatchTask.BeforeProcess()
----------
Base.Process()
BaseEmail.Process()
BatchTask.SendEmail()
BatchTask.Process()
----------
Base.AfterProcess()
```

`BaseEmail`基于`Base`拓展了email功能，`BaseEmail`的子类`BatchTask`只需要实现email逻辑即可使用通用流程。

注意`EmailWapper`的使用，采用的匿名接口方式，在`BaseEmail`中，如果直接使用b.SendEmail()，无法调用`BatchTask`的SendEmail方法，这里的b是BaseEmail。需要调用匿名接口：b.EmailWapper.SendEmail()，不过这种方式需要给EmailWapper赋值，不赋值会空指针。