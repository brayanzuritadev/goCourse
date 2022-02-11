package main

import "fmt"

func main() {
	//slicen
	//al tener inicialmente 3 datos la capacidad del slicen es de 3
	numeros := []int{1, 2, 300}

	fmt.Println(numeros, len(numeros), cap(numeros))

	//agregar datos al slicen
	//al aumentarse un dato mas al slicen
	//este sobre pasa su capacidad inical, por lo tanto se aumenta la capacidad
	//automaticamente al doble ===  capacidad = 6
	fmt.Printf("esta es la longitud %d esta es la capacidad %d y esta es la direccion en memoria %p\n", len(numeros), cap(numeros), numeros)
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)

	//en este punto se creo otro slicen al momento de realizar el append.
	fmt.Println(numeros, len(numeros), cap(numeros))
	fmt.Printf("esta es la longitud %d esta es la capacidad %d y esta es la direccion en memoria %p\n", len(numeros), cap(numeros), numeros)
	//sub slicen

	subSlicen := numeros[:2]
	fmt.Printf("esta es la longitud %d esta es la capacidad %d y esta es la direccion en memoria %p\n", len(subSlicen), cap(subSlicen), subSlicen)

	numeros[0] = 100

	//aun cuando se saca el subslicen antes de modificar lo que contiene el slicen padre
	//en el slicen hijo se modifica automaticamente los cmabios realizados en el slicen padre

	fmt.Println(subSlicen)
	fmt.Println(numeros)
	numeros[1] = 200
	fmt.Println(subSlicen)

	//caso interesante
	numero := []int{}
	fmt.Printf("%v %d %d %p\n", numero, cap(numero), len(numero), numero)
	numero = append(numero, 10)
	fmt.Println(numero, cap(numero), len(numero))
}
