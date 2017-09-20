package transport

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDeepCopyRequest(t *testing.T) {
	src := newTestRequest()
	dest := src.DeepCopy()
	assertRequestIsTestRequest(t, src)
	assertRequestIsTestRequest(t, dest)

	dest.Method = "POST"
	dest.Timeout = 0 * time.Second
	dest.Headers["beep"] = ""
	dest.Baggage["beep"] = ""
	dest.TransportHeaders["beep"] = ""
	dest.ShardKey = ""
	dest.Body = []byte("")
	dest.TargetService = ""

	assertRequestIsTestRequest(t, src)
}

func newTestRequest() *Request {
	return &Request{
		Method:           "GET",
		Timeout:          1 * time.Second,
		Headers:          map[string]string{"beep": "boop"},
		Baggage:          map[string]string{"beep": "bleep"},
		TransportHeaders: map[string]string{"beep": "meep"},
		ShardKey:         "key",
		Body:             []byte("hello world"),
		TargetService:    "bob",
	}
}

func assertRequestIsTestRequest(t *testing.T, req *Request) {
	r := newTestRequest()
	assert.Equal(t, *r, *req)
}

func TestDeepCopyMap(t *testing.T) {
	src := map[string]string{"foo": "bar"}
	dest := deepCopyMap(src)
	assert.Equal(t, "bar", src["foo"])
	assert.Equal(t, "bar", dest["foo"])

	dest["foo"] = "baz"
	assert.Equal(t, "bar", src["foo"])
	assert.Equal(t, "baz", dest["foo"])

	src["foo"] = "qux"
	assert.Equal(t, "qux", src["foo"])
	assert.Equal(t, "baz", dest["foo"])

	dest["zim"] = "zam"
	_, srcHasZim := src["zim"]
	assert.False(t, srcHasZim)
	assert.Equal(t, "zam", dest["zim"])
}

func TestDeepCopyBytes(t *testing.T) {
	initValue := []byte("foo")
	src := initValue
	dest := deepCopyBytes(src)
	assert.True(t, bytes.Equal(initValue, src))
	assert.True(t, bytes.Equal(initValue, dest))

	newValue := []byte("boo")
	dest[0] = 'b'
	assert.True(t, bytes.Equal(initValue, src))
	assert.True(t, bytes.Equal(newValue, dest))

	lastValue := []byte("food")
	src = append(src, 'd')
	assert.True(t, bytes.Equal(lastValue, src))
	assert.True(t, bytes.Equal(newValue, dest))
}
