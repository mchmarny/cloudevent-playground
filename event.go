package main

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

func makeCloudEvent(dataSize int) *CloudEvent {
	ce := &CloudEvent{
		Data: getRandomBytes(dataSize),
	}
	ce.ID = string(getRandomBytes(36))    // UUID
	ce.Source = string(getRandomBytes(6)) // App Name
	ce.Type = "com.dapr.event.sent"
	ce.SpecVersion = "v1.0"
	ce.DataContentType = "application/cloudevents+json"
	return ce
}
