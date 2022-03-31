package main

import "fmt"

// make( []Type, size, cap )

func iter(m map[string]string){
	for k, v := range m{
		fmt.Printf("key: %s value: %s\n", k, v)
	}
	/*
		for k := range m{
			这样写法 k 是 map 的 key
		}
	 */
}


func main() {
	// 第一种初始化方法
	var map1 map[string]int
	map1 = map[string]int{"one": 1, "two": 2}
	// 第二种
	map3 := make(map[string]float32)
	map3 = map[string]float32{"a": 5.5}
	// 第三种
	map4 := map[string]float32{"a": 2.11, "b": 2.33}

	fmt.Printf("%p\n", &map4)

	// mapCreated := make(map[string]float) 等价于 mapCreated := map[string]float{}
	fmt.Println(map1)
	fmt.Println(map3)
	fmt.Println(map4)

	// 切片作为 map 的 value
	m1 := make(map[int][]int)
	m2 := make(map[int]*[]int)

	fmt.Println(m1)
	fmt.Println(m2)

	mm := map[string]string{"name": "mm", "sex": "f", "age": "18"}
	iter(mm)
	delete(mm, "name")
	delete(mm, "age")
	delete(mm, "sex")
	fmt.Println(mm)
	// 清空 map 的唯一办法就是重新 make 一个新的 map，不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。
}
