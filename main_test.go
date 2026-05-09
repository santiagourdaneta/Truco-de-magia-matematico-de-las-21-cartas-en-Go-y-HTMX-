package main

import (
    "fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Unit: Verifica la lógica del sandwich matemático
func TestSandwichLogic(t *testing.T) {
	deck := make([]Card, 21)
	for i := 0; i < 21; i++ {
		deck[i] = Card{Value: fmt.Sprintf("%d", i)}
	}
	state = GameState{Deck: deck, Round: 1}

	// Simular elección de columna 0
	reorderDeck(0)

	if len(state.Deck) != 21 {
		t.Errorf("Se perdieron cartas en el reordenamiento, total: %d", len(state.Deck))
	}
}

// Integración: Verifica que /start inicialice el juego
func TestStartEndpoint(t *testing.T) {
	req, _ := http.NewRequest("POST", "/start", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleStart)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code incorrecto: obtenido %v esperado %v", status, http.StatusOK)
	}
}
