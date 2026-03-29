package calculator

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    assert.Equal(t, 5, result)
}

func TestSubtract(t *testing.T) {
    result := Subtract(10, 4)
    assert.Equal(t, 6, result)
}

func TestMultiply(t *testing.T) {
    result := Multiply(3, 4)
    assert.Equal(t, 12, result)
}

func TestDivide(t *testing.T) {
    result := Divide(20, 4)
    assert.Equal(t, 5, result)
}