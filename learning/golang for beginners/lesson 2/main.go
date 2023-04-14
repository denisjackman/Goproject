package main
import (
	"fmt"
	"reflect"
	"strconv"
)
// This is lesson 2 in the course data types
var newname string = "Lisa"
const PI float64 = 3.14159265358979323846
func main() {
/*
	name:= "John"
	var s string = "Hello World"
	var i int = 100
	var b bool = false
	var f float32 = 10.5
	var ff float64 = 77.90 
*/
	const denis string = "Denis"
	const isIrish = true
	const age = 12 
	fmt.Println(newname)
	fmt.Print("Hello World\n")
	var greeting string = "Hello World\n"
	fmt.Print(greeting)
	var city string = "Kolkata\n"
	fmt.Print(city)
	var name string = "KodeKloud"
	var user string = "John"
	fmt.Print("Welcome to ", name, ", ", user,"\n")
	fmt.Print(name)
	fmt.Print(user)
	fmt.Println("Welcome to ", name, ", ", user)
	fmt.Printf("Nice to see you here, at %v\n", name)
	var grades int = 42 
	fmt.Printf("You have scored %d out of 100\n", grades)
	var i int = 78 
	fmt.Printf("Hey, %v! you have scored %d/100 in Physics\n", user,  i)
	var s,t string = "foo", "bar"
	fmt.Println(s)
	fmt.Println(t)
	var (
		a int = 10
		st string = "Hello"
	)
	fmt.Println(a)
	fmt.Println(st)
	sss := "Hello World"
	fmt.Println(sss)
	newcity:= "London"
	{
		country:= "UK"
		fmt.Println(newcity)
		fmt.Println(country)
	}
	fmt.Println(newcity)
	// zero values 
	var as int
	fmt.Printf("%d\n", as)
	var fl float64
	fmt.Printf("%.2f\n", fl)
	//var username string
	//fmt.Print("Enter your name: ")
	//fmt.Scanf("%s", &username)
	//fmt.Println("Hello there,", username)
	var magicname string
	var is_muggle bool
	fmt.Print("Enter your name & are you a muggle: ")
	fmt.Scanf("%s %t", &magicname, &is_muggle)
	fmt.Println(magicname, is_muggle)

	var newgrades int = 42
	var message string = "hello world"
	var isCheck bool = true
	var amount float64 = 5466.54

	fmt.Printf("variable grades = %v is of type %T \n", newgrades, newgrades)
	fmt.Printf("variable message = '%v' is of type %T \n", message, message)
	fmt.Printf("variable isCheck = '%v' is of type %T \n", isCheck, isCheck)
	fmt.Printf("variable amount = %v is of type %T \n", amount, amount)
	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("priyanka"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(46.0))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))

	fmt.Printf(("variable newgrades=%v is of type %v \n"), newgrades, reflect.TypeOf(newgrades))
	fmt.Printf(("variable message='%v' is of type %v \n"), message, reflect.TypeOf(message))

	var ii int = 90
	var ff float64 = float64(ii)
	fmt.Printf("%.2f\n", ff)
	var ss string = strconv.Itoa(ii)
	fmt.Printf("%q\n", ss)
	var asss string = "100"
	iii, err := strconv.Atoi(asss)
	fmt.Printf("%v, %T\n", iii, iii)
	fmt.Printf("%v, %T\n", err, err)

	var aasss string = "100abc"
	iiii, err := strconv.Atoi(aasss)
	fmt.Printf("%v, %T\n", iiii, iiii)
	fmt.Printf("%v, %T\n", err, err)
	fmt.Printf("%v, %T\n", denis, denis)
	fmt.Printf("%v, %T\n", isIrish, isIrish)
	fmt.Printf("%v, %T\n", age, age)
	var radius float64 = 5.0 
	var area float64 = PI * radius * radius
	fmt.Println("Area of circle is", area)
}

