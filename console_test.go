package console;

import (
	"testing"
	"io"
	"os"
	"bytes"
);

func TestConsoleLog(t *testing.T) {

	// test 1
	output := getStdOut(func() {
		log("Hello world")
	});
	if output != "Hello world\n" {
		t.Error("Test 1 of console output for log failed")
	}

	// test 2
	output = getStdOut(func() {
		log("world-", 2, "-", 2.45)
	});
	if output != "world-2-2.45\n" {
		t.Error("Test 2 of console output for log failed")
	}
}

func TestConsoleError(t *testing.T) {

	// test 1
	output := getStdErr(func() {
		error("Hello world")
	});
	if output != "Hello world\n" {
		t.Error("Test 1 of console output for error failed")
	}

	// test 2
	output = getStdErr(func() {
		error("world-", 2, "-", 2.45)
	});
	if output != "world-2-2.45\n" {
		t.Error("Test 2 of console output for error failed")
	}
}

func getStdOut(targetFunction func()) string {

	// replace stdout with a pipe
	old_stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// execute the function
	targetFunction()

	outC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, r)
        outC <- buf.String()
    }()

	// close the pipe and restore stdout
	w.Close()
	os.Stdout = old_stdout

	// return
	return <-outC
}

func getStdErr(targetFunction func()) string {

	// replace stdout with a pipe
	old_stderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	// execute the function
	targetFunction()

	outC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, r)
        outC <- buf.String()
    }()

	// close the pipe and restore stderr
	w.Close()
	os.Stderr = old_stderr

	// return
	return <-outC
}