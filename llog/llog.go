package llog

import (
	"log"
	"shutme/cmds"
)

const (
	lDebug = iota
	lInfo
	lWarn
	lError
	lPanic
	lFetal
)

var LogFile = cmds.BaseName + ".log"

func levelLog(level int, logString string) {
	l := []string{"DEBUG", "INFO", "WARN", "ERROR", "PANIC", "FETAL"}
	log.Printf("[%v] %v\n", l[level], logString)
}

func Debug(String string) {
	levelLog(lDebug, String)
}

func Info(String string) {
	levelLog(lInfo, String)
}

func Warn(String string) {
	levelLog(lWarn, String)
}

func Error(String string) {
	levelLog(lError, String)
}

func Panic(String string) {
	levelLog(lPanic, String)
}

func Fetal(String string) {
	levelLog(lFetal, String)
}
