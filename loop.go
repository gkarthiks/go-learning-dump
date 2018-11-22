package main

import (
	"fmt"
	"strconv"
)

func main() {

    for i := 0; i < 10 ; i++ {
		fmt.Println(i * 3)
		b := checkthis(i * 3)
		println(b)
    }
}

func checkthis(i int) string {
	println(strconv.Itoa(i) + " <<---")
	return strconv.Itoa(i)
}