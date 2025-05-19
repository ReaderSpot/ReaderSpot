/*
Autor Nesto Gavilanes
Objetivos del sistema de gestion de libros electronicos:

	Agregar nuevos libros electrónicos.
	Buscar libros por título, autor o género.
	Eliminar o actualizar registros de libros existentes.
	Marcar libros como “leídos”, “pendientes” o “favoritos”.

Librerias a usar:

	fmt
	database/sql mysql
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Bienvenido al sistema de gestión de Libros Electrónicos")
		fmt.Println("Escribe el número de la opción:")
		fmt.Println("1. Registrar un libro")
		fmt.Println("2. Buscar un libro")
		fmt.Println("3. Administrar libros")
		fmt.Println("4. Salir")
		fmt.Print("Opción: ")

		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			registro_libros_func()
		case "2":
			busqueda_libros_func()
		case "3":
			administrar_libros_func()
		case "4":
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opción no válida. Intenta nuevamente.")
		}

		fmt.Println() // espacio entre iteraciones
	}
}
