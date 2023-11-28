package cmds

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

/*
-s	install; uninstall; start; stop; restart
-t	<hostname>
-n	<number> 	Number of cycle
-i	<second>	Cycle Interval in Second
-c	<command> 	Ignore the default command
-f	bool		Forces the default command to shut down
-h	bool		Try the hibernation command (suspend to hard disk)
*/
var ExecPath string
var BaseName string

var (
	Flag_s string
	Flag_t string
	Flag_c string
	Flag_n int
	Flag_i int
	Flag_f bool
	Flag_b bool
	Flag_y bool
)

func init() {
	ExecPath, _ = os.Executable()
	ExecPath, _ = filepath.EvalSymlinks(ExecPath)
	ext := filepath.Ext(ExecPath)
	BaseName = strings.TrimSuffix(ExecPath, ext)
	ExecPath = filepath.Dir(ExecPath)

	flag.StringVar(&Flag_s, "s", "", "Service control action: \"install\" \"uninstall\" \"start\" \"stop\" \"restart\".")
	flag.StringVar(&Flag_t, "t", "", "Specify the hostname or IP address, and detect if it is online.")
	flag.StringVar(&Flag_c, "c", "", "Action taken upon detecting remote host offline.\nBy default, the shutdown command will be automatically determined.")
	flag.IntVar(&Flag_n, "n", 10, "Cumulative number of disconnections.")
	flag.IntVar(&Flag_i, "i", 60, "Detection interval in second.")
	flag.BoolVar(&Flag_f, "f", false, "Add an enforcement option for shutdown commands.")
	flag.BoolVar(&Flag_b, "b", false, "Try the hibernation (suspend to hard disk) command.")
	flag.BoolVar(&Flag_y, "y", false, "No confirm prompts for Interactive Mode and Install Service.")
	flag.Parse()
}

// Define and perse command line flags
// Param: none
// Retrun: error
func CmdPerser() error {

	// Check Flags logic.
	// The "-t" option must be specified during interactive or service install.
	if Flag_t == "" && (Flag_s == "" || Flag_s == "install") {
		return fmt.Errorf("the '-t' option must be specified during interactive or service install")
	}

	// -c , -f , -b flag mutually exclusive
	i := 0
	if len(Flag_c) != 0 {
		i++
	}

	if Flag_f {
		i++
	}

	if Flag_b {
		i++
	}

	if i > 1 {
		return fmt.Errorf("the -c, -f, -b options are not allowed to be used simultaneously")
	}

	if len(Flag_c) == 0 {
		return ShutMeCmdPerse()
	}

	return nil
}

var ShutMeCmd string

// Command line to perform shutdown for current system
// Param : none
// Retern: string, If null, the current system is not supported
func ShutMeCmdPerse() error {
	switch runtime.GOOS {
	case "linux":
		if !Flag_b {
			ShutMeCmd = "shutdown -t 0 -h"
			if Flag_f {
				ShutMeCmd += " -n"
			}
		} else {
			ShutMeCmd = "echo disk > /sys/power/state"
			// ShutMeCmd ="systemctl hibernate"
			// Check to see if suspend to hard disk is supported: "cat /sys/power/state"
		}

	case "windows":
		if !Flag_b {
			ShutMeCmd = "shutdown -s -t 0"
		} else {
			ShutMeCmd = "shutdown -h -t 0"
		}
		if Flag_f {
			ShutMeCmd += " -f"
		}

	//case "darwin":
	//	ShutMeCmd = "sudo shutdown -h now"

	default:
		ShutMeCmd = ""
		return fmt.Errorf("the shutdown command on this OS is not currently supported, you can specify it with option \"-c\"")
	}
	return nil
}

// ShutMe Actions: command and event
func ShutMeRun() error {
	var args []string
	cmdShut := strings.SplitN(ShutMeCmd, " ", 2)
	if len(cmdShut) > 1 {
		args = strings.Split(cmdShut[1], " ")
	}
	cmdl := exec.Command(cmdShut[0], args...)
	_, err := cmdl.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
