package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomBytes(t *testing.T) {
	d := 10
	c := getRandomBytes(d)
	assert.NotNil(t, c)
	assert.Lenf(t, c, d, "invalid generated content size")
}
