package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生产者每秒生成一个字符串，并通过通道传给消费者，生产者使用两个 goroutine 并发运行，消费者在 main() 函数的 goroutine 中进行处理。

func producer(header string, channel chan <- string)  {
	for {
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

func customer(channel <- chan string)  {
	for {
		message := <- channel
		fmt.Println(message)
	}
}

func main()  {
	channel := make(chan string)
	go producer("cat", channel)
	go producer("dog", channel)
	go customer(channel)
	time.Sleep(time.Second * 5)
}
