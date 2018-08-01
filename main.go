package main

import (
	"os"

	"github.com/shomali11/slacker"
)

func handle(request slacker.Request, response slacker.ResponseWriter) {
	response.Reply("Hey!")
}

func handleWifi(request slacker.Request, response slacker.ResponseWriter) {
	ssid := os.Getenv("WIFI_SSID")
	sauce := os.Getenv("WIFI_SAUCE")
	if sauce != "" && ssid != "" {
		response.Reply("To connect our killer wifi use network name " + ssid + " and password " + sauce)
	} else if sauce == "" && ssid != "" {
		response.Reply("The wifi isn't protected and the SSID is " + ssid)
	} else {
		response.Reply("Wifi isn't setup! Set it up and make some friends!")
	}

}

func main() {
	bot := slacker.NewClient(os.Getenv("API_TOKEN"))
	bot.Command("hello <name>", "Say hello to someone", handle)
	bot.Command("wifi", "get the wifi information", handleWifi)
	bot.Command("ping", "Ping!", func(request slacker.Request, response slacker.ResponseWriter) {
		response.Reply("pong")
	})
	err := bot.Listen()
	if err != nil {
		panic(err)
	}
}
