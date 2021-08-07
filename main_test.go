package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert := assert.New(t)
	var a = 42
	assert.Equal(a, 42, "The two words should be the same.")
}
