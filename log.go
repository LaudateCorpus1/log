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
var request = log.New(os.Stderr, "R: ", log.Ldate|log.Ltime)

var callDepth int = 2

func Errorf(format string, v ...interface{}) {
	error.Output(callDepth,f(format, v...))
}

func Errorln(v ...interface{}) {
	error.Output(callDepth,ln(v...))
}

func Warningf(format string, v ...interface{}) {
	warning.Output(callDepth,f(format, v...))
}

func Warningln(v ...interface{}) {
	warning.Output(callDepth,ln(v...))
}

func Infof(format string, v ...interface{}) {
	info.Output(callDepth,f(format, v...))
}

func Infoln(v ...interface{}) {
	info.Output(callDepth,ln(v...))
}

func Requestln(v ...interface{}) {
	request.Println(v...)
}

func Tracef(format string, v ...interface{}) {
	trace.Output(callDepth,f(format, v...))
}

func Traceln(v ...interface{}) {
	trace.Output(callDepth,ln(v...))
}

func Panicln(v ...interface{}) {
	string := ln(v...)
	trace.Output(callDepth, string)
	panic(string)
}

func Panicf(format string, v ...interface{}) {
	string := f(format,v...)
	trace.Output(callDepth, string)
	panic(string)
}

func Slackf(format string, v ...interface{}) {
	slack.Output(callDepth,f(format, v...))
}

func SlackLn(v ...interface{}) {
	slack.Output(callDepth,ln(v...))
}

func f(format string, v ...interface{}) string {
	return ln(fmt.Sprintf(format, v...))
}

func ln(a ...interface{}) string {
	return fmt.Sprintln(a...)
}
