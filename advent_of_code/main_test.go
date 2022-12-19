package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	elf, calories := findMostCalores("test.txt")
	assert.Equal(t, elf, 4)
	assert.Equal(t, 24000, calories)
}
