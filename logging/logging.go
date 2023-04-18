// logging.go
package logging

// Http logging for important messages BITCH!!!
func HttpLogger() []string {
	logTypes := []string{
		"INFO",
		"WARNING",
		"ERROR",
	}
	return logTypes
}
