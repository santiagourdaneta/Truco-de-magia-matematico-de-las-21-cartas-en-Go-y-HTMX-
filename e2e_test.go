package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/assert"
)

func TestEndToEndMagicTrick(t *testing.T) {
	// 1. Iniciar el servidor en una goroutine
	go main()

	// Esperar a que el servidor levante el puerto 8080
	time.Sleep(2 * time.Second)

	pw, err := playwright.Run()
	assert.NoError(t, err)
	defer pw.Stop()

	// Usamos Headless: true para velocidad 
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true), 
	})
	assert.NoError(t, err)
	defer browser.Close()

	page, err := browser.NewPage()
	assert.NoError(t, err)

	// 2. Navegar a la raíz
	_, err = page.Goto("http://localhost:8080")
	assert.NoError(t, err)

	// 3. Click en "Empezar Juego" (esto dispara HTMX hx-post="/start")
	btnStart := page.Locator("button:has-text('Empezar Juego')")
	err = btnStart.Click()
	assert.NoError(t, err)

	// 4. Realizar las 3 rondas
	for i := 1; i <= 3; i++ {
		// Esperar a que aparezca el texto de la ronda
		rondaText := fmt.Sprintf("Ronda %d de 3", i)
		err = page.GetByText(rondaText).WaitFor()
		assert.NoError(t, err)

		// Seleccionar la primera columna (esto dispara hx-post="/pick")
		col := page.Locator(".column").First()
		err = col.Click()
		assert.NoError(t, err)
	}

	// 5. Verificar el resultado final
	// El template muestra "✨ Tu carta es la:"
	resultado := page.Locator(".result p")
	err = resultado.WaitFor()
	assert.NoError(t, err)

	textoFinal, _ := resultado.TextContent()
	fmt.Println("Resultado del truco:", textoFinal)
	assert.Contains(t, textoFinal, "Tu carta es la")
}