package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/shomali11/slacker"
)

// handleHello - Simple Hello handler
func handleHello(request slacker.Request, response slacker.ResponseWriter) {
	name := request.Param("name")
	if name == "" {
		response.Reply("Usage: @billy hello Name")
		return
	}
	response.Reply("Hey " + name + "!")
}

// handleWifi - Simple handler for giving out wifi information
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
	bot.Init(func() {
		log.Println("Connected!")
	})

	bot.Err(func(err string) {
		log.Println(err)
	})

	bot.Command("hello <name>", "Say hello to someone", handleHello)
	bot.Command("wifi", "Retrieve wifi information", handleWifi)
	bot.Command("ping", "Check on the bot", func(request slacker.Request, response slacker.ResponseWriter) {
		response.Reply("pong")
	})

	bot.DefaultCommand(func(request slacker.Request, response slacker.ResponseWriter) {
		userInfo, _ := bot.GetUserInfo(request.Event().User)
		name := userInfo.Profile.DisplayName
		if name == "" {
			name = strings.Split(userInfo.Profile.RealName, " ")[0]
		}
		response.Reply("I'm not sure what you're looking for, " + name + ". Try the 'help' command.")
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		panic(err)
	}
}
