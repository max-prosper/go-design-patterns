package barrier

import (
	"strings"
	"testing"
)

func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}

		result := captureBarrierOutput(endpoints...)

		if !strings.Contains(result, "Accept-Encoding") || !strings.Contains(result, "User-Agent") {
			t.Fail()
		}
		t.Log(result)
	})

	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"http://malformed-url", "http://httpbin.org/User-Agent"}

		result := captureBarrierOutput(endpoints...)

		if !strings.Contains(result, "ERROR") {
			t.Fail()
		}
		t.Log(result)
	})

	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}

		timeoutMilliseconds = 1

		result := captureBarrierOutput(endpoints...)

		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}
		t.Log(result)
	})
}
