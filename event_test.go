package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	smallContent  = 1000
	mediumContent = 100000
	largeContent  = 10000000
)

func TestGetRandomBytes(t *testing.T) {
	d := 10
	c := getRandomBytes(d)
	assert.NotNil(t, c)
	assert.Lenf(t, c, d, "invalid generated content size")
}

func bench(i int, t interface{}, b *testing.B) {
	c := getCloudEventContent(i)
	b.ResetTimer()
	_ = json.Unmarshal(c, t)
}

func BenchmarkSmall(b *testing.B) {
	var t interface{}
	bench(smallContent, &t, b)
}

func BenchmarkSmallWithData(b *testing.B) {
	bench(smallContent, &CloudEvent{}, b)
}

func BenchmarkSmallWithoutData(b *testing.B) {
	bench(smallContent, &CloudEventMeta{}, b)
}

func BenchmarkMedium(b *testing.B) {
	var t interface{}
	bench(mediumContent, &t, b)
}

func BenchmarkMediumWithData(b *testing.B) {
	bench(mediumContent, &CloudEvent{}, b)
}

func BenchmarkMediumWithoutData(b *testing.B) {
	bench(mediumContent, &CloudEventMeta{}, b)
}

func BenchmarkLarge(b *testing.B) {
	var t interface{}
	bench(largeContent, &t, b)
}

func BenchmarkLargeWithData(b *testing.B) {
	bench(largeContent, &CloudEvent{}, b)
}

func BenchmarkLargeWithoutData(b *testing.B) {
	bench(largeContent, &CloudEventMeta{}, b)
}
