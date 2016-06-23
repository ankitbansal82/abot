// Package driver defines interfaces to be implemented by messaging platform drivers as used
// by package messenger.
package driver

import "github.com/julienschmidt/httprouter"

// Driver is the interface that must be implemented by an messaging platform driver.
type Driver interface {
	// Open returns a new connection to the SMS server. Authentication is
	// handled by the individual drivers.
	Open(r *httprouter.Router) (Conn, error)
}

// Conn is a connection to the external SMS service.
type Conn interface {
	// Send message to an existing chat an SMS from one number to another. It's up to each individual
	// SMS driver to specify the format of the numbers. The From number is
	// handled by the drivers themselves.
	Send(chatID int, msg string) error

	// Close the connection.
	Close() error
}
