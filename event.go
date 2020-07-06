package main

import (
	"encoding/json"
	"math/rand"
	"time"
)

const charset = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`

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

func getRandomBytes(length int) []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return b
}
