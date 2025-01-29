package main

import (
	"fmt" //imported for IO opeartions
	// "math"
)

func main() {
	// fmt.Println("Hello world")

	// fmt.Println("go"+"lang");
	// fmt.Println(1+1)
	// fmt.Println(5/3)

	// fmt.Println(true && false)
	// fmt.Println(true || false)
	// fmt.Println(!true)

	//variables
	// var a="initial";
	// var a=true;
	// var b="hi";
	// var c=213;
	// fmt.Println(a,b,c)

	// var a int = 12
	// var b bool = true;
	// var c string = "string";
	// // c:=true
	// d:=123
	// // fmt.Println(a,b,c);
	// fmt.Println(a,b,d);
	// fmt.Println(c)
	// var a=5
	// a=8
	// fmt.Println(a)

	//CONSTANTS

	// constants of character, string, boolean, and numeric values.
	// const s string="constant-3"
	// var s string="bjb "
	// s="ugn"
	// fmt.Println(s)

	// const cons=4000
	// const cons=4e2
	// const cons=4e2/11
	// const cons = 4e6

	// fmt.Println(cons)
	// fmt.Println(int64(cons))
	// fmt.Println(math.Sin(1))

	// loops
	// i:=1
	// for i<=3 {
	//     fmt.Println(i);
	//     i=i+1;
	// }

	// for j:=0;j<=3;j++ {
	//     fmt.Println(j)
	// }

	// for i:= range 3 {
	//     fmt.Println("range",i)
	// }

	// for {
	//     fmt.Println("loop")
	//     break;
	// }

	// for n:=range 6{
	//     if n%2==0{
	//         continue
	//     }
	//     fmt.Println(n)
	// }

	// if else
	// if 7>2 {
	//     fmt.Println("yes")
	// }else{
	//     fmt.Println("no")
	// }
	// if 7<2 {
	//     fmt.Println("yes")
	// }else{
	//     fmt.Println("no")
	// }

	// if num:=5 ; num<0{
	//     fmt.Println("num is negative")
	// }else if num<7 {
	//     fmt.Println("num is less than 5")
	// }

	//switch case
	// i := 1
	// switch i {
	// case 1:
	// 	fmt.Print("one")
	// case 2:
	// 	fmt.Println("two")
	// }

	// i := 9
	// switch i {
	// case 1,2,4:
	// 	fmt.Print("one or two")
	// case 3:
	// 	fmt.Println("three")

	// default:
	//     fmt.Println("default")
	// }

	// whatAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case bool:
	// 		fmt.Println("bool")
	// 	case int:
	// 		fmt.Println("int")
	// 	default:
	// 		fmt.Printf("Don't know type %T\n", t)
	// 	}
	// }

	// whatAmI(true)
	// whatAmI(23)
	// whatAmI("hey")

	//arrays
	var a [5]int
	a[2] = 6
	// fmt.Println(a);
	fmt.Println(len(a))

	b := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(b)

	c := [...]int{1, 2, 3}
	fmt.Println(c)

	d := [...]int{100, 3: 400} //3-index, in between all will be zero
	fmt.Println(d)

	var e [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			e[i][j] = i + j
		}
	}

	// var f=[2][3]int{
	//     {1,2,3},
	//     {4,5,6},
	// }
	// var f=[2][3]int{
	//     {1,2,3},
	//     {4,5,6},
	// }

	f := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(f)

}
