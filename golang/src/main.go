package main

import (
		"fmt"
	)

func main(){
	
}

/*
 * Reto #0
 * Enunciado: Escribe un programa que muestre por consola (con un print) los números de 1 a 100 (ambos incluidos y con un salto de línea entre cada impresión), sustituyendo los siguientes:
 * - Múltiplos de 3 por la palabra "fizz".
 * - Múltiplos de 5 por la palabra "buzz".
 * - Múltiplos de 3 y de 5 a la vez por la palabra "fizzbuzz".
 */

func Fizzbuzz(){
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
 * Reto #1
 * Enunciado: Escribe una función que reciba dos palabras (String) y retorne verdadero o falso (Boolean) según sean o no anagramas.
 * Un Anagrama consiste en formar una palabra reordenando TODAS las letras de otra palabra inicial.
 * NO hace falta comprobar que ambas palabras existan.
 * Dos palabras exactamente iguales no son anagrama.
 */

func IsAnagrama(unaPalabra string, otraPalabra string) bool{
	var auxiliar = []rune(unaPalabra)
	var length int = len(auxiliar) - 1
	for i, j := 0, length; i < j; i, j = i+1, j-1 {
		auxiliar[i], auxiliar[j] = auxiliar[j], auxiliar[i]
	}
	return string(auxiliar) == otraPalabra
}

/*
 * Reto #2
 * Enunciado: Escribe un programa que imprima los 50 primeros números de la sucesión de Fibonacci empezando en 0.
 * La serie Fibonacci se compone por una sucesión de números en la que el siguiente siempre es la suma de los dos anteriores.
 * 0, 1, 1, 2, 3, 5, 8, 13...
	0 1  2  3  4  5  6  7
 */

func Fibonacci()  {
	var n1, n2, n int = 0, 0, 0
	for i := 0; i < 50; i++ {
		fmt.Println(n)
		n2 = n1
		n1 = n
		n = n2 + n1
		if i == 0 {
			n = 1
		}
	}
}

/*
 * Reto #3
 * Enunciado: Escribe un programa que se encargue de comprobar si un número es o no primo.
 * Hecho esto, imprime los números primos entre 1 y 100.
 */

func IsPrimo(unNumero int) bool{
	var divisores int = 2
	for i := unNumero; i > 0; i-- {
		if unNumero % i == 0 {
			divisores--
		} 
	}
	return divisores == 0
}

func PrimosEntreDosNumeros(desdeNumero int, hastaNumero int)  {
	for i := desdeNumero; i <= hastaNumero; i++ {
		if IsPrimo(i){
			fmt.Println(i)
		}
	}
}

/*
 * Enunciado: Crea UNA ÚNICA FUNCIÓN (importante que sólo sea una) que sea capaz de calcular y retornar el área de un polígono.
 * - La función recibirá por parámetro sólo UN polígono a la vez.
 * - Los polígonos soportados serán Triángulo, Cuadrado y Rectángulo.
 * - Imprime el cálculo del área de un polígono de cada tipo.
 */

type FormaGeometrica interface {
	Area() float64
}


type Triangulo struct{
	base float64
	altura float64
}

func (t Triangulo) Area() float64 {
	return t.altura * t.base / 2
}

type Cuadrado struct{
	lado float64
}

func (c Cuadrado) Area() float64  {
	return c.lado * c.lado
}

type Rectángulo struct{
	base float64
	altura float64
}

func (r Rectángulo) Area() float64 {
	return r.altura * r.base
}

func CalcularArea(forma FormaGeometrica) {
	fmt.Println("El area de la forma es: ", forma.Area())
}

