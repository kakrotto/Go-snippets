package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	var scene sync.Map
	go func() {
		for i:=0; i < 50; i++{
			scene.Store(strconv.Itoa(i), i)
		}
	}()
	go func() {
		for i:=0; i < 50;i++{
			scene.Delete(strconv.Itoa(i))
		}
	}()

	go func() {
		for {
			scene.Range(func(k, v interface{}) bool {
				fmt.Println(k ,v)
				return true
			})
		}

	}()
	time.Sleep(10000000000)
}
