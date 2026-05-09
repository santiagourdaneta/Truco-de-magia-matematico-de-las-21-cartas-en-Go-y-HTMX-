package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
    "os"
 
)

type Card struct{ Value string }
type GameState struct {
	Deck  []Card
	Round int
}

var state GameState

var uiTemplate = `
{{ if gt .Round 3 }}
    <div class="result">
        <p>✨ Tu carta es la: <strong>{{ (index .Deck 10).Value }}</strong> ✨</p>
        <button hx-post="/start" hx-target="#game-ui">Jugar de Nuevo</button>
    </div>
{{ else }}
    <p>Ronda {{ .Round }} de 3</p>
    <div class="column-container">
        {{ range $i, $col := .Cols }}
        <div class="column" hx-post="/pick?col={{ $i }}" hx-target="#game-ui">
            <strong>Columna {{ add $i 1 }}</strong>
            {{ range $col }}
            <div class="card">{{ .Value }}</div>
            {{ end }}
        </div>
        {{ end }}
    </div>
{{ end }}
`

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Truco 21 Cartas - Go + HTMX</title>
    <script src="https://unpkg.com/htmx.org@1.9.10"></script>
    <style>

         body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background: #1a1a1a; color: #eee; text-align: center; margin: 0; padding: 20px; }
.column-container { 
    display: flex; 
    flex-wrap: wrap; 
    justify-content: center; 
    gap: 15px; 
    margin-top: 20px; 
}
.column { 
    background: #2a2a2a; 
    border: 2px solid #444; 
    border-radius: 12px; 
    padding: 15px; 
    width: 100%; 
    max-width: 140px; 
    cursor: pointer; 
    transition: transform 0.2s, border-color 0.2s; 
}
.column:hover { border-color: #007bff; transform: translateY(-5px); }
.card { 
    padding: 8px; 
    background: #333; 
    margin: 4px 0; 
    border-radius: 4px; 
    font-size: 1.1em;
}
@media (max-width: 600px) {
    .column { max-width: 90px; padding: 8px; font-size: 0.9em; }
}

    </style>
</head>
<body>
    <h1>🎩 El Mentalista Matemático</h1>
    <p>Memoriza una carta y selecciona la columna donde aparece.</p>
    <div id="game-ui">
        <button hx-post="/start" hx-target="#game-ui">Empezar Juego</button>
    </div>
</body>
</html>
`)
	})
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/start", handleStart)
	http.HandleFunc("/pick", handlePick)

    port := os.Getenv("PORT")
    if port == "" {
            port = "8080"
    }

    address := ":" + port
    fmt.Println("Servidor iniciado en", address)

	if err := http.ListenAndServe(address, nil); err != nil {
    fmt.Printf("Error al iniciar servidor: %v\n", err)
}
}

func handleStart(w http.ResponseWriter, r *http.Request) {
	deck := []Card{}
	suits := []string{"♠", "♥", "♦", "♣"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for i := 0; i < 21; i++ {
		deck = append(deck, Card{Value: values[i%13] + suits[i%4]})
	}
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	state = GameState{Deck: deck, Round: 1}
	renderUI(w)
}

func reorderDeck(colIdx int) {
    cols := [3][]Card{}
    for i, card := range state.Deck {
        cols[i%3] = append(cols[i%3], card)
    }

    var order []int
    if colIdx == 1 {
        order = []int{0, 1, 2}
    } else if colIdx == 0 { //
        order = []int{1, 0, 2}
    } else {
        order = []int{0, 2, 1}
    }

    newDeck := []Card{}
    for _, idx := range order {
        newDeck = append(newDeck, cols[idx]...)
    }

    state.Deck = newDeck
    state.Round++
}

func handlePick(w http.ResponseWriter, r *http.Request) {
    colIdx := 0
    if _, err := fmt.Sscanf(r.URL.Query().Get("col"), "%d", &colIdx); err != nil {
    colIdx = 0 // Valor por defecto si falla
}
    reorderDeck(colIdx)
    renderUI(w)
}

func renderUI(w http.ResponseWriter) {
	cols := [3][]Card{}
	for i, card := range state.Deck {
		cols[i%3] = append(cols[i%3], card)
	}

	funcMap := template.FuncMap{"add": func(a, b int) int { return a + b }}
	t := template.Must(template.New("ui").Funcs(funcMap).Parse(uiTemplate))

	// Estructura de datos para el template
	data := struct {
		Deck  []Card
		Round int
		Cols  [3][]Card
	}{
		Deck:  state.Deck,
		Round: state.Round,
		Cols:  cols,  
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}