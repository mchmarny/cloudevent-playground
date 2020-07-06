# cloudevent-playground

Simple test on the unmarshalling gains when you only need the metadata of [CloudEvent](https://github.com/cloudevents/spec/blob/v1.0/spec.md)


```shell
make run 

48.83% faster with      100 bytes data (w:   52652, wo:   25712)
82.76% faster with      100 bytes data (w:    8699, wo:    7199)
84.99% faster with     1000 bytes data (w:   21390, wo:   18179)
83.69% faster with    10000 bytes data (w:  156117, wo:  130660)
80.88% faster with   100000 bytes data (w: 1300480, wo: 1051816)
73.66% faster with  1000000 bytes data (w: 7753188, wo: 5710777)
73.85% faster with 10000000 bytes data (w:71495854, wo:52796504)
```

```go
// CloudEvent is the envelope 
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
```
