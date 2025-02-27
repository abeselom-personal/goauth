// Package store provides an interface to interact with a database.
//
// The package is designed to be easy to use and provides a simple
// interface for connecting and disconnecting from a database.
//
// The package also provides a simple way to create a database from a
// configuration file.
//
// The package is designed to be extensible and makes it easy to add
// support for different databases.
package store

// Database is an interface that defines methods for connecting and disconnecting from a database.
type Database interface {
	// Connect establishes a connection to the database and returns an error if the connection could not be established.
	Connect() error

	// Disconnect closes the connection to the database and returns an error if the disconnection process fails.
	Disconnect() error
}
