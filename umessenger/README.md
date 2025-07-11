# UMessenger

A packet for uniform sending rules. It just post the message but not wait for the response.

Use envelope to define the rules of sending messages.

---

## Envelope

An envelope contains information formatted in []byte, as well as several fields for determining the destination of the message.

```go
type UEnvelope struct {
	UUID      string `json:"uuid"` // UUID of the follower that sent the message
	methods string // Methods used to send the message, such as "POST", "GET", etc.
	sendToURI string // To send the message to a specific URI.
	Category int    `json:"Category"` // Message category. From 0 to 5, representing result, fatal, error, warn, info, and debug respectively.
	Flag     string `json:"flag"`     // A custom flag field for the flag. Can be null.
	Message  []byte `json:"message"`  // Message content, which can be a string or a byte array.
}
```

The `Message` will upload in the body of the request.