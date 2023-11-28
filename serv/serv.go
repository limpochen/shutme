package serv

import (
	"fmt"
	"os"
	"shutme/cmds"

	"github.com/kardianos/service"
)

const (
	ServName = "shutme"
	DispName = "ShutMe Helper"
)

type Config struct {
	Name             string
	DisplayName      string
	Description      string
	Exec             string
	Args             []string
	WorkingDirectory string
	Option           service.KeyValue
}

// Initialize system service
// Param : none
// Return: service.Service, error
func ServInit() (service.Service, error) {
	var servs service.Service
	var err error
	// Reassemble command-line arguments into system service calls

	var args []string
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-s" {
			i++
			continue
		}

		if os.Args[i] == "-y" {
			continue
		}

		args = append(args, os.Args[i])
	}

	// system service description
	desc := fmt.Sprintf("Probe the remote host %s every %d seconds, ", cmds.Flag_t, cmds.Flag_i)
	desc += fmt.Sprintf("detect offline for %d times (about %d minutes)", cmds.Flag_n, cmds.Flag_i*cmds.Flag_n/60)
	desc += fmt.Sprintf(", and execute: \"%s\".", cmds.Flag_c)

	svcConfig := &service.Config{
		Name:             ServName,
		DisplayName:      DispName,
		Description:      desc,
		Arguments:        args,
		WorkingDirectory: cmds.ExecPath,
	}

	prg := &program{}
	servs, err = service.New(prg, svcConfig)
	if err != nil {
		return nil, err
	}
	return servs, nil
}

// System service command line control
// Param : service.Service
// Return: error
func ServCtrl(servs service.Service) (status string, err error) {
	switch cmds.Flag_s {
	case "install":
		err = servs.Install()
		if err != nil {
			return "", err
		}

		status = "installed"
		fallthrough

	case "start":
		err = servs.Start()
		if err != nil {
			return "", err
		}

		if status == "" {
			status = "started"
		} else {
			status += " and started"
		}

	case "stop":
		err = servs.Stop()
		if err != nil {
			return "", err
		}

		status = "stopped"

	case "restart":
		err = servs.Restart()
		if err != nil {
			return "", err
		}

		status = "restarted."

	case "uninstall":
		st, _ := servs.Status()
		if st == service.StatusRunning {
			service.Control(servs, "stop")
			status = "stopped"
		}

		err = servs.Uninstall()
		if err != nil {
			return "", err
		}

		if status == "" {
			status = "removed"
		} else {
			status += " and removed"
		}

	default:
		err = fmt.Errorf("error in service control command")
		return "", err
	}

	return "Service is " + status + ".", nil
}
