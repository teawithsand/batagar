package events

import "io"

// Listener is single listener, which has single channel used for receiving events.
type Listener interface {
	io.Closer
	Chan() <-chan interface{}
}

// Listeners represents listener factory.
type Listeners interface {
	NewListener() Listener
}

// Emiter emits events to all listeners.
type Emiter interface {
	Emit(e interface{})
}
