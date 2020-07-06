# cloudevent-playground

Simple test on the unmarshalling gains when you only need the metadata of [CloudEvent](https://github.com/cloudevents/spec/blob/v1.0/spec.md)


```shell
make bench 

BenchmarkUnmarshalSmallContentWithData-8       	1000000000	         0.000019 ns/op
BenchmarkUnmarshalSmallContentWithoutData-8    	1000000000	         0.000015 ns/op
BenchmarkUnmarshalMediumContentWithData-8      	1000000000	         0.000793 ns/op
BenchmarkUnmarshalMediumContentWithoutData-8   	1000000000	         0.000711 ns/op
BenchmarkUnmarshalLargeContentWithData-8       	1000000000	         0.0674 ns/op
BenchmarkUnmarshalLargeContentWithoutData-8    	1000000000	         0.0564 ns/op
PASS
ok  	github.com/mchmarny/cloudevent-playground	2.261s
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
