package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(1)
	l.PushBack("allen")
	l.PushBack(map[string]int{"allen": 1, "zhao": 2})
	fmt.Println(l.)
}
