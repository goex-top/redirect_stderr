package stderr

import (
	"os"
	"testing"
)

func Test_redirectStderr(t *testing.T) {
	logFilename := "foo.log"
	// redirect stdout and stderr to log file
	logFile, _ := os.OpenFile(logFilename, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0644)
	defer logFile.Close()
	os.Stderr.Write([]byte("before\n"))
	RedirectStderr(logFile)
	os.Stderr.Write([]byte("after\n"))
	// Output to console:
	// before
	// after
	// Output to foo.log: nothing - it is an empty file

}
