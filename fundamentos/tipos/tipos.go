package main

import (
	"fmt"
	"reflect"
	"math"
)

func main() {
	fmt.Println(1,2,1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	var b byte = 255
	fmt.Println("Byte inteiro é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("Maximo de", i1)
	fmt.Println("Tipo do i1 e", reflect.TypeOf(i1))
}