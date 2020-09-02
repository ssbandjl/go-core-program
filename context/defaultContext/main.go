package main

/*
默认上下文：context 包中最常用的方法还是 context.Background、context.TODO，这两个方法都会返回预先初始化好的私有变量 background 和 todo，它们会在同一个 Go 程序中被复用
context.Background 是上下文的默认值，所有其他的上下文都应该从它衍生（Derived）出来；
context.TODO 应该只在不确定应该使用哪种上下文时使用
*/
