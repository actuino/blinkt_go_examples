package main

import (
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"log"
	"runtime"
	"time"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"strconv"
)
import . "github.com/alexellis/blinkt_go"


type colors map[string][3]int

type FileMessage struct {
	Name      string    `json:"Name"`
	Channel string 	`json:"Channel"`
	Type    string `json:"Type"`
	Payload [8][8]int `json:"Payload"`
}


func main() {
	brightness := 0.3
	display_server_host := "192.168.7.3"
	if "" != os.Getenv("DISPLAY_SERVER_HOST") {
		display_server_host = os.Getenv("DISPLAY_SERVER_HOST")
	}
	display_server_port := 80
	if "" != os.Getenv("DISPLAY_SERVER_PORT") {
		port, err := strconv.Atoi(os.Getenv("DISPLAY_SERVER_PORT"))
		if err == nil {
			display_server_port = port
		}	 
	}
	log.Println("--- Server: ", display_server_host, display_server_port)
    blinkt := NewBlinkt(brightness)
    blinkt.SetClearOnExit(true)
    blinkt.Setup()

    Delay(100)
        
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	dat, err := ioutil.ReadFile("./colors.json")
	if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var palette colors
    json.Unmarshal(dat, &palette)

	c, err := gosocketio.Dial(
		gosocketio.GetUrl(display_server_host, display_server_port, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Fatal(err)
	}

	err = c.On("message", func(h *gosocketio.Channel, args string) {
		log.Println("--- Got message: ", args)
	})
	if err != nil {
		log.Fatal(err)
	}
	err = c.On("file", func(h *gosocketio.Channel, args FileMessage) {
		log.Println("--- Got json file: ", args)
		blinkt.Clear()
		for i, colorIndex := range args.Payload[0] {
			rgb := palette[strconv.Itoa(colorIndex)];
			blinkt.SetPixel(i, rgb[0], rgb[1], rgb[2])
            blinkt.Show()
		}
	})
	if err != nil {
		log.Fatal(err)
	}	

	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Fatal("Disconnected")
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("Connected")
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(1 * time.Second)
	}	
	c.Close()

	log.Println(" [x] Complete")
}
