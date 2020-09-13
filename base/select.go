package main

import (
	"log"
	"sync"
	"time"
)

/*
select 语法类似switch， case 后面必须是通信，要么发送，要么接收
*/
func main() {
	var waitGroup sync.WaitGroup
	var num chan int = make(chan int, 0)
	var count int = 0
	waitGroup.Add(1)
	ticker := time.NewTicker(time.Second * 1)
	go func(ticker *time.Ticker) {
		for {
			num <- count
			count++
			log.Println("ticker")
		}
	}(ticker)

	go func() {
		for true {
			select {
			case <- num:
				log.Println("num")
				waitGroup.Done()
			default:
				log.Println("default")
			}
			time.Sleep(time.Second * 1)
		}
	}()
	waitGroup.Wait()
}
