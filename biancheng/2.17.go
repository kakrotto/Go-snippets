package main

import "fmt"

//将枚举值转换为字符串
// 声明芯片类型
type ChipType int

const (
	None ChipType = iota
	CPU    // 中央处理器
	GPU    // 图形处理器
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}

	return "N/A"
}

func main() {

	// 输出CPU的值并以整型格式显示
	fmt.Printf("%s %d", CPU, CPU)
	// String() 方法的 ChipType 在使用上和普通的常量没有区别。当这个类型需要显示为字符串时，Go语言会自动寻找 String() 方法并进行调用。
}