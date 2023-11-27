//go:generate goversioninfo
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"shutme/cmds"
	"shutme/llog"
	"shutme/probe"
	"shutme/serv"
	"shutme/shutmedown"
	"strings"
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

	if cmds.Flag_c == "" {
		err = shutmedown.ShutMeCmdPerse()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	}

	if len(cmds.Flag_s) != 0 {
		if cmds.Flag_s == "install" {
			if !Confirm() {
				os.Exit(0)
			}
		}
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
		logFile, err := os.OpenFile(llog.LogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
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

	if !Confirm() {
		return
	}

	logFile, err := os.OpenFile(llog.LogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(io.MultiWriter(os.Stderr, logFile))
	defer logFile.Close()

	cs := make(chan os.Signal, 1)
	signal.Notify(cs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go probe.ProbeRemote()

	<-cs
	fmt.Fprintln(os.Stderr, "Program interrupted.")

}

// Print some necessary information and confirm that you want to continue
// Param : none
// Return: boolean
func Confirm() bool {
	var res string

	fmt.Printf("Attempt to detect remote host status... ")
	if _, err := probe.Ping(cmds.Flag_t); err != nil {
		fmt.Fprintf(os.Stderr, "Failed.\nCommunication cannot be established with the remote host %s, program terminates.\n", cmds.Flag_t) //todo
		return false
	} else {
		fmt.Printf("OK.\n")
	}

	if !cmds.Flag_y {

		fmt.Println("WARNING, Once the network failure occurs after the program is running, it will trigger the shutdown behavior.")
		fmt.Printf("Are you sure you want to do this? \nPress 'YES' to continue:")
		fmt.Scanln(&res)
		if strings.ToUpper(res) != "YES" {
			return false
		}
	}

	return true
}
