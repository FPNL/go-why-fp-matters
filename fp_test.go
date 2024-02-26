package fp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	givenCons := cons(1, cons(2, cons(4, cons(5, nil))))

	var expectedSum int64 = 12
	assert.Equal(t, sum(givenCons).Int(), expectedSum)
}

func TestProduct(t *testing.T) {
	givenCons := cons(1, cons(2, cons(4, cons(5, nil))))

	var expectedProduct int64 = 40
	assert.Equal(t, product(givenCons).Int(), expectedProduct)
}

func TestAnytrue(t *testing.T) {
	givenCons := cons(false, cons(false, cons(false, cons(true, nil))))

	assert.True(t, anytrue(givenCons).Bool())
}

func TestAlltrue(t *testing.T) {
	givenCons := cons(false, cons(false, cons(false, cons(true, nil))))

	assert.False(t, alltrue(givenCons).Bool())
}
