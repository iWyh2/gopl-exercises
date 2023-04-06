package ch2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	1m = 3.28084 ft
	1 ft = 0.3048 m
*/

type Foot float64
type Meter float64

func MtoF(m Meter) Foot {
	return Foot(m * 3.28084)
}

func FtoM(f Foot) Meter {
	return Meter(f * 0.3048)
}

func (f Foot) String() string {
	return fmt.Sprintf("%gft", f)
}
func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

// LengthTransfer Convert meters to feet or convert feet to meters
func LengthTransfer() {
	if os.Args != nil && len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			l, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				_, err := fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				if err != nil {
					fmt.Println("Have Error!")
				}
				os.Exit(1)
			}
			fmt.Printf("%s = %s, %s = %s\n",
				Foot(l), FtoM(Foot(l)), Meter(l), MtoF(Meter(l)))
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			l, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				_, err := fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				if err != nil {
					fmt.Println("Have Error!")
				}
				os.Exit(1)
			}
			fmt.Printf("%s = %s, %s = %s\n",
				Foot(l), FtoM(Foot(l)), Meter(l), MtoF(Meter(l)))
		}
	}
}
