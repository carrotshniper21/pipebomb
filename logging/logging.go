// logging.go
package logging

import (
	"math/rand"
	"time"
)

// Http logging for important messages BITCH!!!
func HttpLogger() []string {
	logTypes := []string{
		"INFO",
		"WARNING",
		"ERROR",
	}
	return logTypes
}

// Random007Phrase returns a random 007 phrase
func Random007Phrase() string {
	phrases := []string{
		"Shaken, not stirred.",
		"The name is Bond, James Bond.",
		"License to kill.",
		"007 reporting for duty.",
		"Keeping the British end up, Sir.",
	}

	rand.Seed(time.Now().UnixNano())
	return phrases[rand.Intn(len(phrases))]
}
