package main

import (
	"encoding/json"
	"math/rand"
	"sync"
	"time"
)

const charset = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`

// sharedRand is reused across calls so a burst of events does not reseed per call.
var sharedRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// CloudEvent is the CloudEvent envelope
type CloudEvent struct {
	CloudEventMeta
	Data []byte `json:"data"`
}

// CloudEventMeta is the Cloud Event metadata
type CloudEventMeta struct {
	ID              string `json:"id"`
	Source          string `json:"source"`
	Type            string `json:"type"`
	SpecVersion     string `json:"specversion"`
	DataContentType string `json:"datacontenttype"`
	Subject         string `json:"subject"`
}

func getCloudEventContent(dataSize int) []byte {
	ce := &CloudEvent{
		Data: getRandomBytes(dataSize),
	}
	ce.ID = string(getRandomBytes(36))    // UUID
	ce.Source = string(getRandomBytes(6)) // App Name
	ce.Type = "com.dapr.event.sent"
	ce.SpecVersion = "v1.0"
	ce.DataContentType = "application/cloudevents+json"
	c, _ := json.Marshal(ce)
	return c
}

// getCloudEvents builds n events concurrently.
func getCloudEvents(n, dataSize int) [][]byte {
	out := make([][]byte, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out[i] = getCloudEventContent(dataSize)
		}()
	}
	wg.Wait()
	return out
}

// parseCloudEvent decodes an event produced by getCloudEventContent.
func parseCloudEvent(b []byte) *CloudEvent {
	var ce CloudEvent
	json.Unmarshal(b, &ce)
	return &ce
}

func getRandomBytes(length int) []byte {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[sharedRand.Intn(len(charset))]
	}
	return b
}
