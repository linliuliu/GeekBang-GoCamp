///这里使用用有个报错 手动改了本地learn.go 的go.mod require tag
//PS D:\GoProject\src\github.com\linliuliu\GeekBang-GoCamp\homework2> go mod tidy
//go: found learn.go/chapter02/015.fatrate.refactor/calc in learn.go v0.0.0-00010101000000-000000000000
//go: github.com/linliuliu/GeekBang-GoCamp/homework2/modulework imports
//learn.go/chapter02/015.fatrate.refactor/calc: module D:\GoProject\src\github.com\armstrongli\learn.go: parsing ..\..\..\armstrongli\learn.go\go.mod: D:\GoP
//roject\src\github.com\armstrongli\learn.go\go.mod:6:2: require github.com/armstrongli/go-bmi: version "v0.0.0-00010101000000" invalid: must be of the form v1.2.3
////
//
//PS D:\GoProject\src\github.com\linliuliu\GeekBang-GoCamp\homework2> go mod tidy
//go: found learn.go/chapter02/015.fatrate.refactor/calc in learn.go v0.0.0-00010101000000-000000000000
//PS D:\GoProject\src\github.com\linliuliu\GeekBang-GoCamp\homework2> go mod vendor
//
/////	//
package main

import (
	"fmt"
	calc "learn.go/chapter02/015.fatrate.refactor/calc" //注意 包名= module+path
)

func main(){
	fmt.Println("hello module")
	bmi ,_:= calc.CalcBMI(1.7,63)
	calc.CalcFatRate(bmi,28,"男")

}


