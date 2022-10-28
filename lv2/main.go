package main

import (
	"fmt"
	"time"
)

var adjust = make(chan int, 1)

func main() {

	for {
		go 装弹()
		<-adjust
		go 瞄准()
		<-adjust
		go 发射()
		<-adjust //尽力了。。
	}

}

func 装弹() {
	time.Sleep(time.Second)

	adjust <- 1
	fmt.Print("装弹->")

}
func 瞄准() {
	time.Sleep(2 * time.Second)

	adjust <- 2
	fmt.Print("瞄准->")

}
func 发射() {
	time.Sleep(3 * time.Second)

	adjust <- 3
	fmt.Print("发射!\n")

}
