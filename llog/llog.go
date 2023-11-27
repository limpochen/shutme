package llog

import (
	"log"
	"shutme/cmds"
)

const (
	Debug = iota
	Info
	Warn
	Error
	Panic
	Fetal
)

var LogFile = cmds.BaseName + ".log"

func levelLog(level int, logString string) {
	l := []string{"DEBUG", "INFO", "WARN", "ERROR", "PANIC", "FETAL"}
	log.Println("["+l[level]+"]", logString)
}

func DebugLog(String string) {
	levelLog(Debug, String)
}

func InfoLog(String string) {
	levelLog(Info, String)
}

func WarnLog(String string) {
	levelLog(Warn, String)
}

func ErrorLog(String string) {
	levelLog(Error, String)
}
