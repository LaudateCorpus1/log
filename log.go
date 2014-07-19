package log

import (
	"log"
	"io/ioutil"
	"os"
	"fmt"
	"runtime"
	"path/filepath"
)

var trace = log.New(ioutil.Discard, "T: ", log.Ldate|log.Ltime)
var info = log.New(os.Stderr, "I: ", log.Ldate|log.Ltime)
var warning = log.New(os.Stderr, "W: ", log.Ldate|log.Ltime)
var error = log.New(os.Stderr, "E: ", log.Ldate|log.Ltime)
var request = log.New(os.Stderr, "R: ", log.Ldate|log.Ltime)

func Errorf(format string, v ...interface{}) {
	error.Print(f(format, v...))
}

func Errorln(v ...interface{}) {
	error.Print(ln(v...))
}

func Warningf(format string, v ...interface{}) {
	warning.Print(f(format, v...))
}

func Warningln(v ...interface{}) {
	warning.Print(ln(v...))
}

func Infof(format string, v ...interface{}) {
	info.Print(f(format, v...))
}

func Infoln(v ...interface{}) {
	info.Print(ln(v...))
}

func Requestln(v ...interface{}) {
	request.Println(v...)
}

func Tracef(format string, v ...interface{}) {
	trace.Print(f(format, v...))
}

func Traceln(v ...interface{}) {
	trace.Print(ln(v...))
}

func Panicln(v ...interface{}) {
	error.Panic(ln(v...))
}

func Panicf(format string, v ...interface{}) {
	error.Panic(f(format, v...))
}

func f(format string, v ...interface{}) string {
	return ln(fmt.Sprintf(format, v...))
}

func ln(a ...interface{}) string {

	_, file, line, _ := runtime.Caller(2)

	filename := filepath.Base(file)
	directory := filepath.Dir(file)
	directoryName := filepath.Base(directory)

	//f := runtime.FuncForPC(pc)
	info := fmt.Sprintln(a...)
	return fmt.Sprintf("%s%c%s:%d: %v", directoryName, filepath.Separator, filename, line, info)
}
