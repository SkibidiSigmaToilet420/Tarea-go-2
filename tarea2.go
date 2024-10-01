package main

import (
	"fmt"
	"strings"
)

func main() {

	//funcion Contains verifica si una cadena de caracteres contiene a otra responde con true o false

	fmt.Println("Contains")
	fmt.Println(("1)"), strings.Contains("Amarillo", "Ama")) //respuesta true porque Amarillo contiene Ama

	fmt.Println(("2)"), strings.Contains("Rojo", "Azul")) //respuesta false porque Rojo no contien la cadena de caracteres Azul

	//Funcion Count cuenta cuantas letras hay en una cadena de caracteres
	fmt.Println("Count")
	fmt.Println("1)", strings.Count("Rojo", "o"))   //tiene 2 "o" en rojo
	fmt.Println("2)", strings.Count("Morado", "u")) //no hay "u" en morado

	//Funcio HasPrefix verifica si una palabra comienza por otra cadena
	fmt.Println("HasPrefix")
	fmt.Println("1)", strings.HasPrefix("carro", "ca")) //respuesta true porque carro empieza por "ca"
	fmt.Println("2)", strings.HasPrefix("carro", "te")) //respuesta false porque carro no empieza por te

	//Funcio HasSuffix verifica si una palabra comienza por otra cadena
	fmt.Println("HasSuffix")
	fmt.Println("1)", strings.HasSuffix("Perro", "ro")) //respuesta true porque perro termina en "ro"
	fmt.Println("2)", strings.HasSuffix("gato", "mu"))  //respuesta false porque gato no termina en mu

}
