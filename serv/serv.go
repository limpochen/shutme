package serv

import (
	"fmt"
	"shutme/cmds"
	"shutme/shutmedown"
	"strconv"

	"github.com/kardianos/service"
)

const (
	ServName = "ShutMe"
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

	args := []string{
		"-t", cmds.Flag_t,
		"-n", strconv.Itoa(cmds.Flag_n),
		"-i", strconv.Itoa(cmds.Flag_i),
		"-c", fmt.Sprintf("\"%s\"", shutmedown.ShutMeCmd),
	}

	// system service description
	desc := fmt.Sprintf("Probe the remote host %s every %d seconds, ", cmds.Flag_t, cmds.Flag_i)
	desc += fmt.Sprintf("detect offline for %d times (%d minutes)", cmds.Flag_n, cmds.Flag_i*cmds.Flag_n/60)
	desc += fmt.Sprintf(", and execute: \"%s\".", shutmedown.ShutMeCmd)

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
func ServCtrl(servs service.Service) error {
	var err error
	switch cmds.Flag_s {
	case "install":
		err = servs.Install()
		if err != nil {
			return err
		}

		fmt.Println("The service is installed.")
		fallthrough

	case "start":
		err = servs.Start()
		if err != nil {
			return err
		}

		fmt.Println("The service is started.")

	case "stop":
		err = servs.Stop()
		if err != nil {
			return err
		}

		fmt.Println("The service is stopped")

	case "restart":
		err = servs.Restart()
		if err != nil {
			return err
		}

		fmt.Println("The service is restarted.")

	case "uninstall":
		st, _ := servs.Status()
		if st == service.StatusRunning {
			service.Control(servs, "stop")
		}

		err = servs.Uninstall()
		if err != nil {
			return err
		}

		fmt.Println("The service is removed.")

	default:
		err = fmt.Errorf("error in service control command")
		return err
	}

	return nil
}
