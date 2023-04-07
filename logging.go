// logging.go
package main

import (
	"math/rand"
	"time"
)

func random007Phrase() string {
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
