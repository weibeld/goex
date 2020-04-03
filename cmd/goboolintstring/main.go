// goboolintstring is a tool for printing a boolean, an integer, and a string.
//
// Usage:
//
//  goboolintstring
//
package main

import (
	"fmt"
	"github.com/weibeld/goex/mybool"
	"github.com/weibeld/goex/myint"
	"github.com/weibeld/goex/mystring"
)

func main() {
	fmt.Println(mybool.Get())
	fmt.Println(myint.Get())
	fmt.Println(mystring.Get())
}
