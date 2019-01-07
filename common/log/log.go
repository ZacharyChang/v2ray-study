//go:generate protoc -I $GOPATH/src --go_out=plugins=grpc:$GOPATH/src v2ray.com/v2ray-study/common/log/log.proto
package log // import "v2ray.com/v2ray-study/common/log"

import (
	"sync"

	"v2ray.com/v2ray-study/common/serial"
)

// Message is the interface for all log messages.
type Message interface {
	String() string
}

// Handler is the interface for log handler.
type Handler interface {
	Handle(msg Message)
}

// GeneralMessage is a general log message that can contain all kind of content.
type GeneralMessage struct {
	Severity Severity
	Content  interface{}
}

// String implements Message.
func (m *GeneralMessage) String() string {
	return serial.Concat("[", m.Severity, "] ", m.Content)
}

// Record writes a message into log stream.
func Record(msg Message) {
	logHandler.Handle(msg)
}

var (
	logHandler syncHandler
)

// RegisterHandler register a new handler as current log handler. Previous registered handler will be discarded.
func RegisterHandler(handler Handler) {
	if handler == nil {
		panic("Log handler is nil")
	}
	logHandler.Set(handler)
}

type syncHandler struct {
	sync.RWMutex
	// Note: Handler is the interface for log
	Handler
}

// Note: Handle implements Handler
func (h *syncHandler) Handle(msg Message) {
	// Note: read lock provided by sync.RWMutex
	h.RLock()
	defer h.RUnlock()

	if h.Handler != nil {
		h.Handler.Handle(msg)
	}
}

// Note: set handler with write lock
func (h *syncHandler) Set(handler Handler) {
	h.Lock()
	defer h.Unlock()

	h.Handler = handler
}
