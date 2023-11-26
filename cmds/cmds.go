package cmds

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
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
var ShutMeCmd string
var MyLogFile string

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
	MyLogFile = strings.TrimSuffix(ExecPath, ext) + ".log"
	ExecPath = filepath.Dir(ExecPath)
}

// Define and perse command line flags
// Param: none
// Retrun: error
func CmdLine() error {
	flag.StringVar(&Flag_s, "s", "", "Service control action: \"install\" \"uninstall\" \"start\" \"stop\" \"restart\".")
	flag.StringVar(&Flag_t, "t", "", "Specify the hostname or IP address, and detect if it is online.")
	flag.StringVar(&Flag_c, "c", "", "Action taken upon detecting remote host offline.")
	flag.IntVar(&Flag_n, "n", 10, "Cumulative number of disconnections.")
	flag.IntVar(&Flag_i, "i", 60, "Detection interval (in second).")
	flag.BoolVar(&Flag_f, "f", false, "Forces the default command to shut down.")
	flag.BoolVar(&Flag_b, "b", false, "Try the hibernation command (suspend to hard disk).")
	flag.BoolVar(&Flag_y, "y", false, "No confirm prompts for Interactive Mode and Install Service.")
	flag.Parse()

	// Check Flags logic.
	// The "-t" option must be specified during interactive or service install.
	if Flag_t == "" && (Flag_s == "" || Flag_s == "install") {
		return fmt.Errorf("the '-t' option must be specified during interactive or service install")
	}

	// -c , -f , -h flag mutually exclusive
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
		return fmt.Errorf("the -c, -f, -h options are not allowed to be used simultaneously")
	}

	if len(Flag_c) != 0 {
		ShutMeCmd = Flag_c
	} else {
		ShutMeCmd = ShutMeCmdPerse()
		if ShutMeCmd == "" {
			return fmt.Errorf("the shutdown command on this OS is not currently supported, you can specify it with option \"-c\"")
		}
	}
	return nil
}

// Print some necessary information and confirm that you want to continue
// Param : none
// Return: boolean
func Confirm() bool {
	var res string

	fmt.Printf("Attempt to detect remote host status... ")
	if _, err := Ping(Flag_t); err != nil {
		fmt.Fprintf(os.Stderr, "Failed.\nCommunication cannot be established with the remote host %s, program terminates.\n", Flag_t) //todo
		return false
	} else {
		fmt.Printf("OK.\n")
	}

	if !Flag_y {

		fmt.Println("WARNING, Once the network failure occurs after the program is running, it will trigger the shutdown behavior.")
		fmt.Printf("Are you sure you want to do this? \nPress 'YES' to continue:")
		fmt.Scanln(&res)
		if strings.ToUpper(res) != "YES" {
			return false
		}
	}

	return true
}
