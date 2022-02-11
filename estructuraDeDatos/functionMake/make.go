package main

import "fmt"

func main() {

	//me permite crear un slicen con longitud y capacidad definida
	numeros := make([]int, 2, 3)
	numeros[0] = 100
	numeros[1] = 100
	fmt.Printf("direccion en memoria es %p\n", numeros)
	numeros = append(numeros, 400) //esta es la manera de agregar un dato a un slicen
	fmt.Printf("direccion en memoria es %p\n", numeros)
	fmt.Println(numeros, len(numeros), cap(numeros))
}
