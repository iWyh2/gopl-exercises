package ch1

import (
	"fmt"
	"time"
)

func PerformanceTest() {
	fmt.Println("Now, we will test the performance of the two Echo program")
	fmt.Println("First, testing the Echo function")
	time1 := time.Now()
	Echo()
	time2 := time.Now()
	fmt.Println("We get the time of the Echo program's performance: ",
		time2.Sub(time1))
	fmt.Println("Then, testing the EchoPlus function")
	time1 = time.Now()
	EchoPlus()
	time2 = time.Now()
	fmt.Println("We get the time of the EchoPlus program's performance: ",
		time2.Sub(time1))
}
