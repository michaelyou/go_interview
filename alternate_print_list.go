/*
问题描述

使用两个 goroutine 交替打印序列，一个 goroutinue 打印数字， 另外一个goroutine打印字母，
最终效果如下 12AB34CD56EF78GH910IJ 。

解题思路

问题很简单，使用 channel 来控制打印的进度。使用两个 channel ，来分别控制数字和字母的打印序列，
数字打印完成后通过 channel 通知字母打印, 字母打印完成后通知数字打印，然后周而复始的工作。
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	chan_n := make(chan bool, 1)
	chan_c := make(chan bool)
	done := make(chan bool)

	go func() {
		for i := 1; i < 11; i += 2 {
			<-chan_n
			fmt.Print(i)
			fmt.Print(i + 1)
			chan_c <- true
		}
	}()

	go func() {
		string_list := [10]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
		for i := 0; i < 10; i += 2 {
			<-chan_c
			fmt.Print(string_list[i])
			fmt.Print(string_list[i+1])
			chan_n <- true
		}
		fmt.Print("\n")
		done <- true
	}()

	chan_n <- true
	<-done
}
