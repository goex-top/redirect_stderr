// +build darwin

// Log the panic under unix to the log file

package stderr

import (
	"os"
)

// redirectStderr to the file passed in
func RedirectStderr(f *os.File) {

}
