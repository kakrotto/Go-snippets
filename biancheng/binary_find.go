package main

import "fmt"

func BinaryFind(arr *[]int, left, right, val int) {
	if left > right {
		fmt.Println("找不到")
		return
	}
	mid := (left + right) / 2

	if (*arr)[mid] > val {
		BinaryFind(arr, left, mid-1, val)
	}else if (*arr)[mid] < val {
		BinaryFind(arr, mid+1, right, val)
	}else {
		fmt.Printf("找到了，下标为：%v\n", mid)
	}
}

func main() {
	arr := []int{1, 2, 5, 7, 15, 25, 30, 36, 39, 51, 67, 78, 80, 82, 85, 91, 92, 97}
	BinaryFind(&arr, 0, len(arr) - 1, 30)
	fmt.Println("main arr = ",arr)

}