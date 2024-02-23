package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	givenCons := cons(1, cons(2, cons(4, cons(5, nil))))
	
	var expectedSum int64 = 12
	assert.Equal(t, sum(givenCons).Int(), expectedSum )
}


func TestProduct(t *testing.T) {
	givenCons := cons(1, cons(2, cons(4, cons(5, nil))))
	
	var expectedProduct int64 = 40
	assert.Equal(t, product(givenCons).Int() , expectedProduct)
}


func TestAnytrue(t *testing.T) {
	givenCons := cons(false, cons(false, cons(false, cons(true, nil))))

	var expectedAnyTrue bool = true
	assert.Equal(t, anytrue(givenCons).Bool() , expectedAnyTrue)
}


func TestAlltrue(t *testing.T) {
	givenCons := cons(false, cons(false, cons(false, cons(true, nil))))

	var expectedAllTrue bool = false
	assert.Equal(t, alltrue(givenCons).Bool() , expectedAllTrue)
}
