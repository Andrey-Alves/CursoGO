//declarar pacoti principal
package main

//importar funcao
import "fmt"

//declaracao de variavel CONST
const ebulicaoF = 212.0

//funcao principal
func main(){

	//atribuir variavel  var nome = isso
	tempF := ebulicaoF //variavel do valor da temperatura em graus F
	
	//atribuir variavel  NomeDaVariavel := isso
	var tempC = (tempF - 32)*5/9 //conversao de F para C
	
	//
	fmt.Println("A temperatura de ebulicao da agua em F = ", tempF)
	//
	fmt.Println("A temperatura de ebulicao da agua em C = ", tempC)

}