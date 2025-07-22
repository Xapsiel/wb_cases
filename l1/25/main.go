package main

import (
	"fmt"
	"time"
)

func sleep(duration int) {
	<-time.After(time.Duration(duration) * time.Second)
}
func main() {
	fmt.Println(time.Now())
	fmt.Println("Засыпаем")
	sleep(2)
	fmt.Println("Продолжаем работу")
	fmt.Println(time.Now())
}
