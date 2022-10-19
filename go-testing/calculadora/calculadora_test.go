package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := -2

	// Se ejecuta el test
	resultado := Restar(num1, num2)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}

func TestSort(t *testing.T) {
	unsorted := []int{2, 1, 3, 5, 4, 6, 7, 9}
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 9}

	QuickSort(unsorted, 0, len(unsorted)-1)
	assert.Equal(t, unsorted, sorted, "no ordena bien")
}

func TestDivide(t *testing.T) {
	_, err := Divide(1, 0)
	assert.ErrorContains(t, err, "El denominador no puede ser 0")
	assert.EqualErrorf(t, err, "El denominador no puede ser 0", "AAAAAAAAAAAA")
}
