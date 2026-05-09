# 🎩 The Mentalist Engine

**The Mentalist Engine** es una experiencia web interactiva diseñada bajo los principios de **Minimalismo Extremo** y **Zero Node**. Utiliza la potencia de **Go** en el backend y **HTMX** para ofrecer el clásico truco de magia matemático de las 21 cartas con un consumo de recursos mínimo, optimizado para funcionar con fluidez incluso en hardware limitado.

## 🚀 Filosofía del Proyecto

Este motor ha sido construido bajo restricciones técnicas estrictas para garantizar un rendimiento excepcional:

*   **HTML-First**: Enfoque centrado en el servidor, evitando la sobrecarga de frameworks de JavaScript modernos.
*   **Zero Node**: Eliminación total de dependencias de Node.js, `npm` y procesos de compilación de frontend pesados.
*   **Arquitectura Minimalista**: Diseñado específicamente para ejecutarse de forma ágil en equipos con recursos de hardware limitados, como procesadores antiguos o poca RAM.
*   **Interactividad Eficiente**: Uso de HTMX para gestionar actualizaciones parciales del DOM (swaps), manteniendo la experiencia fluida sin recargas completas de página.

## 🛠️ Stack Tecnológico

*   **Lenguaje**: [Go (Golang)](https://go.dev/) para la lógica del servidor, gestión de estado y renderizado de plantillas.
*   **Frontend**: [HTMX](https://htmx.org/) para comunicación asíncrona y [Bulma CSS](https://bulma.io/) para una interfaz limpia y responsiva.
*   **Testing**: [Playwright-Go](https://github.com/playwright-community/playwright-go) para pruebas de extremo a extremo (E2E) automáticas.

## 🃏 El Truco de las 21 Cartas

El motor implementa el algoritmo matemático de reordenamiento por columnas:
1. El usuario elige una carta mentalmente de un mazo de 21.
2. Tras tres rondas de selección de columnas, el sistema utiliza la técnica del "sándwich" (posicionando la columna elegida en medio de las otras dos).
3. La carta elegida se desplaza matemáticamente a la posición 11, donde es revelada al final del juego.

## 🔧 Instalación y Uso

### Requisitos previos
*   Go 1.20 o superior.

### Pasos para la ejecución
1. **Clonar el repositorio**:
   ```bash
   git clone [https://github.com/santiagourdaneta/Truco-de-magia-matematico-de-las-21-cartas-en-Go-y-HTMX-](https://github.com/santiagourdaneta/Truco-de-magia-matematico-de-las-21-cartas-en-Go-y-HTMX-)
   cd Truco-de-magia-matematico-de-las-21-cartas-en-Go-y-HTMX-

   Ejecutar el servidor:

   go run main.go

   Ejecutar los tests:

   go test -v ./...

   ---
**Desarrollado por:** [Santiago Urdaneta Anton](https://github.com/santiagourdaneta)