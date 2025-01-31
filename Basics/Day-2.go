// package main

// import (
// 	"fmt"
// 	// "unicode/utf8"
// )

// "fmt"
// "reflect"

//struct
// type Student struct{
// 	Name string
// 	age int
// }

// type Animal struct {
// 	Name string
// }

// type Bird struct{
// 	Name string
// 	Speed int
// }

// type Bird struct {
// 	Animal
// 	Speed int
// }

// tags and reflection
// type Animal struct {
// 	Name string `required max_len:"10"`
// 	Age  int    `tag has any value`
// }

//closures
// func outer() func() int{
// 	x:=10
// 	return func() int {
// 		x++;
// 		return x;
// 	}
// }

// func main() {
// fmt.Println("hi")

// Utkarsh:=Student{
// 	Name: "utkarsh",
// 	age: 22,
// }
// a:=2
// Utkarsh.Name="rohan"
// fmt.Println(Utkarsh,a)

// animal:=Animal{
// 	Name: "dog",
// }

// animal:=Bird{
// 	Name: "parrot",
// 	Speed: 23,
// }

// animal:=Bird{
// Animal: Animal{"parrot"},
// Speed: 23,
// }
// animal.Name="bird-2"

// fmt.Println(animal)
// a:=2
// b:=a

// b=4
// fmt.Println(animal.Name,a,b)

// kiwi := struct {
// 	name string
// 	age  int
// }{
// 	name: "utkarsh",
// 	age:  2,
// }

// // b:=kiwi
// b := &kiwi
// b.age = 23
// fmt.Println(kiwi, *b)

// peacock:=Bird{
// 	Animal: Animal{"Utkarsh"},
// 	Speed: 23,
// }
// peacock2:=Bird{
// 	Animal: Animal{"tkarsh"},
// 	Speed: 23,
// }

//maps with composite keys
// mp:=map[Bird]int {
// 	peacock:3,
// }
// fmt.Println(mp[peacock2])
// // a:=3
// // fmt.Println("type of %T", peacock)
// fmt.Printf("type of %T", peacock)

//tags and reflection
// fmt.Println("struct tags")
// a := Animal{"abc", 23}
// // fmt.Println(a)

// t:=reflect.TypeOf(a)
// // field,_:=t.FieldByName("Age")
// field,_:=t.FieldByName("Name")
// // fmt.Println(t)
// fmt.Println(field.Tag)
// fmt.Println("hi")

//closure example
// closurefn:=outer();
// fmt.Println(closurefn())
// fmt.Println(closurefn())
// fmt.Println(closurefn())

//Strings and Runes

// const s="สวัสดี";
// const s="abc";
// fmt.Println(len(s))

// for i:=0;i<len(s);i++{
// 	fmt.Printf("%x ",s[i]) // x is used for format a value as a hexadecimal value
// }

// fmt.Println("Rune count: ",utf8.RuneCountInString(s))

// for idx,runeVal :=range s{
// fmt.Printf("%#U    ",runeVal)   //%#U for printing unicode code point
// have a diff of 3 in index because index tell the byte it start with, and each rune in this is of 3 byte
// fmt.Println(idx)
// idx++;
// }

// }

package main

import "fmt"

type Animal struct {
	Species string
}

func (a Animal) speak() {
	fmt.Printf("%s is the species of Animal\n", a.Species)
}

type Person struct {
	Name string
}

func (a Person) speak() {
	fmt.Printf("%s is the name of person\n", a.Name)
}

type Speak interface {
	speak()
}

func introduce(a Speak) {
	a.speak()
}

func main() {
	fmt.Println("Interfaces in golang")

	animal := Animal{
		"Lion",
	}
	person := Person{
		"Utkarsh",
	}

	introduce(animal)
	introduce(person)

}
