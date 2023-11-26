//go:generate goversioninfo
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"shutme/cmds"
	"shutme/serv"
	"syscall"

	"github.com/kardianos/service"
)

func main() {
	fmt.Printf("ShutMe Helper v0.8.1.0 Copyright(C) 2023  limpo@live.com\n\n")

	err := cmds.CmdLine()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		//flag.PrintDefaults()
		//log.Println(err)
		os.Exit(1)
	}

	if len(cmds.Flag_s) != 0 {
		servs, err := serv.ServInit()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		err = serv.ServCtrl(servs)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	if !service.Interactive() {
		logFile, err := os.OpenFile(cmds.MyLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}
		log.SetOutput(logFile)
		defer logFile.Close()

		servs, err := serv.ServInit()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		err = servs.Run()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		return
	}

	if !cmds.Confirm() {
		return
	}

	logFile, err := os.OpenFile(cmds.MyLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(io.MultiWriter(os.Stderr, logFile))
	defer logFile.Close()

	cs := make(chan os.Signal, 1)
	signal.Notify(cs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go cmds.ProbeRemote()

	<-cs
	fmt.Fprintln(os.Stderr, "Program interrupted.")

}
