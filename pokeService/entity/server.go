package entity

import "net/http"

// ServerHandler is the function that handles a server request
type ServerHandler func(w http.ResponseWriter, r *http.Request)

// ServerProperties defines properties about the server
type ServerProperties struct {
	Address string
}
