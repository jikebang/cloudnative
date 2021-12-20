package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int,1)
	go func() {
		defer func() {
			fmt.Println("go fun close")
		}()
		/*for  {
			select {
			case a:=<-ch:
				fmt.Println(a)
			default:
				fmt.Println("is default")
				time.Sleep(800*time.Millisecond)
			}
		}*/
		/*for ii := range ch{
			fmt.Println(ii)
		}*/
		for {
			fmt.Println(<-ch)
		}
	}()
	for i:=0;i<3;i++ {
		ch<-i
		time.Sleep(time.Second)
	}
	close(ch)
	time.Sleep(3*time.Second)
}