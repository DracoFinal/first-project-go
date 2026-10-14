package main

import "fmt"

func main() {
	Array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // declaro o array de 10 posições predefinidas, tipagem int
	for _, value := range Array {                   // uso o comando forrange, omitindo o valor das posições (index)
		if value%2 == 0 { // crio outro if com a condição quê se o resto do "valor" dividido por 2 for igual a zero,
			// ele irá realizar as condições
			fmt.Println("É par") // se a condição da divisão igual a 0 for verdaeira, irá printar a mensagem escrita.
		} else { // se a condição do if for falsa, ele irá realizar os resultados do "else"
			fmt.Println("É ímpar") // resultado do else
		} // fecha scopo do else
	} // fecha scopo do for range
} // fecha o scopo da func main
