package main

import (
	"fmt"
	"io"

	"github.com/gopherjs/gopherjs/js"
)

// Underscores are to prevent intellisense from picking these
// up and adding them inside the code. It also prevents exporting.

var (
	// Logs to console. Press F12 to open in browser
	_consoleWriter = jsWriterToWriter(js.Global.Get("console"), "log")
	// Writes inside document HTML. Warning: Overwrites all HTML!
	_documentWriter = jsWriterToWriter(js.Global.Get("document"), "write")
	// Default logger when using logf.
	_GlobalLogger *logger = &logger{w: _consoleWriter}
)

func logf(format string, a ...interface{}) {
	_GlobalLogger.Logf(format, a...)
}

type logger struct {
	w io.Writer
}

func (l *logger) Logf(format string, a ...interface{}) {
	l.w.Write([]byte(fmt.Sprintf(format, a...)))
}

type jsWriter interface {
	Call(name string, args ...interface{}) *js.Object
}

type jsWriterObj struct {
	writeCall string
	jsWriter
}

func jsWriterToWriter(j jsWriter, call string) io.Writer {
	return jsWriterObj{call, j}
}

func (j jsWriterObj) Write(b []byte) (int, error) {
	j.Call(j.writeCall, string(b))
	return len(b), nil
}
