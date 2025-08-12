package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

// seleccionarPalabra elige una palabra al azar de la lista.
func seleccionarPalabra() string {
	diccionario := []string{
		"golang", "programacion", "computadora", "desarrollo",
		"teclado", "algoritmo", "variable", "funcion",
		"paquete", "interfaz", "estructura", "puntero",
	}
	// Inicializa el generador de números aleatorios para que la palabra sea diferente cada vez.
	rand.Seed(time.Now().UnixNano())
	return diccionario[rand.Intn(len(diccionario))]
}

// mostrarEstadoJuego imprime el estado actual del juego en la consola.
func mostrarEstadoJuego(palabraOculta string, letrasAdivinadas map[rune]bool, intentosFallidos []string, vidas int) {
	fmt.Println("\n" + "=".repeat(30))

	// Muestra el dibujo del ahorcado según las vidas restantes.
	mostrarAhorcado(vidas)

	// Construye y muestra la palabra con guiones bajos para las letras no adivinadas.
	fmt.Print("Palabra: ")
	for _, letra := range palabraOculta {
		if letrasAdivinadas[letra] {
			fmt.Printf("%c ", letra)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()

	// Muestra información adicional.
	fmt.Printf("Vidas restantes: %d\n", vidas)
	fmt.Printf("Letras incorrectas: %s\n", strings.Join(intentosFallidos, ", "))
	fmt.Println("=".repeat(30))
}

// procesarAdivinanza valida y procesa la letra ingresada por el usuario.
// Devuelve un booleano indicando si la letra era correcta.
func procesarAdivinanza(palabraOculta string, letra rune, letrasAdivinadas map[rune]bool, intentosFallidos *[]string) bool {
	letra = unicode.ToLower(letra) // Convertir a minúscula para ser consistente.

	// Verificar si la letra ya fue intentada (correcta o incorrecta).
	if letrasAdivinadas[letra] {
		fmt.Printf("Ya adivinaste la letra '%c'. Intenta con otra.\n", letra)
		return true // No se penaliza, pero no es un nuevo acierto.
	}
	for _, fallida := range *intentosFallidos {
		if string(letra) == fallida {
			fmt.Printf("Ya intentaste con la letra '%c' y es incorrecta. Intenta con otra.\n", letra)
			return false // No se penaliza, pero sigue siendo un fallo.
		}
	}

	// Comprobar si la letra está en la palabra.
	if strings.ContainsRune(palabraOculta, letra) {
		letrasAdivinadas[letra] = true
		fmt.Printf("¡Bien hecho! La letra '%c' está en la palabra.\n", letra)
		return true
	} else {
		*intentosFallidos = append(*intentosFallidos, string(letra))
		fmt.Printf("¡Incorrecto! La letra '%c' no está en la palabra.\n", letra)
		return false
	}
}

// verificarVictoria comprueba si todas las letras de la palabra han sido adivinadas.
func verificarVictoria(palabraOculta string, letrasAdivinadas map[rune]bool) bool {
	for _, letra := range palabraOculta {
		if !letrasAdivinadas[letra] {
			return false // Si falta al menos una letra por adivinar, no ha ganado.
		}
	}
	return true
}

// mostrarAhorcado imprime el arte ASCII correspondiente al número de vidas.
func mostrarAhorcado(vidas int) {
	etapas := []string{
		// 0 vidas (final)
		`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
		// 1 vida
		`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`,
		// 2 vidas
		`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
		// 3 vidas
		`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
		// 4 vidas
		`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
		// 5 vidas
		`
  +---+
  |   |
  O   |
      |
      |
      |
=========`,
		// 6 vidas (inicio)
		`
  +---+
  |   |
      |
      |
      |
      |
=========`,
	}
	// El número de vidas corresponde a un índice en el array de etapas.
	if vidas >= 0 && vidas < len(etapas) {
		fmt.Println(etapas[len(etapas)-1-vidas])
	}
}

func main() {
	var nombre string
	vidas := 6
	palabraOculta := seleccionarPalabra()
	letrasAdivinadas := make(map[rune]bool)
	var intentosFallidos []string

	fmt.Println("¡Bienvenido al Juego del Ahorcado en Go!")
	fmt.Print("Por favor, escribe tu nombre: ")
	fmt.Scanln(&nombre)
	fmt.Printf("Hola, %s. ¡Adivina la palabra!\n", nombre)
	time.Sleep(1 * time.Second)

	// Bucle principal del juego
	for vidas > 0 {
		mostrarEstadoJuego(palabraOculta, letrasAdivinadas, intentosFallidos, vidas)

		// Pedir letra al usuario
		fmt.Print("Ingresa una letra: ")
		var entrada string
		fmt.Scanln(&entrada)

		if len(entrada) == 0 {
			fmt.Println("No ingresaste nada. Inténtalo de nuevo.")
			continue
		}

		// Procesar la primera letra de la entrada
		letra := rune(entrada[0])
		if !procesarAdivinanza(palabraOculta, letra, letrasAdivinadas, &intentosFallidos) {
			vidas--
		}

		// Comprobar si el jugador ha ganado
		if verificarVictoria(palabraOculta, letrasAdivinadas) {
			fmt.Printf("\n¡FELICIDADES, %s! Has adivinado la palabra: %s\n", nombre, palabraOculta)
			return // Termina el juego
		}
	}

	// Si el bucle termina, el jugador ha perdido
	mostrarAhorcado(0) // Muestra el dibujo final
	fmt.Printf("\n¡Oh no, %s! Te has quedado sin vidas.\n", nombre)
	fmt.Printf("La palabra correcta era: %s\n", palabraOculta)
}
