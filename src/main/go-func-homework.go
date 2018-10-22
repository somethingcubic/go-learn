package main

import "fmt"

func main() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		fmt.Println("第", i, "步开始")
		//在defer定义的时候，i已经做了值拷贝了
		defer fmt.Println("defer i = ", i)
		//同下
		defer func() { fmt.Println("defer_closure i = ", i) }()
		//i 匿名函数中 引用拷贝 来自于for循环中的i的引用地址
		fs[i] = func() { fmt.Println("closure i = ", i) }
		fmt.Println(fs[i])
		fmt.Println("第", i, "步结束")
	}
	fmt.Println("分隔符")

	for _, f := range fs {
		f()
	}
}
