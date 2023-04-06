package ch1

import (
	"fmt"
	"os"
)

func EchoPlus() {
	for i, v := range os.Args[0:] {
		fmt.Println("index: ", i, "command: ", v)
	}
}
