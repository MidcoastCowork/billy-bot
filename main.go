package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/shomali11/slacker"
)

func handle(request slacker.Request, response slacker.ResponseWriter) {
	name := request.Param("name")
	if name == "" {
		response.Reply("Usage: @billy hello Name")
		return
	}
	response.Reply("Hey " + name + "!")
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Billy!")
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

	bot := slacker.NewClient(os.Getenv("API_TOKEN"))
	bot.Init(func() {
		log.Println("Connected!")
	})

	bot.Err(func(err string) {
		log.Println(err)
	})
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
