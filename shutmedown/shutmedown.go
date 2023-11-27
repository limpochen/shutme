package shutmedown

import (
	"fmt"
	"os/exec"
	"runtime"
	"shutme/cmds"
	"strings"
)

var ShutMeCmd string

// Command line to perform shutdown for current system
// Param : none
// Retern: string, If null, the current system is not supported
func ShutMeCmdPerse() error {

	switch runtime.GOOS {
	case "linux":
		if !cmds.Flag_b {
			ShutMeCmd = "shutdown -t 0 -h"
			if cmds.Flag_f {
				ShutMeCmd += " -n"
			}
		} else {
			ShutMeCmd = "echo disk > /sys/power/state"
			// ShutMeCmd ="systemctl hibernate"
			// Check to see if suspend to hard disk is supported: "cat /sys/power/state"
		}

	case "windows":
		if !cmds.Flag_b {
			ShutMeCmd = "shutdown -s -t 0"
		} else {
			ShutMeCmd = "shutdown -h -t 0"
		}
		if cmds.Flag_f {
			ShutMeCmd += " -f"
		}

	//case "darwin":
	//	ShutMeCmd = "sudo shutdown -h now"

	default:
		ShutMeCmd = ""
		return fmt.Errorf("this os is not suppered")
	}
	return nil
}

// ShutMe Actions: command and event
func ShutMeRun() error {
	var args []string
	cmds := strings.SplitN(ShutMeCmd, " ", 2)
	if len(cmds) > 1 {
		args = strings.Split(cmds[1], " ")
	}
	cmdl := exec.Command(cmds[0], args...)
	_, err := cmdl.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
