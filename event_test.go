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

func BenchmarkUnmarshalSmallContentWithData(b *testing.B) {
	bench(smallContent, &CloudEvent{}, b)
}

func BenchmarkUnmarshalSmallContentWithoutData(b *testing.B) {
	bench(smallContent, &CloudEventMeta{}, b)
}

func BenchmarkUnmarshalMediumContentWithData(b *testing.B) {
	bench(mediumContent, &CloudEvent{}, b)
}

func BenchmarkUnmarshalMediumContentWithoutData(b *testing.B) {
	bench(mediumContent, &CloudEventMeta{}, b)
}

func BenchmarkUnmarshalLargeContentWithData(b *testing.B) {
	bench(largeContent, &CloudEvent{}, b)
}

func BenchmarkUnmarshalLargeContentWithoutData(b *testing.B) {
	bench(largeContent, &CloudEventMeta{}, b)
}
