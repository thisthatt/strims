// +build js

package wasmio

import (
	"syscall/js"
)

const webSocketMTU = 64 * 1024

// WebSocketProxy ...
type WebSocketProxy interface {
	MTU() int
	Write(b []byte) (int, error)
	Read(b []byte) (n int, err error)
	Close() error
}

// NewWebSocketProxy ...
func NewWebSocketProxy(bridge js.Value, uri string) (WebSocketProxy, error) {
	return newChannel(dataChannelMTU, bridge, "openWebSocket", uri)
}
