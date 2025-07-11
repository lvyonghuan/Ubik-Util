package umessenger

func (messenger *UMessenger) PostLog(log string, level int) {
	go func() { // Convert log to bytes, assuming a function exists for this
		var logBytes = []byte(log)

		envelope := messenger.NewEnvelope("PUT", level, "", logBytes)
		messenger.PostMessage(envelope)
	}() // Use a goroutine to avoid blocking the main thread
}
