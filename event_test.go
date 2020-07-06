package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeCloudEvent(t *testing.T) {
	d := 10
	ce := makeCloudEvent(d)
	assert.NotNil(t, ce)
	assert.Lenf(t, ce.Data, d, "invalid generated cloud event data size")
}
