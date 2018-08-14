package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {

	systray.SetIcon(getIcon("assets/hauxe.jpg"))
	systray.SetTitle("Ấm Sò")

	medium := systray.AddMenuItem("Medium Blog", "My Medium Blog")
	github := systray.AddMenuItem("Github Page", "My Github Page")
	stackoverflow := systray.AddMenuItem("StackOverflow Page", "StackOverflow Page")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quits this app")

	go func() {
		for {
			select {
			case <-medium.ClickedCh:
				openWebpage("https://medium.com/@hau12a1")
			case <-github.ClickedCh:
				openWebpage("https://github.com/hauxe")
			case <-stackoverflow.ClickedCh:
				openWebpage("https://stackoverflow.com/users/5039386/hau-ma")
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func openWebpage(url string) {
	cmd := exec.Command("open", "-a", "Google Chrome", url)
	err := cmd.Run()

	if err != nil {
		log.Println(err)
	}
}

func onExit() {
	// Cleaning stuff here.
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
