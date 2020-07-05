package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

const charset = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`

var (
	logger                = log.New(os.Stdout, "", 0)
	seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func getRandomChars(length int) []byte {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return b
}

type cloudEventEnvelope struct {
	cloudEventMeta
	Data interface{} `json:"data"`
}

type cloudEventMeta struct {
	ID              string `json:"id"`
	Source          string `json:"source"`
	Type            string `json:"type"`
	SpecVersion     string `json:"specversion"`
	DataContentType string `json:"datacontenttype"`
	Subject         string `json:"subject"`
}

func main() {
	runTest(100)
	runTest(1000)
	runTest(10000)
	runTest(100000)
	runTest(1000000)
	runTest(10000000)
}

func runTest(size int) {
	s := makeStruct(size)
	c, _ := json.Marshal(s)
	d1 := unmarshal(c, &cloudEventEnvelope{})
	d2 := unmarshal(c, &cloudEventMeta{})
	logger.Printf("%.2f faster using %d bytes data", float64(d1)/float64(d2), len(c))
}

func unmarshal(c []byte, t interface{}) int64 {
	s := time.Now()
	if err := json.Unmarshal(c, t); err != nil {
		panic(err)
	}
	return time.Since(s).Nanoseconds()
}

func makeStruct(size int) *cloudEventEnvelope {
	ce := &cloudEventEnvelope{
		Data: getRandomChars(size),
	}
	ce.ID = string(getRandomChars(36))    // UUID
	ce.Source = string(getRandomChars(6)) // App Name
	ce.Type = "com.dapr.event.sent"
	ce.SpecVersion = "v1.0"
	ce.DataContentType = "application/cloudevents+json"
	return ce
}
