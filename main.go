package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

var (
	logger = log.New(os.Stdout, "", 0)
)

func main() {

	runTest(100) // warmup
	runTest(100)
	runTest(1000)
	runTest(10000)
	runTest(100000)
	runTest(1000000)
	runTest(10000000)
}

func runTest(size int) {
	s := makeCloudEvent(size)
	c, _ := json.Marshal(s)
	d1 := timeUnmarshal(c, &CloudEvent{})
	d2 := timeUnmarshal(c, &CloudEventMeta{})
	logger.Printf("%.2f%% faster with %8d bytes data (w:%8d, wo:%8d)",
		float64(d2)/float64(d1)*100, size, d1, d2)
}

func timeUnmarshal(c []byte, t interface{}) int64 {
	s := time.Now()
	if err := json.Unmarshal(c, t); err != nil {
		logger.Fatalf("error deserializing content: %v to %T type", err, t)
	}
	return time.Since(s).Nanoseconds()
}
