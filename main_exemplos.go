package main

import "fmt"

func hands_on() {
	// fmt.Println("Hello World")
	// var numeroTeste int32 = 12
	// numeroTeste := 12

	fmt.Println("Digite um número")
	var numero int32
	fmt.Scanf("%d", &numero)

	fmt.Println(fmt.Sprintf("Você digitou o número %d", numero))

}
