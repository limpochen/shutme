package cmds

import (
	"os/exec"
	"runtime"
	"strings"
)

// Command line to perform shutdown for current system
// Param : none
// Retern: string, If null, the current system is not supported
func ShutMeCmdPerse() string {
	var shutmecmd string

	switch runtime.GOOS {
	case "linux":
		if !Flag_b {
			shutmecmd = "shutdown -t 0 -h"
			if Flag_f {
				shutmecmd = ShutMeCmd + " -n"
			}
		} else {
			shutmecmd = "echo disk > /sys/power/state"
			// ShutMeCmd ="systemctl hibernate"
			// Check to see if suspend to hard disk is supported: "cat /sys/power/state"
		}

	case "windows":
		if !Flag_b {
			shutmecmd = "shutdown -s -t 0"
		} else {
			shutmecmd = "shutdown -h -t 0"
		}
		if Flag_f {
			shutmecmd = shutmecmd + " -f"
		}

	//case "darwin":
	//	ShutMeCmd = "sudo shutdown -h now"

	default:
		shutmecmd = ""
	}
	return shutmecmd
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
