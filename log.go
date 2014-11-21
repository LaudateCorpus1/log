package log

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	slack   *log.Logger
	request *log.Logger
}

func New(prefix string) *Logger {

	l := &Logger{
		info:    log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
		warning: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
		error:   log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile),
		slack:   log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
		request: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}

	l.SetPrefix(prefix)
	return l
}

var DefaultLogger = New("")

type BLogger interface {
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})
	Warningf(format string, v ...interface{})
	Warningln(v ...interface{})
	Infof(format string, v ...interface{})
	Infoln(v ...interface{})
	Requestln(v ...interface{})
	Panicln(v ...interface{})
	Panicf(format string, v ...interface{})
	Slackf(format string, v ...interface{})
	SlackLn(v ...interface{})
}

var callDepth int = 2

func (this *Logger) SetPrefix(prefix string) {

	if prefix != "" {
		prefix = prefix + " "
	}

	this.info.SetPrefix(prefix + "I: ")
	this.warning.SetPrefix(prefix + "W: ")
	this.error.SetPrefix(prefix + "E: ")
	this.slack.SetPrefix(prefix + "SLACK: ")
	this.request.SetPrefix(prefix + "R: ")
}

func (this *Logger) Errorf(format string, v ...interface{}) {
	this.error.Output(callDepth, f(format, v...))
}

func (this *Logger) Errorln(v ...interface{}) {
	this.error.Output(callDepth, ln(v...))
}

func (this *Logger) Warningf(format string, v ...interface{}) {
	this.warning.Output(callDepth, f(format, v...))
}

func (this *Logger) Warningln(v ...interface{}) {
	this.warning.Output(callDepth, ln(v...))
}

func (this *Logger) Infof(format string, v ...interface{}) {
	this.info.Output(callDepth, f(format, v...))
}

func (this *Logger) Infoln(v ...interface{}) {
	this.info.Output(callDepth, ln(v...))
}

func (this *Logger) Requestln(v ...interface{}) {
	this.request.Println(v...)
}

func (this *Logger) Panicln(v ...interface{}) {
	string := ln(v...)
	this.error.Output(callDepth, string)
	panic(string)
}

func (this *Logger) Panicf(format string, v ...interface{}) {
	string := f(format, v...)
	this.error.Output(callDepth, string)
	panic(string)
}

func (this *Logger) Slackf(format string, v ...interface{}) {
	this.slack.Output(callDepth, f(format, v...))
}

func (this *Logger) SlackLn(v ...interface{}) {
	this.slack.Output(callDepth, ln(v...))
}

//Globals
func Errorf(format string, v ...interface{}) {
	DefaultLogger.Errorf(format, v...)
}

func Errorln(v ...interface{}) {
	DefaultLogger.Errorln(v...)
}

func Warningf(format string, v ...interface{}) {
	DefaultLogger.Warningf(format, v...)
}

func Warningln(v ...interface{}) {
	DefaultLogger.Warningln(v...)
}

func Infof(format string, v ...interface{}) {
	DefaultLogger.Infof(format, v...)
}

func Infoln(v ...interface{}) {
	DefaultLogger.Infoln(v...)
}

func Requestln(v ...interface{}) {
	DefaultLogger.Requestln(v...)
}

func Panicln(v ...interface{}) {
	DefaultLogger.Panicln(v...)
}

func Panicf(format string, v ...interface{}) {
	DefaultLogger.Panicf(format, v...)
}

func Slackf(format string, v ...interface{}) {
	DefaultLogger.Slackf(format, v...)
}

func SlackLn(v ...interface{}) {
	DefaultLogger.SlackLn(v...)
}

func f(format string, v ...interface{}) string {
	return ln(fmt.Sprintf(format, v...))
}

func ln(a ...interface{}) string {
	return fmt.Sprintln(a...)
}
