package ch1

import (
	"fmt"
	"os"
	"strings"
)

func Echo() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
