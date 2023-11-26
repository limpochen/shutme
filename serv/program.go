package serv

import (
	"log"
	"os"
	"shutme/cmds"

	"github.com/kardianos/service"
)

type program struct{}

// System service program start enrty
// Param : service.Service
// Return: error
func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	log.Println("Shutme Service started.")
	return nil
}

// System service program stop enrty
// Param : service.Service
// Return: error
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	log.Println("Shutme Service stopped.")
	return nil
}

// System service program execution entry
// Param : none
// Return: none
func (p *program) run() {
	if _, err := cmds.Ping(cmds.Flag_t); err != nil {
		//cmds.MyLog(cmds.Error, "Communication cannot be established with the remote host, service terminates.\n")
		log.Printf("Communication cannot be established with the remote host, service terminates.\n\n")
		os.Exit(1)
	}
	cmds.ProbeRemote()
}
