package main
import "fmt"

var a int = 20
var b string = "Andrey"

func main(){

	//convertendo a variavel "a" tipo int, em variavel "c" tipo float
	var c float64 = float64(a)
	
	fmt.Println(c)
}