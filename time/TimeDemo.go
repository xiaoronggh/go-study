package main

import (
	"log"
	"sync"
	"time"
)

/*
timer 定时器，固定时间才会执行 timer 需要配置Reset函数才能实现定时器，Reset可能会失败
ticker 定时器，跟timer的区别是，定义完成就会立即执行
*/

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	var ticker *time.Ticker = time.NewTicker(time.Second * 2)
	go func(ticker *time.Ticker){
		defer waitGroup.Done()
		for {
			<- ticker.C
			log.Println("it is time ticker")
		}
	}(ticker)

	var timer *time.Timer = time.NewTimer(time.Second * 2)

	go func(timer *time.Timer){
		defer waitGroup.Done()
		for {
			<- timer.C
			log.Println("it is time Timer")
			timer.Reset(time.Second * 2)
		}
	}(timer)
	log.Println("start...")
	waitGroup.Wait()
}
