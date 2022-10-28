package main

import (
	"fmt"
	"time"
)

func main() {
	杰哥不要啦()
}
func 打电动() {
	fmt.Println("输了啦，都是你害的")
}

func 欢迎来我家玩() string {
	//花费5s前往杰哥家
	time.Sleep(5 * time.Second)
	return "登dua郞"
}

type hhh func(func(), func()) string

func 杰哥不要啦() {
	打电动()
	好康的 := 欢迎来我家玩()
	fmt.Println(好康的) //回调函数？
}
