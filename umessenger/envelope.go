package umessenger

import "github.com/lvyonghuan/Ubik-Util/uconst"

// Message categories
const (
	result = 0 //To submit nodes' results
	Fatal  = uconst.Fatal
	Error  = uconst.Error
	Warn   = uconst.Warn
	Info   = uconst.Info
	Debug  = uconst.Debug
)

type UEnvelope struct {
	uuid    string // UUID of the follower that sent the message
	methods string // Methods used to send the message, such as "POST", "GET", etc.

	Category int    `json:"Category"` // Message category. From 0 to 5, representing result, fatal, error, warn, info, and debug respectively.
	Flag     string `json:"flag"`     // A custom flag field for the flag. Can be null.
	Message  []byte `json:"message"`  // Message content, a formatted byte array.
}

func (messenger *UMessenger) NewEnvelope(methods string, category int, flag string, message []byte) UEnvelope {
	return UEnvelope{
		uuid:     messenger.uuid,
		methods:  methods,
		Category: category,
		Flag:     flag,
		Message:  message,
	}
}
