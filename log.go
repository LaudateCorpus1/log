package log

import (
	"log"
	"io/ioutil"
	"os"
	"fmt"
)

var trace = log.New(ioutil.Discard, "T: ", log.Ldate|log.Ltime|log.Lshortfile)
var info = log.New(os.Stderr, "I: ", log.Ldate|log.Ltime|log.Lshortfile)
var warning = log.New(os.Stderr, "W: ", log.Ldate|log.Ltime|log.Lshortfile)
var error = log.New(os.Stderr, "E: ", log.Ldate|log.Ltime|log.Lshortfile)
var slack = log.New(os.Stderr, "SLACK: ", log.Ldate|log.Ltime|log.Lshortfile)
var request = log.New(os.Stderr, "R: ", log.Ldate|log.Ltime|log.Lshortfile)

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

func Slackf(format string, v ...interface{}) {
	slack.Print(f(format, v...))
}

func SlackLn(v ...interface{}) {
	slack.Print(ln(v...))
}

func f(format string, v ...interface{}) string {
	return ln(fmt.Sprintf(format, v...))
}

func ln(a ...interface{}) string {
	return fmt.Sprintln(a...)
}
