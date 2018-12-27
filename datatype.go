package main

import (
	"fmt"
	"time"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(true))

	date := time.Now()
	fmt.Println("Time now is $1", date)
	fmt.Println(reflect.TypeOf(date))

	data := date.Format("20180901")
	fmt.Println(data)
	fmt.Println(reflect.TypeOf(data))

	layout := "20060102"
	str := "20180901"
	resultant, _ := time.Parse(layout, str)
	fmt.Println(resultant)
	fmt.Println(reflect.TypeOf(resultant))

}