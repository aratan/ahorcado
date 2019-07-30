package main

import (
	"fmt"
	//"strconv"
	"strings"
	"time"
)

func main() {
	var lee string
	var adivina string

	fmt.Println("juego del ahorcado. \n Escriba su nombre.")
	fmt.Scanln(&lee)
	fmt.Println("Hola " + lee)
	time.Sleep(1 * time.Second)
	fmt.Println("Comienza a adivinar...")
	

	dic := []string{"apple", "banana"}
var log bool	
fmt.Println("Tiene una longitud de: ", len(dic[0]))

	for vida := 5; vida > 0; {

		fmt.Scanln(&adivina)
		log = strings.Contains(dic[0], adivina)
		fmt.Println(adivina + " es: ")


		if log == true {
			fmt.Println("Contiene la letra: ", adivina)
			if adivina == dic[0]{
				fmt.Println("Ganaste!!!")
				break
			}
		} else {
			vida = vida - 1
			fmt.Println("No tiene: ", adivina, " tu vida es: ", vida)

		}
		
		
	}
}
