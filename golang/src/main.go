package main

import (
		"fmt"
	)

func main(){
	primosEntreDosNumeros(1, 100)
}

/*
 * Reto #0
 * Enunciado: Escribe un programa que muestre por consola (con un print) los números de 1 a 100 (ambos incluidos y con un salto de línea entre cada impresión), sustituyendo los siguientes:
 * - Múltiplos de 3 por la palabra "fizz".
 * - Múltiplos de 5 por la palabra "buzz".
 * - Múltiplos de 3 y de 5 a la vez por la palabra "fizzbuzz".
 */

func fizzbuzz(){
	for i := 1; i < 101; i++ {
		var multiploOfThree bool = i%3 == 0
		var multiploOfFive bool = i%5 == 0
		if multiploOfFive && multiploOfThree {
			fmt.Println("fizzbuzz")
		}else if multiploOfFive {
			fmt.Println("buzz")
		}else if multiploOfThree {
			fmt.Println("buzz")
		}else{
			fmt.Println(i)
		}
	}
}

/*
 * Enunciado: Escribe una función que reciba dos palabras (String) y retorne verdadero o falso (Boolean) según sean o no anagramas.
 * Un Anagrama consiste en formar una palabra reordenando TODAS las letras de otra palabra inicial.
 * NO hace falta comprobar que ambas palabras existan.
 * Dos palabras exactamente iguales no son anagrama.
 */

func isAnagrama(unaPalabra string, otraPalabra string) bool{
	var auxiliar = []rune(unaPalabra)
	var length int = len(auxiliar) - 1
	for i, j := 0, length; i < j; i, j = i+1, j-1 {
		auxiliar[i], auxiliar[j] = auxiliar[j], auxiliar[i]
	}
	return string(auxiliar) == otraPalabra
}

/*
 * Reto #3
 * Enunciado: Escribe un programa que se encargue de comprobar si un número es o no primo.
 * Hecho esto, imprime los números primos entre 1 y 100.
 */

func isPrimo(unNumero int) bool{
	var divisores int = 2
	for i := unNumero; i > 0; i-- {
		if unNumero % i == 0 {
			divisores--
		} 
	}
	return divisores == 0
}

func primosEntreDosNumeros(desdeNumero int, hastaNumero int)  {
	for i := desdeNumero; i <= hastaNumero; i++ {
		if isPrimo(i){
			fmt.Println(i)
		}
	}
}