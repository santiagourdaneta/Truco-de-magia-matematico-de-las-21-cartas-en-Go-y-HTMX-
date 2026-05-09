package main

import (
	"fmt"
	"math/rand"
 
)

func main() {
 
	const TotalPruebas = 10000
	exitos := 0

	for i := 0; i < TotalPruebas; i++ {
		// Crear baraja de 21 números
		deck := make([]int, 21)
		for j := 0; j < 21; j++ {
			deck[j] = j
		}
		rand.Shuffle(21, func(a, b int) { deck[a], deck[b] = deck[b], deck[a] })

		// Elegir una carta objetivo al azar de la baraja
		target := deck[rand.Intn(21)]

		// Simular 3 rondas
		for round := 0; round < 3; round++ {
			cols := [3][]int{}
			targetCol := -1
			for idx, card := range deck {
				c := idx % 3
				cols[c] = append(cols[c], card)
				if card == target {
					targetCol = c
				}
			}

			// Reordenar: Columna del objetivo al centro
			newDeck := []int{}
			others := []int{}
			for c := 0; c < 3; c++ {
				if c != targetCol {
					others = append(others, c)
				}
			}
			newDeck = append(newDeck, cols[others[0]]...)
			newDeck = append(newDeck, cols[targetCol]...)
			newDeck = append(newDeck, cols[others[1]]...)
			deck = newDeck
		}

		// La carta número 11 (índice 10) debe ser el target
		if deck[10] == target {
			exitos++
		}
	}

	fmt.Printf("--- Verificación del Algoritmo ---\n")
	fmt.Printf("Pruebas ejecutadas: %d\n", TotalPruebas)
	fmt.Printf("Éxitos: %d\n", exitos)
	fmt.Printf("Precisión: %.2f%%\n", float64(exitos)/float64(TotalPruebas)*100)
}
