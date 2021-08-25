package console

import (
	"fmt"
	"os"
)

// write all arguments to a file f
func logToFile(f *os.File, args ... interface{}) {

	// print each param independently
	for _, arg := range args {
		fmt.Fprint(f, arg)
	}
	// print the endline
	fmt.Fprint(f, "\n")

}

// logs various arguments to a newline terminated string in stdout
func log(args ... interface{}) {
	logToFile(os.Stdout, args...)
}


// logs various arguments to a newline terminated string in stderr
func error(args ... interface{}) {
	logToFile(os.Stderr, args...)
}